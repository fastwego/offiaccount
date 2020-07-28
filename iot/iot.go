package iot

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiApplycode         = "/intp/marketcode/applycode"
	apiApplycodequery    = "/intp/marketcode/applycodequery"
	apiApplycodedownload = "/intp/marketcode/applycodedownload"
	apiCodeactive        = "/intp/marketcode/codeactive"
	apiCodeactivequery   = "/intp/marketcode/codeactivequery"
	apiTickettocode      = "/intp/marketcode/tickettocode"
)

/*
/intp/marketcode/applycode



See: https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html

POST https://api.weixin.qq.com/intp/marketcode/applycode?access_token=ACCESSTOKEN
*/
func Applycode(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiApplycode, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/intp/marketcode/applycodequery



See: https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html

POST https://api.weixin.qq.com/intp/marketcode/applycodequery?access_token=ACCESSTOKEN
*/
func Applycodequery(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiApplycodequery, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/intp/marketcode/applycodedownload



See: https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html

POST https://api.weixin.qq.com/intp/marketcode/applycodedownload?access_token=ACCESSTOKE
*/
func Applycodedownload(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiApplycodedownload, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/intp/marketcode/codeactive



See: https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html

POST https://api.weixin.qq.com/intp/marketcode/codeactive?access_token=ACCESSTOKEN
*/
func Codeactive(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiCodeactive, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/intp/marketcode/codeactivequery



See: https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html

POST https://api.weixin.qq.com/intp/marketcode/codeactivequery?access_token=ACCESSTOKEN
*/
func Codeactivequery(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiCodeactivequery, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/intp/marketcode/tickettocode



See: https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html

POST https://api.weixin.qq.com/intp/marketcode/tickettocode?access_token=ACCESSTOKEN
*/
func Tickettocode(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiTickettocode, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
