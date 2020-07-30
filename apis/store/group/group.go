// Package group 微信小店/分组管理
package group

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiAdd         = "/merchant/group/add"
	apiDel         = "/merchant/group/del"
	apiPropertyMod = "/merchant/group/propertymod"
	apiProductMod  = "/merchant/group/productmod"
	apiGetAll      = "/merchant/group/getall"
	apiGetById     = "/merchant/group/getbyid"
)

/*
增加分组



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/group/add?access_token=ACCESS_TOKEN
*/
func Add(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAdd, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
删除分组



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/group/del?access_token=ACCESS_TOKEN
*/
func Del(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDel, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
修改分组属性



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/group/propertymod?access_token=ACCESS_TOKEN
*/
func PropertyMod(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiPropertyMod, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
修改分组商品



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/group/productmod?access_token=ACCESS_TOKEN
*/
func ProductMod(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiProductMod, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取所有分组



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

GET https://api.weixin.qq.com/merchant/group/getall?access_token=ACCESS_TOKEN
*/
func GetAll(ctx *offiaccount.OffiAccount) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetAll)
}

/*
根据分组ID获取分组信息



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/group/getbyid?access_token=ACCESS_TOKEN
*/
func GetById(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetById, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}
