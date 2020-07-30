// Package product 微信小店/商品管理
package product

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiCreate           = "/merchant/create"
	apiDel              = "/merchant/del"
	apiUpdate           = "/merchant/update"
	apiGet              = "/merchant/get"
	apiGetByStatus      = "/merchant/getbystatus"
	apiModProductStatus = "/merchant/modproductstatus"
	apiGetSub           = "/merchant/category/getsub"
	apiGetSku           = "/merchant/category/getsku"
	apiGetProperty      = "/merchant/category/getproperty"
)

/*
增加商品



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/create?access_token=ACCESS_TOKEN
*/
func Create(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreate, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
删除商品



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/del?access_token=ACCESS_TOKEN
*/
func Del(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDel, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
修改商品



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/update?access_token=ACCESS_TOKEN
*/
func Update(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdate, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
查询商品



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/get?access_token=ACCESS_TOKEN
*/
func Get(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGet, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取指定状态的所有商品



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/getbystatus?access_token=ACCESS_TOKEN
*/
func GetByStatus(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetByStatus, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
商品上下架



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/modproductstatus?access_token=ACCESS_TOKEN
*/
func ModProductStatus(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiModProductStatus, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取指定分类的所有子分类



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/category/getsub?access_token=ACCESS_TOKEN
*/
func GetSub(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetSub, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取指定子分类的所有SKU



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/category/getsku?access_token=ACCESS_TOKEN
*/
func GetSku(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetSku, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取指定分类的所有属性



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/category/getproperty?access_token=ACCESS_TOKEN
*/
func GetProperty(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetProperty, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}
