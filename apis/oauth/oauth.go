// Package oauth 微信网页开发(oauth)
package oauth

import (
	"bytes"
	"net/url"

	"github.com/fastwego/offiaccount"
)

const (
	apiAccessToken  = "/sns/oauth2/access_token"
	apiRefreshToken = "/sns/oauth2/refresh_token"
	apiGetUserInfo  = "/sns/userinfo"
	apiAuth         = "/sns/auth"
)

/*
通过code换取网页授权access_token

注意：由于公众号的secret和获取到的access_token安全级别都非常高，必须只保存在服务器，不允许传给客户端

See: https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html

POST https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code
*/
func AccessToken(ctx *offiaccount.OffiAccount, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAccessToken+"?"+params.Encode(), bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
刷新access_token

由于access_token拥有较短的有效期，当access_token超时后，可以使用refresh_token进行刷新，refresh_token有效期为30天，当refresh_token失效之后，需要用户重新授权

See: https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html

POST https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=APPID&grant_type=refresh_token&refresh_token=REFRESH_TOKEN
*/
func RefreshToken(ctx *offiaccount.OffiAccount, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiRefreshToken+"?"+params.Encode(), bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
拉取用户信息

如果网页授权作用域为snsapi_userinfo，则此时开发者可以通过access_token和openid拉取用户信息了

See: https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html

POST https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
*/
func GetUserInfo(ctx *offiaccount.OffiAccount, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetUserInfo+"?"+params.Encode(), bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
检验授权凭证（access_token）是否有效



See: https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html

GET https://api.weixin.qq.com/sns/auth?access_token=ACCESS_TOKEN&openid=OPENID
*/
func Auth(ctx *offiaccount.OffiAccount, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiAuth + "?" + params.Encode())
}
