package account

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiCreate   = "/cgi-bin/qrcode/create"
	apiShorturl = "/cgi-bin/shorturl"
)

/*
创建二维码ticket

每次创建二维码ticket需要提供一个开发者自行设定的参数

See: https://developers.weixin.qq.com/doc/offiaccount/Account_Management/Generating_a_Parametric_QR_Code.html

POST https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=TOKEN
*/
func Create(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiCreate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
长链接转短链接

将一条长链接转成短链接。主要使用场景： 开发者用于生成二维码的原链接（商品、支付二维码等）太长导致扫码速度和成功率下降，将原长链接通过此接口转成短链接再生成二维码将大大提升扫码速度和成功率

See: https://developers.weixin.qq.com/doc/offiaccount/Account_Management/URL_Shortener.html

POST https://api.weixin.qq.com/cgi-bin/shorturl?access_token=ACCESS_TOKEN
*/
func Shorturl(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiShorturl, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
