package nontax

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiGetbillauthurl          = "/nontax/getbillauthurl"
	apiCreatebillcard          = "/nontax/createbillcard"
	apiInsertbill              = "/nontax/insertbill"
	apiQueryfee                = "/nontax/queryfee"
	apiUnifiedorder            = "/nontax/unifiedorder"
	apiGetorder                = "/nontax/getorder"
	apiRefund                  = "/nontax/refund"
	apiDownloadbill            = "/nontax/downloadbill"
	apiNotifyinconsistentorder = "/nontax/notifyinconsistentorder"
	apiMocknotification        = "/nontax/mocknotification"
	apiMockqueryfee            = "/nontax/mockqueryfee"
	apiMicropay                = "/nontax/micropay"
	apiGetorderlist            = "/nontax/getorderlist"
	apiGetauthurl              = "/intp/realname/getauthurl"
	apiGetrealname             = "/nontax/getrealname"
	apiQuerystate              = "/nontax/vehicle/querystate"
	apiEntrancenotify          = "/nontax/vehicle/entrancenotify"
	apiPayapply                = "/nontax/vehicle/payapply"
)

/*
/nontax/getbillauthurl



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Nontax_Bill/API_list.html

POST https://api.weixin.qq.com/nontax/getbillauthurl?access_token={access_token}
*/
func Getbillauthurl(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetbillauthurl, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/nontax/createbillcard



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Nontax_Bill/API_list.html

POST https://api.weixin.qq.com/nontax/createbillcard?access_token={access_token}
*/
func Createbillcard(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiCreatebillcard, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/nontax/insertbill



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Nontax_Bill/API_list.html

POST https://api.weixin.qq.com/nontax/insertbill?access_token={access_token}
*/
func Insertbill(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiInsertbill, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/nontax/queryfee



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/queryfee?access_token=$AccessToken
*/
func Queryfee(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiQueryfee, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/nontax/unifiedorder



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/unifiedorder?access_token=$AccessToken
*/
func Unifiedorder(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUnifiedorder, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/nontax/getorder



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/getorder?access_token=$AccessToken
*/
func Getorder(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetorder, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/nontax/refund



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/refund?access_token=$AccessToken
*/
func Refund(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiRefund, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/nontax/downloadbill



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/downloadbill?access_token=$AccessToken
*/
func Downloadbill(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDownloadbill, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/nontax/notifyinconsistentorder



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/notifyinconsistentorder?access_token=$AccessToken
*/
func Notifyinconsistentorder(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiNotifyinconsistentorder, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/nontax/mocknotification



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/mocknotification?access_token=$AccessToken
*/
func Mocknotification(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiMocknotification, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/nontax/mockqueryfee



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/mockqueryfee?access_token=$AccessToken
*/
func Mockqueryfee(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiMockqueryfee, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/nontax/micropay



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/micropay?access_token=$AccessToken
*/
func Micropay(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiMicropay, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/nontax/getorderlist



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/getorderlist?access_token=$AccessToken
*/
func Getorderlist(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetorderlist, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/intp/realname/getauthurl



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/intp/realname/getauthurl?access_token=ACCESS_TOKEN
*/
func Getauthurl(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetauthurl, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/nontax/getrealname



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html

POST https://api.weixin.qq.com/nontax/getrealname?access_token=ACCESS_TOKEN
*/
func Getrealname(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetrealname, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/nontax/vehicle/querystate



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/Non-tax_driver_platform.html

POST https://api.weixin.qq.com/nontax/vehicle/querystate?access_token=$AccessToken
*/
func Querystate(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiQuerystate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/nontax/vehicle/entrancenotify



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/Non-tax_driver_platform.html

POST https://api.weixin.qq.com/nontax/vehicle/entrancenotify?access_token=$AccessToken
*/
func Entrancenotify(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiEntrancenotify, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/nontax/vehicle/payapply



See: https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/Non-tax_driver_platform.html

POST https://api.weixin.qq.com/nontax/vehicle/payapply?access_token=$AccessToken
*/
func Payapply(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiPayapply, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
