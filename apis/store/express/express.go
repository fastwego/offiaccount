// Package express 微信小店/邮费模板管理管理
package express

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiAdd     = "/merchant/express/add"
	apiDel     = "/merchant/express/del"
	apiUpdate  = "/merchant/express/update"
	apiGetById = "/merchant/express/getbyid"
	apiGetAll  = "/merchant/express/getall"
)

/*
增加邮费模板



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/express/add?access_token=ACCESS_TOKEN
*/
func Add(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAdd, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
删除邮费模板



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/express/del?access_token=ACCESS_TOKEN
*/
func Del(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDel, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
修改邮费模板



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/express/update?access_token=ACCESS_TOKEN
*/
func Update(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdate, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取指定ID的邮费模板



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/express/getbyid?access_token=ACCESS_TOKEN
*/
func GetById(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetById, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取所有邮费模板



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

GET https://api.weixin.qq.com/merchant/express/getall?access_token=ACCESS_TOKEN
*/
func GetAll(ctx *offiaccount.OffiAccount) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetAll)
}
