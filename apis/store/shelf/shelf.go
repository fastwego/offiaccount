// Package shelf 微信小店/货架管理
package shelf

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiAdd     = "/merchant/shelf/add"
	apiDel     = "/merchant/shelf/del"
	apiMod     = "/merchant/shelf/mod"
	apiGetAll  = "/merchant/shelf/getall"
	apiGetById = "/merchant/shelf/getbyid"
)

/*
增加货架



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/shelf/add?access_token=ACCESS_TOKEN
*/
func Add(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAdd, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
删除货架



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/shelf/del?access_token=ACCESS_TOKEN
*/
func Del(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDel, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
修改货架



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/shelf/mod?access_token=ACCESS_TOKEN
*/
func Mod(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMod, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取所有货架



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

GET https://api.weixin.qq.com/merchant/shelf/getall?access_token=ACCESS_TOKEN
*/
func GetAll(ctx *offiaccount.OffiAccount) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetAll)
}

/*
根据货架ID获取货架信息



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/shelf/getbyid?access_token=ACCESS_TOKEN
*/
func GetById(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetById, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}
