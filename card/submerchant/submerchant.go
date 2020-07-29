package submerchant

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiSubmerchantSubmit   = "/card/submerchant/submit"
	apiGetApplyProtocol    = "/card/getapplyprotocol"
	apiSubmerchantUpdate   = "/card/submerchant/update"
	apiSubmerchantGet      = "/card/submerchant/get"
	apiSubmerchantBatchget = "/card/submerchant/batchget"
)

/*
创建子商户

支持母商户调用该接口传入子商户的相关资料，并获取子商户ID，用于子商户的卡券功能管理。 子商户的资质包括：商户名称、商户logo（图片）、卡券类目、授权函（扫描件或彩照）、授权函有效期截止时间

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html

POST https://api.weixin.qq.com/card/submerchant/submit?access_token=TOKEN
*/
func SubmerchantSubmit(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSubmerchantSubmit, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
卡券开放类目查询

通过调用该接口查询卡券开放的类目ID，类目会随业务发展变更，请每次用接口去查询获取实时卡券类目

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html

POST https://api.weixin.qq.com/card/getapplyprotocol?access_token=TOKEN
*/
func GetApplyProtocol(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetApplyProtocol, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
更新子商户

支持调用该接口更新子商户信息

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html

POST https://api.weixin.qq.com/card/submerchant/update?access_token=TOKEN
*/
func SubmerchantUpdate(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSubmerchantUpdate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
拉取单个子商户信息

通过指定的子商户appid，拉取该子商户的基础信息。 注意，用母商户去调用接口，但接口内传入的是子商户的appid

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html

POST https://api.weixin.qq.com/card/submerchant/get?access_token=TOKEN
*/
func SubmerchantGet(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSubmerchantGet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
批量拉取子商户信息

母商户可以通过该接口批量拉取子商户的相关信息，一次调用最多拉取100个子商户的信息，可以通过多次拉去满足不同的查询需求

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html

POST https://api.weixin.qq.com/card/submerchant/batchget?access_token=TOKEN
*/
func SubmerchantBatchget(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSubmerchantBatchget, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
