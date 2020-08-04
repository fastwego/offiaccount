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
	return ctx.Client.HTTPPost(apiCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除商品



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/del?access_token=ACCESS_TOKEN
*/
func Del(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDel, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
修改商品



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/update?access_token=ACCESS_TOKEN
*/
func Update(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询商品



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/get?access_token=ACCESS_TOKEN
*/
func Get(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取指定状态的所有商品



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/getbystatus?access_token=ACCESS_TOKEN
*/
func GetByStatus(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetByStatus, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
商品上下架



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/modproductstatus?access_token=ACCESS_TOKEN
*/
func ModProductStatus(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiModProductStatus, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取指定分类的所有子分类



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/category/getsub?access_token=ACCESS_TOKEN
*/
func GetSub(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetSub, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取指定子分类的所有SKU



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/category/getsku?access_token=ACCESS_TOKEN
*/
func GetSku(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetSku, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取指定分类的所有属性



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/category/getproperty?access_token=ACCESS_TOKEN
*/
func GetProperty(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetProperty, bytes.NewReader(payload), "application/json;charset=utf-8")
}
