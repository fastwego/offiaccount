package invoice

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiGetauthurl          = "/card/invoice/getauthurl"
	apiGetauthdata         = "/card/invoice/getauthdata"
	apiRejectinsert        = "/card/invoice/rejectinsert"
	apiMakeoutinvoice      = "/card/invoice/makeoutinvoice"
	apiClearoutinvoice     = "/card/invoice/clearoutinvoice"
	apiQueryinvoceinfo     = "/card/invoice/queryinvoceinfo"
	apiSeturl              = "/card/invoice/seturl"
	apiCreatecard          = "/card/invoice/platform/createcard"
	apiSetpdf              = "/card/invoice/platform/setpdf"
	apiGetpdf              = "/card/invoice/platform/getpdf"
	apiInsert              = "/card/invoice/insert"
	apiUpdatestatus        = "/card/invoice/platform/updatestatus"
	apiGetinvoiceinfo      = "/card/invoice/reimburse/getinvoiceinfo"
	apiGetinvoicebatch     = "/card/invoice/reimburse/getinvoicebatch"
	apiUpdateinvoicestatus = "/card/invoice/reimburse/updateinvoicestatus"
	apiUpdatestatusbatch   = "/card/invoice/reimburse/updatestatusbatch"
	apiGetusertitleurl     = "/card/invoice/biz/getusertitleurl"
	apiGetselecttitleurl   = "/card/invoice/getauthdata)card/invoice/biz/getselecttitleurl"
	apiGetselecttitleurl   = "/card/invoice/biz/getselecttitleurl"
	apiScantitle           = "/card/invoice/scantitle"
)

/*
/card/invoice/getauthurl



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html

POST https://api.weixin.qq.com/card/invoice/getauthurl?access_token={access_token}
*/
func Getauthurl(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetauthurl, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/getauthdata



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html

POST https://api.weixin.qq.com/card/invoice/getauthdata?access_token={access_token}
*/
func Getauthdata(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetauthdata, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/rejectinsert



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html

POST https://api.weixin.qq.com/card/invoice/rejectinsert?access_token={access_token}
*/
func Rejectinsert(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiRejectinsert, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/makeoutinvoice



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html

POST https://api.weixin.qq.com/card/invoice/makeoutinvoice?access_token={access_token}
*/
func Makeoutinvoice(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiMakeoutinvoice, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/clearoutinvoice



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html

POST https://api.weixin.qq.com/card/invoice/clearoutinvoice?access_token={access_token}
*/
func Clearoutinvoice(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiClearoutinvoice, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/queryinvoceinfo



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html

POST https://api.weixin.qq.com/card/invoice/queryinvoceinfo?access_token={access_token}
*/
func Queryinvoceinfo(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiQueryinvoceinfo, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/seturl



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html

POST https://api.weixin.qq.com/card/invoice/seturl?access_token={access_token}
*/
func Seturl(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSeturl, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/platform/createcard



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html

POST https://api.weixin.qq.com/card/invoice/platform/createcard?access_token={access_token}
*/
func Createcard(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiCreatecard, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/platform/setpdf



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html

POST https://api.weixin.qq.com/card/invoice/platform/setpdf?access_token={access_token}
*/
func Setpdf(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSetpdf, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/platform/getpdf



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html

POST https://api.weixin.qq.com/card/invoice/platform/getpdf?action=get_url&amp;access_token={access_token}
*/
func Getpdf(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetpdf, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/insert



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html

POST https://api.weixin.qq.com/card/invoice/insert?access_token={access_token}
*/
func Insert(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiInsert, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/platform/updatestatus



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html

POST https://api.weixin.qq.com/card/invoice/platform/updatestatus?access_token={access_token}
*/
func Updatestatus(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdatestatus, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/reimburse/getinvoiceinfo



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Reimburser_API_List.html

POST https://api.weixin.qq.com/card/invoice/reimburse/getinvoiceinfo?access_token={access_token}
*/
func Getinvoiceinfo(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetinvoiceinfo, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/reimburse/getinvoicebatch



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Reimburser_API_List.html

POST https://api.weixin.qq.com/card/invoice/reimburse/getinvoicebatch?access_token={access_token}
*/
func Getinvoicebatch(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetinvoicebatch, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/reimburse/updateinvoicestatus



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Reimburser_API_List.html

POST https://api.weixin.qq.com/card/invoice/reimburse/updateinvoicestatus?access_token={access_token}
*/
func Updateinvoicestatus(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdateinvoicestatus, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/reimburse/updatestatusbatch



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Reimburser_API_List.html

POST https://api.weixin.qq.com/card/invoice/reimburse/updatestatusbatch?access_token={access_token}
*/
func Updatestatusbatch(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdatestatusbatch, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/biz/getusertitleurl



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Quick_issuing/Interface_Instructions.html

POST https://api.weixin.qq.com/card/invoice/biz/getusertitleurl?access_token={access_token
*/
func Getusertitleurl(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetusertitleurl, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/getauthdata)card/invoice/biz/getselecttitleurl



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Quick_issuing/Interface_Instructions.html

POST https://api.weixin.qq.com/card/invoice/getauthdata)card/invoice/biz/getselecttitleurl?access_token={access_token/
*/
func Getselecttitleurl(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetselecttitleurl, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/biz/getselecttitleurl



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Quick_issuing/Interface_Instructions.html

POST https://api.weixin.qq.com/card/invoice/biz/getselecttitleurl?access_token={access_token}
*/
func Getselecttitleurl(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetselecttitleurl, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/scantitle



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Quick_issuing/Interface_Instructions.html

POST https://api.weixin.qq.com/card/invoice/scantitle?access_token={access_token}
*/
func Scantitle(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiScantitle, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
