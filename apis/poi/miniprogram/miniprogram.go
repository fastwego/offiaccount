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

// Package miniprogram 微信门店小程序
package miniprogram

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiGetMerchantCategory  = "/wxa/get_merchant_category"
	apiApplyMerchant        = "/wxa/apply_merchant"
	apiGetMerchantAuditInfo = "/wxa/get_merchant_audit_info"
	apiModifyMerchant       = "/wxa/modify_merchant"
	apiGetDistrict          = "/wxa/get_district"
	apiSearchMapPoi         = "/wxa/search_map_poi"
	apiCreateMapPoi         = "/wxa/create_map_poi"
	apiAddStore             = "/wxa/add_store"
	apiUpdateStore          = "/wxa/update_store"
	apiCardStorewxaGet      = "/card/storewxa/get"
	apiGetStoreInfo         = "/wxa/get_store_info"
	apiGetStoreList         = "/wxa/get_store_list"
	apiDelStore             = "/wxa/del_store"
)

/*
拉取门店小程序类目



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

GET https://api.weixin.qq.com/wxa/get_merchant_category?access_token=TOKEN
*/
func GetMerchantCategory(ctx *offiaccount.OffiAccount) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetMerchantCategory)
}

/*
创建门店小程序



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/apply_merchant?access_token=TOKEN
*/
func ApplyMerchant(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiApplyMerchant, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
查询门店小程序审核结果



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

GET https://api.weixin.qq.com/wxa/get_merchant_audit_info?access_token=TOKEN
*/
func GetMerchantAuditInfo(ctx *offiaccount.OffiAccount) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetMerchantAuditInfo)
}

/*
修改门店小程序信息



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/modify_merchant?access_token=TOKEN
*/
func ModifyMerchant(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiModifyMerchant, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
从腾讯地图拉取省市区信息



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

GET https://api.weixin.qq.com/wxa/get_district?access_token=TOKEN
*/
func GetDistrict(ctx *offiaccount.OffiAccount) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetDistrict)
}

/*
在腾讯地图中搜索门店



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/search_map_poi?access_token=TOKEN
*/
func SearchMapPoi(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSearchMapPoi, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
在腾讯地图中创建门店



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/create_map_poi?access_token=TOKEN
*/
func CreateMapPoi(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreateMapPoi, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
添加门店



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/add_store?access_token=TOKEN
*/
func AddStore(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddStore, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
更新门店信息



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/update_store?access_token=TOKEN
*/
func UpdateStore(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdateStore, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取门店小程序配置的卡券



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/card/storewxa/get?access_token=ACCESS_TOKEN
*/
func CardStorewxaGet(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCardStorewxaGet, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取单个门店信息



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/get_store_info?access_token=TOKEN
*/
func GetStoreInfo(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetStoreInfo, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取门店信息列表



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/get_store_list?access_token=TOKEN
*/
func GetStoreList(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetStoreList, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
删除门店



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/del_store?access_token=TOKEN
*/
func DelStore(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelStore, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}
