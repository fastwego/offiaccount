package ai

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiSemantic               = "/semantic/semproxy/search"
	apiAddVoiceToRecoForText  = "/cgi-bin/media/voice/addvoicetorecofortext"
	apiQueryRecoResultForText = "/cgi-bin/media/voice/queryrecoresultfortext"
	apiTranslateContent       = "/cgi-bin/media/voice/translatecontent"
	apiIDCard                 = "/cv/ocr/idcard"
	apiBankcard               = "/cv/ocr/bankcard"
	apiDrivingLicense         = "/cv/ocr/drivinglicense"
	apiBizLicense             = "/cv/ocr/bizlicense"
	apiComm                   = "/cv/ocr/comm"
	apiQRCode                 = "/cv/img/qrcode"
	apiSuperResolution        = "/cv/img/superresolution"
	apiAICrop                 = "/cv/img/aicrop"
)

/*
语义理解



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Natural_Language_Processing.html

POST https://api.weixin.qq.com/semantic/semproxy/search?access_token=YOUR_ACCESS_TOKEN
*/
func Semantic(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSemantic, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
提交语音


See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/AI_Open_API.html

POST http://api.weixin.qq.com/cgi-bin/media/voice/addvoicetorecofortext?access_token=ACCESS_TOKEN&format=&voice_id=xxxxxx&lang=zh_CN
*/
func AddVoiceToRecoForText(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAddVoiceToRecoForText, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取语音识别结果



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/AI_Open_API.html

POST http://api.weixin.qq.com/cgi-bin/media/voice/queryrecoresultfortext?access_token=ACCESS_TOKEN&voice_id=xxxxxx&lang=zh_CN
*/
func QueryRecoResultForText(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiQueryRecoResultForText, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
微信翻译



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/AI_Open_API.html

POST http://api.weixin.qq.com/cgi-bin/media/voice/translatecontent?access_token=ACCESS_TOKEN&lfrom=xxx&lto=xxx
*/
func TranslateContent(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiTranslateContent, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
身份证OCR识别



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html

POST https://api.weixin.qq.com/cv/ocr/idcard?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func IDCard(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiIDCard, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
银行卡OCR识别



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html

POST https://api.weixin.qq.com/cv/ocr/bankcard
*/
func Bankcard(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiBankcard, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
行驶证/驾驶证 OCR识别



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html

POST https://api.weixin.qq.com/cv/ocr/drivinglicense?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func DrivingLicense(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDrivingLicense, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
营业执照OCR识别



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html

POST http://api.weixin.qq.com/cv/ocr/bizlicense?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func BizLicense(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiBizLicense, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
通用印刷体OCR识别



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html

POST http://api.weixin.qq.com/cv/ocr/comm?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func Comm(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiComm, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
二维码/条码识别



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Img_Proc.html

POST https://api.weixin.qq.com/cv/img/qrcode?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func QRCode(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiQRCode, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
图片高清化



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Img_Proc.html

POST https://api.weixin.qq.com/cv/img/superresolution?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func SuperResolution(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSuperResolution, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
图片智能裁剪



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Img_Proc.html

POST http://api.weixin.qq.com/cv/img/aicrop?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func AICrop(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAICrop, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
