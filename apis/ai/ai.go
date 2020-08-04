// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package ai 智能接口
package ai

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path"

	"github.com/fastwego/offiaccount"
)

const (
	apiSemantic               = "/semantic/semproxy/search"
	apiAddVoiceToRecoForText  = "/cgi-bin/media/voice/addvoicetorecofortext"
	apiQueryRecoResultForText = "/cgi-bin/media/voice/queryrecoresultfortext"
	apiTranslateContent       = "/cgi-bin/media/voice/translatecontent"
	apiOCRIDCard              = "/cv/ocr/idcard"
	apiOCRBankcard            = "/cv/ocr/bankcard"
	apiOCRDrivingLicense      = "/cv/ocr/drivinglicense"
	apiOCRBizLicense          = "/cv/ocr/bizlicense"
	apiOCRCommon              = "/cv/ocr/comm"
	apiQRCode                 = "/cv/img/qrcode"
	apiSuperResolution        = "/cv/img/superresolution"
	apiAICrop                 = "/cv/img/aicrop"
)

/*
语义理解



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Natural_Language_Processing.html

POST https://api.weixin.qq.com/semantic/semproxy/search?access_token=YOUR_ACCESS_TOKEN
*/
func Semantic(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSemantic, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
提交语音



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/AI_Open_API.html

POST(@media) https://api.weixin.qq.com/cgi-bin/media/voice/addvoicetorecofortext?access_token=ACCESS_TOKEN&format=&voice_id=xxxxxx&lang=zh_CN
*/
func AddVoiceToRecoForText(ctx *offiaccount.OffiAccount, media string, params url.Values) (resp []byte, err error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("media", path.Base(media))
		if err != nil {
			return
		}
		file, err := os.Open(media)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}

	}()
	return ctx.Client.HTTPPost(apiAddVoiceToRecoForText+"?"+params.Encode(), r, m.FormDataContentType())
}

/*
获取语音识别结果



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/AI_Open_API.html

POST https://api.weixin.qq.com/cgi-bin/media/voice/queryrecoresultfortext?access_token=ACCESS_TOKEN&voice_id=xxxxxx&lang=zh_CN
*/
func QueryRecoResultForText(ctx *offiaccount.OffiAccount, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiQueryRecoResultForText+"?"+params.Encode(), bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
微信翻译



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/AI_Open_API.html

POST https://api.weixin.qq.com/cgi-bin/media/voice/translatecontent?access_token=ACCESS_TOKEN&lfrom=xxx&lto=xxx
*/
func TranslateContent(ctx *offiaccount.OffiAccount, payload []byte, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiTranslateContent+"?"+params.Encode(), bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
身份证OCR识别



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html

POST https://api.weixin.qq.com/cv/ocr/idcard?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func OCRIDCard(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiOCRIDCard, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
银行卡OCR识别



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html

POST https://api.weixin.qq.com/cv/ocr/bankcard
*/
func OCRBankcard(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiOCRBankcard, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
行驶证/驾驶证 OCR识别



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html

POST https://api.weixin.qq.com/cv/ocr/drivinglicense?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func OCRDrivingLicense(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiOCRDrivingLicense, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
营业执照OCR识别



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html

POST https://api.weixin.qq.com/cv/ocr/bizlicense?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func OCRBizLicense(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiOCRBizLicense, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
通用印刷体OCR识别



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html

POST https://api.weixin.qq.com/cv/ocr/comm?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func OCRCommon(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiOCRCommon, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
二维码/条码识别



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Img_Proc.html

POST https://api.weixin.qq.com/cv/img/qrcode?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func QRCode(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiQRCode, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
图片高清化



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Img_Proc.html

POST https://api.weixin.qq.com/cv/img/superresolution?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func SuperResolution(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSuperResolution, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
图片智能裁剪



See: https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Img_Proc.html

POST https://api.weixin.qq.com/cv/img/aicrop?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
*/
func AICrop(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAICrop, bytes.NewReader(payload), "application/json;charset=utf-8")
}
