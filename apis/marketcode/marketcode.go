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

// Package marketcode 一物一码
package marketcode

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
func ApplyCode(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiApplyCode, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询二维码申请单



See: https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html

POST https://api.weixin.qq.com/intp/marketcode/applycodequery?access_token=ACCESSTOKEN
*/
func ApplyCodeQuery(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiApplyCodeQuery, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
下载二维码包



See: https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html

POST https://api.weixin.qq.com/intp/marketcode/applycodedownload?access_token=ACCESSTOKE
*/
func ApplyCodeDownload(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiApplyCodeDownload, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
激活二维码



See: https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html

POST https://api.weixin.qq.com/intp/marketcode/codeactive?access_token=ACCESSTOKEN
*/
func CodeActive(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCodeActive, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询二维码激活状态



See: https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html

POST https://api.weixin.qq.com/intp/marketcode/codeactivequery?access_token=ACCESSTOKEN
*/
func CodeActiveQuery(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCodeActiveQuery, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
code_ticket换code



See: https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html

POST https://api.weixin.qq.com/intp/marketcode/tickettocode?access_token=ACCESSTOKEN
*/
func TicketToCode(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiTicketToCode, bytes.NewReader(payload), "application/json;charset=utf-8")
}
