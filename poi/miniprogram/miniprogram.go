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
func GetMerchantCategory() (resp []byte, err error) {
	return offiaccount.HTTPGet(apiGetMerchantCategory)
}

/*
创建门店小程序



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/apply_merchant?access_token=TOKEN
*/
func ApplyMerchant(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiApplyMerchant, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询门店小程序审核结果



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

GET https://api.weixin.qq.com/wxa/get_merchant_audit_info?access_token=TOKEN
*/
func GetMerchantAuditInfo() (resp []byte, err error) {
	return offiaccount.HTTPGet(apiGetMerchantAuditInfo)
}

/*
修改门店小程序信息



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/modify_merchant?access_token=TOKEN
*/
func ModifyMerchant(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiModifyMerchant, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
从腾讯地图拉取省市区信息



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

GET https://api.weixin.qq.com/wxa/get_district?access_token=TOKEN
*/
func GetDistrict() (resp []byte, err error) {
	return offiaccount.HTTPGet(apiGetDistrict)
}

/*
在腾讯地图中搜索门店



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/search_map_poi?access_token=TOKEN
*/
func SearchMapPoi(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSearchMapPoi, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
在腾讯地图中创建门店



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/create_map_poi?access_token=TOKEN
*/
func CreateMapPoi(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiCreateMapPoi, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
添加门店



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/add_store?access_token=TOKEN
*/
func AddStore(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAddStore, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
更新门店信息



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/update_store?access_token=TOKEN
*/
func UpdateStore(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdateStore, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取门店小程序配置的卡券



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/card/storewxa/get?access_token=ACCESS_TOKEN
*/
func CardStorewxaGet(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiCardStorewxaGet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取单个门店信息



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/get_store_info?access_token=TOKEN
*/
func GetStoreInfo(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetStoreInfo, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取门店信息列表



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/get_store_list?access_token=TOKEN
*/
func GetStoreList(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetStoreList, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
删除门店



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/del_store?access_token=TOKEN
*/
func DelStore(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelStore, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
