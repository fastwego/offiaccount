// Package stock 微信小店/库存管理
package stock

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiAdd    = "/merchant/stock/add"
	apiReduce = "/merchant/stock/reduce"
)

/*
增加库存



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/stock/add?access_token=ACCESS_TOKEN
*/
func Add(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAdd, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
减少库存



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/stock/reduce?access_token=ACCESS_TOKEN
*/
func Reduce(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiReduce, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}
