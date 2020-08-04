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
	return ctx.Client.HTTPPost(apiAdd, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除分组



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/group/del?access_token=ACCESS_TOKEN
*/
func Del(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDel, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
修改分组属性



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/group/propertymod?access_token=ACCESS_TOKEN
*/
func PropertyMod(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiPropertyMod, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
修改分组商品



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/group/productmod?access_token=ACCESS_TOKEN
*/
func ProductMod(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiProductMod, bytes.NewReader(payload), "application/json;charset=utf-8")
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
	return ctx.Client.HTTPPost(apiGetById, bytes.NewReader(payload), "application/json;charset=utf-8")
}
