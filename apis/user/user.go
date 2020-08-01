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

// Package user 用户管理
package user

import (
	"bytes"
	"net/url"

	"github.com/fastwego/offiaccount"
)

const (
	apiUpdateRemark     = "/cgi-bin/user/info/updateremark"
	apiGetUserInfo      = "/cgi-bin/user/info"
	apiBatchGetUserInfo = "/cgi-bin/user/info/batchget"
	apiGet              = "/cgi-bin/user/get"
	apiGetBlackList     = "/cgi-bin/tags/members/getblacklist"
	apiBatchBlackList   = "/cgi-bin/tags/members/batchblacklist"
	apiBatchUnBlackList = "/cgi-bin/tags/members/batchunblacklist"
)

/*
设置用户备注名

开发者可以通过该接口对指定用户设置备注名，该接口暂时开放给微信认证的服务号

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/Configuring_user_notes.html

POST https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=ACCESS_TOKEN
*/
func UpdateRemark(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdateRemark, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取用户基本信息

在关注者与公众号产生消息交互后，公众号可获得关注者的OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的。对于不同公众号，同一用户的openid不同）。公众号可通过本接口来根据OpenID获取用户基本信息，包括昵称、头像、性别、所在城市、语言和关注时间

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId

GET https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
*/
func GetUserInfo(ctx *offiaccount.OffiAccount, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetUserInfo + "?" + params.Encode())
}

/*
批量获取用户基本信息

开发者可通过该接口来批量获取用户基本信息。最多支持一次拉取100条

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId

POST https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=ACCESS_TOKEN
*/
func BatchGetUserInfo(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBatchGetUserInfo, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取用户列表

公众号可通过本接口来获取帐号的关注者列表，关注者列表由一串OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成。一次拉取调用最多拉取10000个关注者的OpenID，可以通过多次拉取的方式来满足需求

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/Getting_a_User_List.html

GET https://api.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&next_openid=NEXT_OPENID
*/
func Get(ctx *offiaccount.OffiAccount, params url.Values) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGet + "?" + params.Encode())
}

/*
获取公众号的黑名单列表

公众号可通过该接口来获取帐号的黑名单列表，黑名单列表由一串 OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html

POST https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token=ACCESS_TOKEN
*/
func GetBlackList(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetBlackList, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
拉黑用户

公众号可通过该接口来拉黑一批用户，黑名单列表由一串 OpenID （加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html

POST https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist?access_token=ACCESS_TOKEN
*/
func BatchBlackList(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBatchBlackList, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
取消拉黑用户

公众号可通过该接口来取消拉黑一批用户，黑名单列表由一串OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html

POST https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist?access_token=ACCESS_TOKEN
*/
func BatchUnBlackList(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBatchUnBlackList, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}
