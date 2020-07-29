package invoice

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiGetAuthUrl                   = "/card/invoice/getauthurl"
	apiGetAuthData                  = "/card/invoice/getauthdata"
	apiRejectInsert                 = "/card/invoice/rejectinsert"
	apiMakeoutinvoice               = "/card/invoice/makeoutinvoice"
	apiClearOutInvoice              = "/card/invoice/clearoutinvoice"
	apiQueryInvoceInfo              = "/card/invoice/queryinvoceinfo"
	apiSetUrl                       = "/card/invoice/seturl"
	apiPlatformCreateCard           = "/card/invoice/platform/createcard"
	apiPlatformSetpdf               = "/card/invoice/platform/setpdf"
	apiPlatformGetpdf               = "/card/invoice/platform/getpdf"
	apiInsert                       = "/card/invoice/insert"
	apiPlatformUpdateStatus         = "/card/invoice/platform/updatestatus"
	apiReimburseGetInvoiceInfo      = "/card/invoice/reimburse/getinvoiceinfo"
	apiReimburseGetInvoiceBatch     = "/card/invoice/reimburse/getinvoicebatch"
	apiReimburseUpdateInvoiceStatus = "/card/invoice/reimburse/updateinvoicestatus"
	apiReimburseUpdateStatusBatch   = "/card/invoice/reimburse/updatestatusbatch"
	apiGetUserTitleUrl              = "/card/invoice/biz/getusertitleurl"
	apiGetSelectTitleUrl            = "/card/invoice/biz/getselecttitleurl"
	apiScanTitle                    = "/card/invoice/scantitle"
)

/*
获取授权页链接

本接口供商户调用。商户通过本接口传入订单号、开票平台标识等参数，获取授权页的链接。在微信中向用户展示授权页，当用户点击了授权页上的“领取发票”/“申请开票”按钮后，即完成了订单号与该用户的授权关系绑定，后续开票平台可凭此订单号发起将发票卡券插入用户卡包的请求，微信也将据此授权关系校验是否放行插卡请求

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html

POST https://api.weixin.qq.com/card/invoice/getauthurl?access_token={access_token}
*/
func GetAuthUrl(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetAuthUrl, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询授权完成状态



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html

POST https://api.weixin.qq.com/card/invoice/getauthdata?access_token={access_token}
*/
func GetAuthData(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetAuthData, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
拒绝开票

户完成授权后，商户若发现用户提交信息错误、或者发生了退款时，可以调用该接口拒绝开票并告知用户。拒绝开票后，该订单无法向用户再次开票。已经拒绝开票的订单，无法再次使用，如果要重新开票，需使用新的order_id，获取授权链接，让用户再次授权

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html

POST https://api.weixin.qq.com/card/invoice/rejectinsert?access_token={access_token}
*/
func RejectInsert(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiRejectInsert, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
统一开票接口-开具蓝票

对于使用微信电子发票开票接入能力的商户，在公众号后台选择任何一家开票平台的套餐，都可以使用本接口实现电子发票的开具

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html

POST https://api.weixin.qq.com/card/invoice/makeoutinvoice?access_token={access_token}
*/
func Makeoutinvoice(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiMakeoutinvoice, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
统一开票接口-发票冲红

对于使用微信电子发票开票接入能力的商户，在公众号后台选择任何一家开票平台的套餐，都可以使用本接口实现电子发票的冲红

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html

POST https://api.weixin.qq.com/card/invoice/clearoutinvoice?access_token={access_token}
*/
func ClearOutInvoice(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiClearOutInvoice, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
统一开票接口-查询已开发票

对于使用微信电子发票开票接入能力的商户，在公众号后台选择任何一家开票平台的套餐，都可以使用本接口实现已开具电子发票的查询

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html

POST https://api.weixin.qq.com/card/invoice/queryinvoceinfo?access_token={access_token}
*/
func QueryInvoceInfo(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiQueryInvoceInfo, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取自身的开票平台识别码

开票平台可以通过此接口获得本开票平台的预开票url，进而获取s_pappid。开票平台将该s_pappid并透传给商户，商户可以通过该s_pappid参数在微信电子发票方案中标识出为自身提供开票服务的开票平台

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html

POST https://api.weixin.qq.com/card/invoice/seturl?access_token={access_token}
*/
func SetUrl(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSetUrl, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
创建发票卡券模板

通过本接口可以为创建一个商户的发票卡券模板，为该商户配置发票卡券模板上的自定义栏位。创建发票卡券模板生成的card_id将在创建发票卡券时被引用，故创建发票卡券模板是创建发票卡券的基础

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html

POST https://api.weixin.qq.com/card/invoice/platform/createcard?access_token={access_token}
*/
func PlatformCreateCard(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiPlatformCreateCard, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
上传PDF

商户或开票平台可以通过该接口上传PDF。PDF上传成功后将获得发票文件的标识，后续可以通过插卡接口将PDF关联到用户的发票卡券上，一并插入到收票用户的卡包中

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html

POST https://api.weixin.qq.com/card/invoice/platform/setpdf?access_token={access_token}
*/
func PlatformSetpdf(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiPlatformSetpdf, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询已上传的PDF文件

用于供发票PDF的上传方查询已经上传的发票或消费凭证PDF

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html

POST https://api.weixin.qq.com/card/invoice/platform/getpdf?action=get_url&access_token={access_token}
*/
func PlatformGetpdf(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiPlatformGetpdf, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
将电子发票卡券插入用户卡包

本接口由开票平台或自建平台商户调用。对用户已经授权过的开票请求，开票平台可以使用本接口将发票制成发票卡券放入用户的微信卡包中

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html

POST https://api.weixin.qq.com/card/invoice/insert?access_token={access_token}
*/
func Insert(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiInsert, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
更新发票卡券状态



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html

POST https://api.weixin.qq.com/card/invoice/platform/updatestatus?access_token={access_token}
*/
func PlatformUpdateStatus(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiPlatformUpdateStatus, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询报销发票信息

通过该接口查询电子发票的结构化信息，并获取发票PDF文件

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Reimburser_API_List.html

POST https://api.weixin.qq.com/card/invoice/reimburse/getinvoiceinfo?access_token={access_token}
*/
func ReimburseGetInvoiceInfo(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiReimburseGetInvoiceInfo, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
批量查询报销发票信息



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Reimburser_API_List.html

POST https://api.weixin.qq.com/card/invoice/reimburse/getinvoicebatch?access_token={access_token}
*/
func ReimburseGetInvoiceBatch(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiReimburseGetInvoiceBatch, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
报销方更新发票状态



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Reimburser_API_List.html

POST https://api.weixin.qq.com/card/invoice/reimburse/updateinvoicestatus?access_token={access_token}
*/
func ReimburseUpdateInvoiceStatus(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiReimburseUpdateInvoiceStatus, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
报销方批量更新发票状态



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Reimburser_API_List.html

POST https://api.weixin.qq.com/card/invoice/reimburse/updatestatusbatch?access_token={access_token}
*/
func ReimburseUpdateStatusBatch(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiReimburseUpdateStatusBatch, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
将发票抬头信息录入到用户微信中

调用接口，获取添加存储发票抬头信息的链接，将链接发给微信用户，用户确认后将保存该信息

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Quick_issuing/Interface_Instructions.html

POST https://api.weixin.qq.com/card/invoice/biz/getusertitleurl?access_token={access_token
*/
func GetUserTitleUrl(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetUserTitleUrl, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取用户抬头（方式一）:获取商户专属二维码，立在收银台

商户调用接口，获取链接，将链接转成二维码，用户扫码，可以选择抬头发给商户

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Quick_issuing/Interface_Instructions.html

POST https://api.weixin.qq.com/card/invoice/biz/getselecttitleurl?access_token={access_token}
*/
func GetSelectTitleUrl(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetSelectTitleUrl, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取用户抬头（方式二）:商户扫描用户的发票抬头二维码

商户扫用户“我的—个人信息—我的发票抬头”里面的抬头二维码后，通过调用本接口，可以获取用户抬头信息

See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Quick_issuing/Interface_Instructions.html

POST https://api.weixin.qq.com/card/invoice/scantitle?access_token={access_token}
*/
func ScanTitle(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiScanTitle, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
