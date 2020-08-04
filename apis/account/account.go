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

// Package account 账号管理
package account

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiCreateQRCode = "/cgi-bin/qrcode/create"
	apiShortUrl     = "/cgi-bin/shorturl"
)

/*
创建二维码ticket

每次创建二维码ticket需要提供一个开发者自行设定的参数

See: https://developers.weixin.qq.com/doc/offiaccount/Account_Management/Generating_a_Parametric_QR_Code.html

POST https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=TOKEN
*/
func CreateQRCode(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreateQRCode, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
长链接转短链接

将一条长链接转成短链接。主要使用场景： 开发者用于生成二维码的原链接（商品、支付二维码等）太长导致扫码速度和成功率下降，将原长链接通过此接口转成短链接再生成二维码将大大提升扫码速度和成功率

See: https://developers.weixin.qq.com/doc/offiaccount/Account_Management/URL_Shortener.html

POST https://api.weixin.qq.com/cgi-bin/shorturl?access_token=ACCESS_TOKEN
*/
func ShortUrl(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiShortUrl, bytes.NewReader(payload), "application/json;charset=utf-8")
}
