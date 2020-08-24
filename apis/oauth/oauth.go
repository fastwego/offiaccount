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

// Package oauth 微信网页开发(oauth)

/*
网页授权流程分为四步：

1、引导用户进入授权页面同意授权，获取code

2、通过code换取网页授权access_token（与基础支持中的access_token不同）

3、如果需要，开发者可以刷新网页授权access_token，避免过期

4、通过网页授权access_token和openid获取用户基本信息（支持UnionID机制）
*/
package oauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/fastwego/offiaccount"
)

var OauthAuthorizeServerUrl = "https://open.weixin.qq.com"

const (
	apiAuthorize    = "/connect/oauth2/authorize"
	apiAccessToken  = "/sns/oauth2/access_token"
	apiRefreshToken = "/sns/oauth2/refresh_token"
	apiUserInfo     = "/sns/userinfo"
	apiAuth         = "/sns/auth"
)

const (
	ScopeSnsapiBase     = "snsapi_base"
	ScopeSnsapiUserinfo = "snsapi_userinfo"
)

/*
获取 用户授权 跳转链接

以snsapi_base为scope发起的网页授权，是用来获取进入页面的用户的openid的，并且是静默授权并自动跳转到回调页的。用户感知的就是直接进入了回调页（往往是业务页面）

以snsapi_userinfo为scope发起的网页授权，是用来获取用户的基本信息的。但这种授权需要用户手动同意，并且由于用户同意过，所以无须关注，就可在授权后获取该用户的基本信息

如果用户同意授权，页面将跳转至 redirect_uri/?code=CODE&state=STATE

See: https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html
*/
func GetAuthorizeUrl(appid string, redirectUri string, scope string, state string) (authorizeUrl string) {
	params := url.Values{}
	params.Add("appid", appid)
	params.Add("redirectUri", redirectUri)
	params.Add("scope", scope)
	params.Add("state", state)
	return OauthAuthorizeServerUrl + apiAuthorize + "?" + params.Encode()
}

type OauthAccessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
}

/*
通过code换取网页授权access_token

注意：由于公众号的secret和获取到的access_token安全级别都非常高，必须只保存在服务器，不允许传给客户端

See: https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html

GET https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code
*/
func GetAccessToken(appid string, secret string, code string) (oauthAccessToken OauthAccessToken, err error) {
	params := url.Values{}
	params.Add("appid", appid)
	params.Add("secret", secret)
	params.Add("code", code)
	params.Add("grant_type", "authorization_code")

	uri := offiaccount.WXServerUrl + apiAccessToken + "?" + params.Encode()
	response, err := http.Get(uri)
	if err != nil {
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &oauthAccessToken)
	if err != nil {
		err = fmt.Errorf("%s", string(body))
		return
	}

	return
}

/*
刷新access_token

由于access_token拥有较短的有效期，当access_token超时后，可以使用refresh_token进行刷新，refresh_token有效期为30天，当refresh_token失效之后，需要用户重新授权

See: https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html

POST https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=APPID&grant_type=refresh_token&refresh_token=REFRESH_TOKEN
*/
func RefreshToken(appid string, refresh_token string) (oauthAccessToken OauthAccessToken, err error) {
	params := url.Values{}
	params.Add("appid", appid)
	params.Add("refresh_token", refresh_token)
	params.Add("grant_type", "refresh_token")

	uri := offiaccount.WXServerUrl + apiRefreshToken + "?" + params.Encode()
	response, err := http.Get(uri)
	if err != nil {
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &oauthAccessToken)
	if err != nil {
		err = fmt.Errorf("%s", string(body))
		return
	}

	return
}

const (
	LANG_zh_CN = "zh_CN"
	LANG_zh_TW = "zh_TW"
	LANG_en    = "en"
)

type OauthUserInfo struct {
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        string   `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

/*
拉取用户信息

如果网页授权作用域为snsapi_userinfo，则此时开发者可以通过access_token和openid拉取用户信息了

See: https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html

POST https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
*/
func GetUserInfo(access_token string, openid string, lang string) (oauthUserInfo OauthUserInfo, err error) {
	params := url.Values{}
	params.Add("access_token", access_token)
	params.Add("openid", openid)
	params.Add("lang", lang)

	uri := offiaccount.WXServerUrl + apiUserInfo + "?" + params.Encode()
	response, err := http.Get(uri)
	if err != nil {
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &oauthUserInfo)
	if err != nil {
		err = fmt.Errorf("%s", string(body))
		return
	}

	return
}

/*
检验授权凭证（access_token）是否有效



See: https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html

GET https://api.weixin.qq.com/sns/auth?access_token=ACCESS_TOKEN&openid=OPENID
*/
func Auth(access_token string, openid string) (isValid bool, err error) {
	params := url.Values{}
	params.Add("access_token", access_token)
	params.Add("openid", openid)

	uri := offiaccount.WXServerUrl + apiAuth + "?" + params.Encode()
	response, err := http.Get(uri)
	if err != nil {
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	s := struct {
		Errcode int    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}{}

	err = json.Unmarshal(body, &s)
	if err != nil {
		err = fmt.Errorf("%s", string(body))
		return
	}

	if s.Errcode == 0 {
		isValid = true
	}

	return
}
