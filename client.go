// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package offiaccount

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var (
	WXServerUrl            = "https://api.weixin.qq.com" // 微信 api 服务器地址
	UserAgent              = "fastwego/offiaccount"
	ErrorAccessTokenExpire = errors.New("access token expire")
)

/*
HttpClient 用于向公众号接口发送请求
*/
type Client struct {
	Ctx *OffiAccount
}

// HTTPGet GET 请求
func (client *Client) HTTPGet(uri string) (resp []byte, err error) {
	newUrl, err := client.applyAccessToken(uri)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodGet, WXServerUrl+newUrl, nil)
	if err != nil {
		return
	}

	return client.httpDo(req)
}

//HTTPPost POST 请求
func (client *Client) HTTPPost(uri string, payload io.Reader, contentType string) (resp []byte, err error) {
	newUrl, err := client.applyAccessToken(uri)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, WXServerUrl+newUrl, payload)
	if err != nil {
		return
	}

	req.Header.Add("Content-Type", contentType)

	return client.httpDo(req)
}

//httpDo 执行 请求
func (client *Client) httpDo(req *http.Request) (resp []byte, err error) {
	req.Header.Add("User-Agent", UserAgent)

	if client.Ctx.Logger != nil {
		client.Ctx.Logger.Printf("%s %s Headers %v", req.Method, req.URL.String(), req.Header)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()

	resp, err = responseFilter(response)

	// 发现 access_token 过期
	if err == ErrorAccessTokenExpire {

		// 主动 通知 access_token 过期
		err = client.Ctx.AccessToken.NoticeAccessTokenExpireHandler(client.Ctx)
		if err != nil {
			return
		}

		// 通知到位后 access_token 会被刷新，那么可以 retry 了
		var accessToken string
		accessToken, err = client.Ctx.AccessToken.GetAccessTokenHandler(client.Ctx)
		if err != nil {
			return
		}

		// 换新
		q := req.URL.Query()
		q.Set("access_token", accessToken)
		req.URL.RawQuery = q.Encode()

		if client.Ctx.Logger != nil {
			client.Ctx.Logger.Printf("retry %s %s Headers %v", req.Method, req.URL.String(), req.Header)
		}

		response, err = http.DefaultClient.Do(req)
		if err != nil {
			return
		}
		defer response.Body.Close()

		resp, err = responseFilter(response)
	}

	return
}

/*
在请求地址上附加上 access_token
*/
func (client *Client) applyAccessToken(oldUrl string) (newUrl string, err error) {
	accessToken, err := client.Ctx.AccessToken.GetAccessTokenHandler(client.Ctx)
	if err != nil {
		return
	}
	if strings.Contains(oldUrl, "?") {
		newUrl = oldUrl + "&access_token=" + accessToken
	} else {
		newUrl = oldUrl + "?access_token=" + accessToken
	}
	return
}

/*
筛查微信 api 服务器响应，判断以下错误：

- http 状态码 不为 200

- 接口响应错误码 errcode 不为 0
*/
func responseFilter(response *http.Response) (resp []byte, err error) {
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("Status %s", response.Status)
		return
	}

	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	errorResponse := struct {
		Errcode int64  `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}{}
	err = json.Unmarshal(resp, &errorResponse)
	if err != nil {
		return
	}

	// 40001(覆盖刷新超过5min后，使用旧 access_token 报错) 获取 access_token 时 AppSecret 错误，或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，或查看是否正在为恰当的公众号调用接口
	// 42001(超过 7200s 后 报错) - access_token 超时，请检查 access_token 的有效期，请参考基础支持 - 获取 access_token 中，对 access_token 的详细机制说明
	if errorResponse.Errcode == 42001 || errorResponse.Errcode == 40001 {
		err = ErrorAccessTokenExpire
		return
	}
	if errorResponse.Errcode != 0 {
		err = errors.New(string(resp))
		return
	}
	return
}

// 防止多个 goroutine 并发刷新冲突
var refreshAccessTokenLock sync.Mutex

/*
从 公众号实例 的 AccessToken 管理器 获取 access_token

如果没有 access_token 或者 已过期，那么刷新

获得新的 access_token 后 过期时间设置为 0.9 * expiresIn 提供一定冗余
*/
func GetAccessToken(ctx *OffiAccount) (accessToken string, err error) {
	accessToken, err = ctx.AccessToken.Cache.Fetch(ctx.Config.Appid)
	if accessToken != "" {
		return
	}

	refreshAccessTokenLock.Lock()
	defer refreshAccessTokenLock.Unlock()

	accessToken, err = ctx.AccessToken.Cache.Fetch(ctx.Config.Appid)
	if accessToken != "" {
		return
	}

	accessToken, expiresIn, err := refreshAccessTokenFromWXServer(ctx.Config.Appid, ctx.Config.Secret)
	if err != nil {
		return
	}

	// 本地缓存 access_token
	d := time.Duration(expiresIn) * time.Second
	_ = ctx.AccessToken.Cache.Save(ctx.Config.Appid, accessToken, d)

	if ctx.Logger != nil {
		ctx.Logger.Printf("%s %s %d\n", "refreshAccessTokenFromWXServer", accessToken, expiresIn)
	}

	return
}

/*
NoticeAccessTokenExpire 只需将本地存储的 access_token 删除，即完成了 access_token 已过期的 主动通知

retry 请求的时候，会发现本地没有 access_token ，从而触发refresh
*/
func NoticeAccessTokenExpire(ctx *OffiAccount) (err error) {
	if ctx.Logger != nil {
		ctx.Logger.Println("NoticeAccessTokenExpire")
	}

	err = ctx.AccessToken.Cache.Delete(ctx.Config.Appid)
	return
}

/*
从微信服务器获取新的 AccessToken

See: https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html
*/
func refreshAccessTokenFromWXServer(appid string, secret string) (accessToken string, expiresIn int, err error) {
	params := url.Values{}
	params.Add("appid", appid)
	params.Add("secret", secret)
	params.Add("grant_type", "client_credential")
	url := WXServerUrl + "/cgi-bin/token?" + params.Encode()

	response, err := http.Get(url)
	if err != nil {
		return
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("GET %s RETURN %s", url, response.Status)
		return
	}

	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	var result = struct {
		AccessToken string  `json:"access_token"`
		ExpiresIn   int     `json:"expires_in"`
		Errcode     float64 `json:"errcode"`
		Errmsg      string  `json:"errmsg"`
	}{}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		err = fmt.Errorf("Unmarshal error %s", string(resp))
		return
	}

	if result.AccessToken == "" {
		err = fmt.Errorf("%s", string(resp))
		return
	}

	return result.AccessToken, result.ExpiresIn, nil
}
