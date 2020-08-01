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

// Package tag 标签管理
package tag

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiNewGuideTagOption       = "/cgi-bin/guide/newguidetagoption"
	apiDelguidetagoption       = "/cgi-bin/guide/delguidetagoption"
	apiAddGuideTagOption       = "/cgi-bin/guide/addguidetagoption"
	apiGetGuideTagOption       = "/cgi-bin/guide/getguidetagoption"
	apiAddGuideBuyerTag        = "/cgi-bin/guide/addguidebuyertag"
	apiGetGuideBuyerTag        = "/cgi-bin/guide/getguidebuyertag"
	apiQueryGuideBuyerByTag    = "/cgi-bin/guide/queryguidebuyerbytag"
	apiDelGuideBuyerTag        = "/cgi-bin/guide/delguidebuyertag"
	apiAddGuideBuyerDisplayTag = "/cgi-bin/guide/addguidebuyerdisplaytag"
	apiGetGuideBuyerDisplayTag = "/cgi-bin/guide/getguidebuyerdisplaytag"
)

/*
/cgi-bin/guide/newguidetagoption



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.newGuideTagOption.html

POST https://api.weixin.qq.com/cgi-bin/guide/newguidetagoption?access_token=ACCESS_TOKEN
*/
func NewGuideTagOption(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiNewGuideTagOption, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
删除指定标签类型



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.delguidetagoption.html

POST https://api.weixin.qq.com/cgi-bin/guide/delguidetagoption?access_token=ACCESS_TOKEN
*/
func Delguidetagoption(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelguidetagoption, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
为标签添加可选值



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.addGuideTagOption.html

POST https://api.weixin.qq.com/cgi-bin/guide/addguidetagoption?access_token=ACCESS_TOKEN
*/
func AddGuideTagOption(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddGuideTagOption, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取标签和可选值



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.getGuideTagOption.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidetagoption?access_token=ACCESS_TOKEN
*/
func GetGuideTagOption(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGuideTagOption, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
为客户设置标签



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.addGuideBuyerTag.html

POST https://api.weixin.qq.com/cgi-bin/guide/addguidebuyertag?access_token=ACCESS_TOKEN
*/
func AddGuideBuyerTag(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddGuideBuyerTag, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
查询客户标签



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.getGuideBuyerTag.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidebuyertag?access_token=ACCESS_TOKEN
*/
func GetGuideBuyerTag(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGuideBuyerTag, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
根据标签值筛选客户



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.queryGuideBuyerByTag.html

POST https://api.weixin.qq.com/cgi-bin/guide/queryguidebuyerbytag?access_token=ACCESS_TOKEN
*/
func QueryGuideBuyerByTag(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiQueryGuideBuyerByTag, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
删除客户标签



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.delGuideBuyerTag.html

POST https://api.weixin.qq.com/cgi-bin/guide/delguidebuyertag?access_token=ACCESS_TOKEN
*/
func DelGuideBuyerTag(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelGuideBuyerTag, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
设置自定义客户信息



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.addGuideBuyerDisplayTag.html

POST https://api.weixin.qq.com/cgi-bin/guide/addguidebuyerdisplaytag?access_token=ACCESS_TOKEN
*/
func AddGuideBuyerDisplayTag(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddGuideBuyerDisplayTag, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取自定义客户信息



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.getGuideBuyerDisplayTag.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidebuyerdisplaytag?access_token=ACCESS_TOKEN
*/
func GetGuideBuyerDisplayTag(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGuideBuyerDisplayTag, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}
