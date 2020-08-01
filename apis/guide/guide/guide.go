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

// Package guide 顾问管理
package guide

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiGetGuideAcct              = "/cgi-bin/guide/getguideacct"
	apiAddGuideAcct              = "/cgi-bin/guide/addguideacct"
	apiUpdateGuideAcct           = "/cgi-bin/guide/updateguideacct"
	apiDelGuideAcct              = "/cgi-bin/guide/delguideacct"
	apiGetGuideAcctList          = "/cgi-bin/guide/getguideacctlist"
	apiGuideCreateQrCode         = "/cgi-bin/guide/guidecreateqrcode"
	apiGetGuideBuyerChatRecord   = "/cgi-bin/guide/getguidebuyerchatrecord"
	apiSetGuideConfig            = "/cgi-bin/guide/setguideconfig"
	apiGetGuideConfig            = "/cgi-bin/guide/getguideconfig"
	apiSetGuideAcctConfig        = "/cgi-bin/guide/setguideacctconfig"
	apiGetGuideAcctConfig        = "/cgi-bin/guide/getguideacctconfig"
	apiPushShowWxaPathMenu       = "/cgi-bin/guide/pushshowwxapathmenu"
	apiNewGuideGroup             = "/cgi-bin/guide/newguidegroup"
	apiGetGuideGroupList         = "/cgi-bin/guide/getguidegrouplist"
	apiGetGroupInfo              = "/cgi-bin/guide/getgroupinfo"
	apiAddGuide2GuideGroup       = "/cgi-bin/guide/addguide2guidegroup"
	apiDelGuide2GuideGroup       = "/cgi-bin/guide/delguide2guidegroup"
	apiGetGroupByGuide           = "/cgi-bin/guide/getgroupbyguide"
	apiDelGuideGroup             = "/cgi-bin/guide/delguidegroup"
	apiGetGuideBuyerRelationList = "/cgi-bin/guide/getguidebuyerrelationlist"
)

/*
获取顾问信息



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideAcct.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguideacct?access_token=ACCESS_TOKEN
*/
func GetGuideAcct(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGuideAcct, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
为服务号添加顾问



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.addGuideAcct.html

POST https://api.weixin.qq.com/cgi-bin/guide/addguideacct?access_token=ACCESS_TOKEN
*/
func AddGuideAcct(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddGuideAcct, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
修改顾问的昵称或头像



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.updateGuideAcct.html

POST https://api.weixin.qq.com/cgi-bin/guide/updateguideacct?access_token=ACCESS_TOKEN
*/
func UpdateGuideAcct(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdateGuideAcct, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
删除顾问



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.delGuideAcct.html

POST https://api.weixin.qq.com/cgi-bin/guide/delguideacct?access_token=ACCESS_TOKEN
*/
func DelGuideAcct(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelGuideAcct, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取服务号顾问列表



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideAcctList.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguideacctlist?access_token=ACCESS_TOKEN
*/
func GetGuideAcctList(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGuideAcctList, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
生成顾问二维码



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.guideCreateQrCode.html

POST https://api.weixin.qq.com/cgi-bin/guide/guidecreateqrcode?access_token=ACCESS_TOKEN
*/
func GuideCreateQrCode(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGuideCreateQrCode, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
/cgi-bin/guide/getguidebuyerchatrecord



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideBuyerChatRecord.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidebuyerchatrecord?access_token=ACCESS_TOKEN
*/
func GetGuideBuyerChatRecord(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGuideBuyerChatRecord, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
设置快捷回复与关注自动回复



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.setGuideConfig.html

POST https://api.weixin.qq.com/cgi-bin/guide/setguideconfig?access_token=ACCESS_TOKEN
*/
func SetGuideConfig(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetGuideConfig, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取快捷回复与关注自动回复



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideConfig.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguideconfig?access_token=ACCESS_TOKEN
*/
func GetGuideConfig(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGuideConfig, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
为服务号设置敏感词与离线自动回复



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.setGuideAcctConfig.html

POST https://api.weixin.qq.com/cgi-bin/guide/setguideacctconfig?access_token=ACCESS_TOKEN
*/
func SetGuideAcctConfig(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetGuideAcctConfig, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取离线自动回复与敏感词



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideAcctConfig.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguideacctconfig?access_token=ACCESS_TOKEN
*/
func GetGuideAcctConfig(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGuideAcctConfig, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
允许微信用户复制小程序页面路径



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.pushShowWxaPathMenu.html

POST https://api.weixin.qq.com/cgi-bin/guide/pushshowwxapathmenu?access_token=ACCESS_TOKEN
*/
func PushShowWxaPathMenu(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiPushShowWxaPathMenu, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
新建顾问分组



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.newGuideGroup.html

POST https://api.weixin.qq.com/cgi-bin/guide/newguidegroup?access_token=ACCESS_TOKEN
*/
func NewGuideGroup(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiNewGuideGroup, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取服务号下所有顾问分组的列表



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideGroupList.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidegrouplist?access_token=ACCESS_TOKEN
*/
func GetGuideGroupList(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGuideGroupList, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取指定顾问分组信息，以及分组内顾问信息



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGroupInfo.html

POST https://api.weixin.qq.com/cgi-bin/guide/getgroupinfo?access_token=ACCESS_TOKEN
*/
func GetGroupInfo(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGroupInfo, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
分组内添加顾问



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.addGuide2GuideGroup.html

POST https://api.weixin.qq.com/cgi-bin/guide/addguide2guidegroup?access_token=ACCESS_TOKEN
*/
func AddGuide2GuideGroup(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddGuide2GuideGroup, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
分组内删除顾问



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.delGuide2GuideGroup.html

POST https://api.weixin.qq.com/cgi-bin/guide/delguide2guidegroup?access_token=ACCESS_TOKEN
*/
func DelGuide2GuideGroup(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelGuide2GuideGroup, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取顾问所在分组



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGroupByGuide.html

POST https://api.weixin.qq.com/cgi-bin/guide/getgroupbyguide?access_token=ACCESS_TOKEN
*/
func GetGroupByGuide(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGroupByGuide, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
删除指定顾问分组



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.delGuideGroup.html

POST https://api.weixin.qq.com/cgi-bin/guide/delguidegroup?access_token=ACCESS_TOKEN
*/
func DelGuideGroup(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelGuideGroup, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取顾问的客户列表



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.getGuideBuyerRelationList.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidebuyerrelationlist?access_token=ACCESS_TOKEN
*/
func GetGuideBuyerRelationList(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGuideBuyerRelationList, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}
