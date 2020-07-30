// Package order 微信小店/订单管理
package order

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiGetById     = "/merchant/order/getbyid"
	apiGetByFilter = "/merchant/order/getbyfilter"
	apiSetDelivery = "/merchant/order/setdelivery"
	apiClose       = "/merchant/order/close"
)

/*
根据订单ID获取订单详情


See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/order/getbyid?access_token=ACCESS_TOKEN
*/
func GetById(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetById, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
根据订单状态/创建时间获取订单详情



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/order/getbyfilter?access_token=ACCESS_TOKEN
*/
func GetByFilter(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetByFilter, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
设置订单发货信息



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/order/setdelivery?access_token=ACCESS_TOKEN
*/
func SetDelivery(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetDelivery, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
关闭订单



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/order/close?access_token=ACCESS_TOKEN
*/
func Close(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiClose, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}
