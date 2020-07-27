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

var WXServerUrl = "https://api.weixin.qq.com"

const (
	ContentTypeApplicationJson = "application/json;charset=utf-8"
)

// HTTPGet GET 请求
func HTTPGet(uri string) (resp []byte, err error) {
	uri = applyAccessToken(uri)
	response, err := http.Get(WXServerUrl + uri)
	if err != nil {
		return
	}
	defer response.Body.Close()
	return responseFilter(response)
}

//HTTPPost POST 请求
func HTTPPost(uri string, payload io.Reader, contentType string) (resp []byte, err error) {
	uri = applyAccessToken(uri)
	response, err := http.Post(WXServerUrl + uri, contentType, payload)
	if err != nil {
		return
	}
	defer response.Body.Close()
	return responseFilter(response)
}

func applyAccessToken(oldUrl string) (newUrl string) {
	accessToken := getAccessToken()
	if strings.Contains(oldUrl, "?") {
		return oldUrl+"&access_token="+accessToken
	}
	return oldUrl+"?access_token="+accessToken
}

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

	if errorResponse.Errcode != 0 {
		err = errors.New(string(resp))
		return
	}

	return
}

var refreshAccessTokenLock sync.Mutex

func getAccessToken() (accessToken string) {
	accessToken, err := accessTokenCache.Fetch(Appid)
	if accessToken != "" {
		return
	}

	refreshAccessTokenLock.Lock()
	defer refreshAccessTokenLock.Unlock()

	accessToken, err = accessTokenCache.Fetch(Appid)
	if accessToken != "" {
		return
	}

	accessToken, expiresIn, err := refreshAccessTokenHandler()
	if err != nil {
		return
	}

	// 提前过期 提供冗余时间
	expiresIn = int(0.9 * float64(expiresIn))
	d := time.Duration(expiresIn) * time.Second
	_ = accessTokenCache.Save(Appid, accessToken, d)

	return
}

/*
从微信服务器获取新的 AccessToken

See: https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html
*/
func RefreshAccessTokenFromWXServer() (accessToken string, expiresIn int, err error) {
	params := url.Values{}
	params.Add("appid", Appid)
	params.Add("secret", Secret)
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
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		Errcode     string `json:"errcode"`
		Errmsg      string `json:"errmsg"`
	}{}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}

	if result.AccessToken == "" {
		err = fmt.Errorf("%s", string(resp))
		return
	}

	return result.AccessToken, result.ExpiresIn, nil
}
