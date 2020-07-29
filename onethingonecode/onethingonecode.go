package onethingonecode

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiApplyCode         = "/intp/marketcode/applycode"
	apiApplyCodeQuery    = "/intp/marketcode/applycodequery"
	apiApplyCodeDownload = "/intp/marketcode/applycodedownload"
	apiCodeActive        = "/intp/marketcode/codeactive"
	apiCodeActiveQuery   = "/intp/marketcode/codeactivequery"
	apiTicketToCode      = "/intp/marketcode/tickettocode"
)

/*
申请二维码


See: https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html

POST https://api.weixin.qq.com/intp/marketcode/applycode?access_token=ACCESSTOKEN
*/
func ApplyCode(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiApplyCode, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询二维码申请单



See: https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html

POST https://api.weixin.qq.com/intp/marketcode/applycodequery?access_token=ACCESSTOKEN
*/
func ApplyCodeQuery(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiApplyCodeQuery, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
下载二维码包



See: https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html

POST https://api.weixin.qq.com/intp/marketcode/applycodedownload?access_token=ACCESSTOKE
*/
func ApplyCodeDownload(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiApplyCodeDownload, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
激活二维码



See: https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html

POST https://api.weixin.qq.com/intp/marketcode/codeactive?access_token=ACCESSTOKEN
*/
func CodeActive(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiCodeActive, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询二维码激活状态



See: https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html

POST https://api.weixin.qq.com/intp/marketcode/codeactivequery?access_token=ACCESSTOKEN
*/
func CodeActiveQuery(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiCodeActiveQuery, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
code_ticket换code



See: https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html

POST https://api.weixin.qq.com/intp/marketcode/tickettocode?access_token=ACCESSTOKEN
*/
func TicketToCode(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiTicketToCode, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
