/*
微信公众平台开发 SDK

See: https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html
*/
package offiaccount

import (
	"github.com/faabiosr/cachego"
	cachegosync "github.com/faabiosr/cachego/sync"
)

type RefreshAccessTokenFunc func() (accessToken string, expiresIn int, err error)

var Appid string
var Secret string

var accessTokenCache cachego.Cache
var refreshAccessTokenHandler RefreshAccessTokenFunc

func init() {
	accessTokenCache = cachegosync.New()
	refreshAccessTokenHandler = RefreshAccessTokenFromWXServer
}

func SetAccessTokenCache(cacheDriver cachego.Cache) {
	accessTokenCache = cacheDriver
}

func SetRefreshAccessTokenHandler(f RefreshAccessTokenFunc) {
	refreshAccessTokenHandler = f
}

func DeleteAccessTokenCache() {
	_ = accessTokenCache.Delete(Appid)
}
