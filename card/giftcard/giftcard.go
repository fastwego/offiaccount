package giftcard

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiGiftCardPageAdd         = "/card/giftcard/page/add"
	apiGiftCardPageGet         = "/card/giftcard/page/get"
	apiGiftCardPageUpdate      = "/card/giftcard/page/update"
	apiGiftCardPageBatchGet    = "/card/giftcard/page/batchget"
	apiSet                     = "/card/giftcard/maintain/set"
	apiGiftCardPayWhitelistAdd = "/card/giftcard/pay/whitelist/add"
	apiGiftCardPaySubmchBind   = "/card/giftcard/pay/submch/bind"
	apiGiftCardWxaSet          = "/card/giftcard/wxa/set"
	apiGiftCardOrderGet        = "/card/giftcard/order/get"
	apiGiftCardOrderBatchGet   = "/card/giftcard/order/batchget"
	apiGeneralCardUpdateUser   = "/card/generalcard/updateuser"
	apiGiftCardOrderRefund     = "/card/giftcard/order/refund"
	apiInvoiceSetBizAttr       = "/card/invoice/setbizattr"
	apiInvoiceGetAuthData      = "/card/invoice/getauthdata"
)

/*
创建-礼品卡货架

开发者可以通过该接口创建一个礼品卡货架并且用于公众号、门店的礼品卡售卖

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/page/add?access_token=ACCESS_TOKEN
*/
func GiftCardPageAdd(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGiftCardPageAdd, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询-礼品卡货架信息

开发者可以查询某个礼品卡货架信息

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/page/get?access_token=ACCESS_TOKEN
*/
func GiftCardPageGet(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGiftCardPageGet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
修改-礼品卡货架信息

开发者可以通过该接口更新礼品卡货架信息

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/page/update?access_token=ACCESS_TOKEN
*/
func GiftCardPageUpdate(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGiftCardPageUpdate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询-礼品卡货架列表

开发者可以通过该接口查询当前商户下所有的礼品卡货架id

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/page/batchget?access_token=ACCESS_TOKEN
*/
func GiftCardPageBatchGet(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGiftCardPageBatchGet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
下架-礼品卡货架

开发者可以通过该接口查询当前商户下所有的礼品卡货架id

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/maintain/set?access_token=ACCESS_TOKEN
*/
func Set(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
申请微信支付礼品卡权限



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/pay/whitelist/add?access_token=TOKEN
*/
func GiftCardPayWhitelistAdd(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGiftCardPayWhitelistAdd, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
 绑定商户号到礼品卡小程序



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/pay/submch/bind?access_token=TOKEN
*/
func GiftCardPaySubmchBind(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGiftCardPaySubmchBind, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
上传小程序代码



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/wxa/set
*/
func GiftCardWxaSet(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGiftCardWxaSet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询-单个礼品卡订单信息

开发者可以通过该接口查询某个订单号对应的订单详情

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/order/get?access_token=ACCESS_TOKEN
*/
func GiftCardOrderGet(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGiftCardOrderGet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
批量查询礼品卡订单信息

开发者可以通过该接口查询该商户某个时间段内创建的所有订单详情

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/order/batchget?access_token=ACCESS_TOKEN
*/
func GiftCardOrderBatchGet(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGiftCardOrderBatchGet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
更新用户礼品卡信息

当礼品卡被使用后，开发者可以通过该接口变更某个礼品卡的余额信息

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/generalcard/updateuser?access_token=TOKEN
*/
func GeneralCardUpdateUser(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGeneralCardUpdateUser, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
退款

开发者可以通过该接口对某一笔订单操作退款，注意该接口比较隐私，请开发者提高操作该功能的权限等级

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/order/refund?access_token=ACCESS_TOKEN
*/
func GiftCardOrderRefund(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGiftCardOrderRefund, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
设置支付后开票信息

商户可以通过该接口设置某个商户号发生收款后在支付消息上出现开票授权按钮

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/invoice/setbizattr?action=set_pay_mch&access_token={access_token}
*/
func InvoiceSetBizAttr(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiInvoiceSetBizAttr, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询开票信息

用户完成授权后，商户可以调用该接口查询某一个订单

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/invoice/getauthdata
*/
func InvoiceGetAuthData(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiInvoiceGetAuthData, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
