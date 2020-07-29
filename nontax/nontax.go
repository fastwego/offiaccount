package nontax

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiGetBillAuthUrl          = "/nontax/getbillauthurl"
	apiCreateBillCard          = "/nontax/createbillcard"
	apiInsertBill              = "/nontax/insertbill"
	apiQueryFee                = "/nontax/queryfee"
	apiUnifiedOrder            = "/nontax/unifiedorder"
	apiGetorder                = "/nontax/getorder"
	apiRefund                  = "/nontax/refund"
	apiDownloadBill            = "/nontax/downloadbill"
	apiNotifyInconsistentOrder = "/nontax/notifyinconsistentorder"
	apiMockNotification        = "/nontax/mocknotification"
	apiMockQueryFee            = "/nontax/mockqueryfee"
	apiMicropay                = "/nontax/micropay"
	apiGetOrderList            = "/nontax/getorderlist"
	apiRealNameGetAuthUrl      = "/intp/realname/getauthurl"
	apiGetRealName             = "/nontax/getrealname"
)

/*
获取授权页链接

通过此接口，获取授权页链接，让用户跳转到授权页

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Nontax_Bill/API_list.html

POST https://api.weixin.qq.com/nontax/getbillauthurl?access_token={access_token}
*/
func GetBillAuthUrl(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetBillAuthUrl, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
创建财政电子票据接口

财政局可以通过这个接口帮助执收单位创建一张财政电子票据模板。同一个财政局可以对应多个执收单位，同一个执收单位，使用同一个card_id，不同的执收单位，使用不同的card_id

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Nontax_Bill/API_list.html

POST https://api.weixin.qq.com/nontax/createbillcard?access_token={access_token}
*/
func CreateBillCard(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiCreateBillCard, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
将财政电子票据添加到用户微信卡包

执收单位完成用户插卡授权后，向财政局请求给某一个订单号进行领取财政电子票据，财政局须调用该接口对用户进行开票

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Nontax_Bill/API_list.html

POST https://api.weixin.qq.com/nontax/insertbill?access_token={access_token}
*/
func InsertBill(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiInsertBill, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询应收信息（提供给委办局）



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/queryfee?access_token=$AccessToken
*/
func QueryFee(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiQueryFee, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
支付下单（提供给委办局）



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/unifiedorder?access_token=$AccessToken
*/
func UnifiedOrder(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUnifiedOrder, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询订单（提供给委办局、银行、财政）



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/getorder?access_token=$AccessToken
*/
func Getorder(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetorder, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
申请退款（提供给银行）



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/refund?access_token=$AccessToken
*/
func Refund(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiRefund, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
下载对帐单（提供给银行）



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/downloadbill?access_token=$AccessToken
*/
func DownloadBill(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDownloadBill, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
通知不一致订单（提供给财政）



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/notifyinconsistentorder?access_token=$AccessToken
*/
func NotifyInconsistentOrder(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiNotifyInconsistentOrder, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
测试支付结果通知

接口提供给接入方在联调前自己调试支付通知接口的加解密和其它基本逻辑。 若version为1，会通知两次，分别测试加解密是否正确和是否有验签。 调用此接口会触发微信后台向 url 发送支付结果通知的测试数据

See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/mocknotification?access_token=$AccessToken
*/
func MockNotification(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiMockNotification, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
测试查询应收信息

此接口提供给接入方在联调前自己调试查询应收信息接口的加解密和其它基本逻辑。 若version为1，会通知两次，分别测试加解密是否正确和是否有验签。 调用此接口会触发微信后台向 url 发送查询应收信息的测试数据

See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/mockqueryfee?access_token=$AccessToken
*/
func MockQueryFee(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiMockQueryFee, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
提交刷卡支付

提交支付请求后微信会同步返回支付结果。 接口返回系统失败时，等待5秒重新调用看返回码。 当结果返回用户支付中需要输入密码时，可每间隔一段时间(建议10秒)重新调用该接口，直到有明确成功、失败，或者超时（建议30秒）

See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/micropay?access_token=$AccessToken
*/
func Micropay(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiMicropay, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询订单列表



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/getorderlist?access_token=$AccessToken
*/
func GetOrderList(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetOrderList, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取用户实名信息-获取授权链接



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/intp/realname/getauthurl?access_token=ACCESS_TOKEN
*/
func RealNameGetAuthUrl(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiRealNameGetAuthUrl, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取实名信息



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/getrealname?access_token=ACCESS_TOKEN
*/
func GetRealName(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetRealName, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
