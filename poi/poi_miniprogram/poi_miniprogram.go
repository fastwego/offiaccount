package poi_miniprogram

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiGetwxastorecatelist  = "/wxa/getwxastorecatelist"
	apiGetMerchantCategory  = "/wxa/get_merchant_category"
	apiApplywxastore        = "/wxa/applywxastore "
	apiApplyMerchant        = "/wxa/apply_merchant"
	apiGetMerchantAuditInfo = "/wxa/get_merchant_audit_info"
	apiModifyMerchant       = "/wxa/modify_merchant"
	apiGetDistrict          = "/wxa/get_district"
	apiSearchMapPoi         = "/wxa/search_map_poi"
	apiCreateMapPoi         = "/wxa/create_map_poi"
	apiAddStore             = "/wxa/add_store"
	apiUpdateStore          = "/wxa/update_store"
	apiGet                  = "/card/storewxa/get"
	apiGetStoreInfo         = "/wxa/get_store_info"
	apiGetStoreList         = "/wxa/get_store_list"
	apiDelStore             = "/wxa/del_store"
)

/*
/wxa/getwxastorecatelist



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/getwxastorecatelist?access_token=TOKEN
*/
func Getwxastorecatelist(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetwxastorecatelist, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/wxa/get_merchant_category



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/get_merchant_category?access_token=TOKEN
*/
func GetMerchantCategory(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetMerchantCategory, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/wxa/applywxastore



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/applywxastore%20?access_token=TOKEN
*/
func Applywxastore(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiApplywxastore, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/wxa/apply_merchant



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/apply_merchant?access_token=TOKEN
*/
func ApplyMerchant(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiApplyMerchant, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/wxa/get_merchant_audit_info



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/get_merchant_audit_info?access_token=TOKEN
*/
func GetMerchantAuditInfo(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetMerchantAuditInfo, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/wxa/modify_merchant



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/modify_merchant?access_token=TOKEN
*/
func ModifyMerchant(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiModifyMerchant, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/wxa/get_district



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/get_district?access_token=TOKEN
*/
func GetDistrict(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetDistrict, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/wxa/search_map_poi



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/search_map_poi?access_token=TOKEN
*/
func SearchMapPoi(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSearchMapPoi, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/wxa/create_map_poi



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/create_map_poi?access_token=TOKEN
*/
func CreateMapPoi(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiCreateMapPoi, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/wxa/add_store



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/add_store?access_token=TOKEN
*/
func AddStore(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAddStore, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/wxa/update_store



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/update_store?access_token=TOKEN
*/
func UpdateStore(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdateStore, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/storewxa/get



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/card/storewxa/get?access_token=ACCESS_TOKEN
*/
func Get(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/wxa/get_store_info



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/get_store_info?access_token=TOKEN
*/
func GetStoreInfo(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetStoreInfo, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/wxa/get_store_list



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/get_store_list?access_token=TOKEN
*/
func GetStoreList(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetStoreList, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/wxa/del_store



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html

POST https://api.weixin.qq.com/wxa/del_store?access_token=TOKEN
*/
func DelStore(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelStore, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
