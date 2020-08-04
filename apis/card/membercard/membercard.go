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

// Package membercard 会员卡专区
package membercard

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiActivateGetUrl      = "/card/membercard/activate/geturl"
	apiActivateTempInfoGet = "/card/membercard/activatetempinfo/get"
	apiActivate            = "/card/membercard/activate"
	apiActivateUserFormSet = "/card/membercard/activateuserform/set"
	apiUserinfoGet         = "/card/membercard/userinfo/get"
	apiUpdateUser          = "/card/membercard/updateuser"
	apiPayGiftcardAdd      = "/card/paygiftcard/add"
	apiPayGiftcardDelete   = "/card/paygiftcard/delete"
	apiPayGiftcardGetById  = "/card/paygiftcard/getbyid"
	apiPayGiftcardBatchGet = "/card/paygiftcard/batchget"
)

/*
获取开卡插件参数

开发者可以通过该接口获取到调用开卡插件所需的参数

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Coupons-Mini_Program_Start_Up.html

POST https://api.weixin.qq.com/card/membercard/activate/geturl?access_token=
*/
func ActivateGetUrl(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiActivateGetUrl, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取用户开卡时提交的信息（跳转型开卡组件）

开发者可以通过该接口获取到用户开卡时填写的字段值

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Coupons-Mini_Program_Start_Up.html

POST https://api.weixin.qq.com/card/membercard/activatetempinfo/get?access_token=TOKEN
*/
func ActivateTempInfoGet(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiActivateTempInfoGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
激活用户领取的会员卡（跳转型开卡组件）

开发者可以通过该接口获取到用户开卡时填写的字段值

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Coupons-Mini_Program_Start_Up.html

POST https://api.weixin.qq.com/card/membercard/activate?access_token=TOKEN
*/
func Activate(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiActivate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
设置开卡字段

开发者在创建时填入wx_activate字段后，需要调用该接口设置用户激活时需要填写的选项，否则一键开卡设置不生效

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Create_a_membership_card.html

POST https://api.weixin.qq.com/card/membercard/activateuserform/set?access_token=TOKEN
*/
func ActivateUserFormSet(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiActivateUserFormSet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
拉取会员信息

支持开发者根据CardID和Code查询会员信息

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Create_a_membership_card.html

POST https://api.weixin.qq.com/card/membercard/userinfo/get?access_token=TOKEN
*/
func UserinfoGet(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUserinfoGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新会员信息

当会员持卡消费后，支持开发者调用该接口更新会员信息。会员卡交易后的每次信息变更需通过该接口通知微信，便于后续消息通知及其他扩展功能

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Create_a_membership_card.html

POST https://api.weixin.qq.com/card/membercard/updateuser?access_token=TOKEN
*/
func UpdateUser(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdateUser, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
设置支付后投放卡券

支持商户设置支付后投放卡券规则，可以区分时间段和金额区间发会员卡

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Manage_Member_Card.html

POST https://api.weixin.qq.com/card/paygiftcard/add?access_token=TOKEN
*/
func PayGiftcardAdd(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiPayGiftcardAdd, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除支付后投放卡券规则

支持商户删除之前设置的规则id

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Manage_Member_Card.html

POST https://api.weixin.qq.com/card/paygiftcard/delete?access_token=TOKEN
*/
func PayGiftcardDelete(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiPayGiftcardDelete, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询支付后投放卡券规则详情

可以查询某个支付即会员规则内容

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Manage_Member_Card.html

POST https://api.weixin.qq.com/card/paygiftcard/getbyid?access_token=TOKEN
*/
func PayGiftcardGetById(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiPayGiftcardGetById, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量查询支付后投放卡券规则

可以批量查询某个商户支付即会员规则内容

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Manage_Member_Card.html

POST https://api.weixin.qq.com/card/paygiftcard/batchget?access_token=TOKEN
*/
func PayGiftcardBatchGet(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiPayGiftcardBatchGet, bytes.NewReader(payload), "application/json;charset=utf-8")
}
