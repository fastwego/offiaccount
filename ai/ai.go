package ai

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiSearch                 = "/semantic/semproxy/search"
	apiAddvoicetorecofortext  = "/cgi-bin/media/voice/addvoicetorecofortext"
	apiQueryrecoresultfortext = "/cgi-bin/media/voice/queryrecoresultfortext"
	apiTranslatecontent       = "/cgi-bin/media/voice/translatecontent"
	apiIdcard                 = "/cv/ocr/idcard"
	apiOcr                    = "/cv/ocr/"
	apiDrivinglicense         = "/cv/ocr/drivinglicense"
	apiBizlicense             = "/cv/ocr/bizlicense"
	apiComm                   = "/cv/ocr/comm"
	apiQrcode                 = "/cv/img/qrcode"
	apiSuperresolution        = "/cv/img/superresolution"
	apiAicrop                 = "/cv/img/aicrop"
)

/*
/semantic/semproxy/search



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Natural_Language_Processing.html

POST https://api.weixin.qq.com/semantic/semproxy/search?access_token=YOUR_ACCESS_TOKEN
*/
func Search(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSearch, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/cgi-bin/media/voice/addvoicetorecofortext



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/AI_Open_API.html

POST http://api.weixin.qq.com/cgi-bin/media/voice/addvoicetorecofortext?access_token=ACCESS_TOKEN&amp;format=&amp;voice_id=xxxxxx&amp;lang=zh_CN
*/
func Addvoicetorecofortext(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAddvoicetorecofortext, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/cgi-bin/media/voice/queryrecoresultfortext



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/AI_Open_API.html

POST http://api.weixin.qq.com/cgi-bin/media/voice/queryrecoresultfortext?access_token=ACCESS_TOKEN&amp;voice_id=xxxxxx&amp;lang=zh_CN
*/
func Queryrecoresultfortext(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiQueryrecoresultfortext, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/cgi-bin/media/voice/translatecontent



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/AI_Open_API.html

POST http://api.weixin.qq.com/cgi-bin/media/voice/translatecontent?access_token=ACCESS_TOKEN&amp;lfrom=xxx&amp;lto=xxx
*/
func Translatecontent(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiTranslatecontent, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/cv/ocr/idcard



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html

POST https://api.weixin.qq.com/cv/ocr/idcard?img_url=ENCODE_URL&amp;access_token=ACCESS_TOCKEN
*/
func Idcard(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiIdcard, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/cv/ocr/



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html

POST https://api.weixin.qq.com/cv/ocr/
*/
func Ocr(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiOcr, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/cv/ocr/drivinglicense



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html

POST https://api.weixin.qq.com/cv/ocr/drivinglicense?img_url=ENCODE_URL&amp;access_token=ACCESS_TOCKEN
*/
func Drivinglicense(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDrivinglicense, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/cv/ocr/bizlicense



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html

POST http://api.weixin.qq.com/cv/ocr/bizlicense?img_url=ENCODE_URL&amp;access_token=ACCESS_TOCKEN
*/
func Bizlicense(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiBizlicense, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/cv/ocr/comm



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html

POST http://api.weixin.qq.com/cv/ocr/comm?img_url=ENCODE_URL&amp;access_token=ACCESS_TOCKEN
*/
func Comm(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiComm, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/cv/img/qrcode



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Img_Proc.html

POST https://api.weixin.qq.com/cv/img/qrcode?img_url=ENCODE_URL&amp;access_token=ACCESS_TOCKEN
*/
func Qrcode(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiQrcode, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/cv/img/superresolution



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Img_Proc.html

POST https://api.weixin.qq.com/cv/img/superresolution?img_url=ENCODE_URL&amp;access_token=ACCESS_TOCKEN
*/
func Superresolution(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSuperresolution, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/cv/img/aicrop



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Img_Proc.html

POST http://api.weixin.qq.com/cv/img/aicrop?img_url=ENCODE_URL&amp;access_token=ACCESS_TOCKEN
*/
func Aicrop(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAicrop, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
