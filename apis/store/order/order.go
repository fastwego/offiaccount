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
