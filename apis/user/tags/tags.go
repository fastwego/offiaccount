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

// Package tags 用户标签管理
package tags

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiCreate         = "/cgi-bin/tags/create"
	apiGet            = "/cgi-bin/tags/get"
	apiUpdate         = "/cgi-bin/tags/update"
	apiDelete         = "/cgi-bin/tags/delete"
	apiGetUsersByTag  = "/cgi-bin/user/tag/get"
	apiBatchTagging   = "/cgi-bin/tags/members/batchtagging"
	apiBatchUnTagging = "/cgi-bin/tags/members/batchuntagging"
	apiGetTagIdList   = "/cgi-bin/tags/getidlist"
)

/*
创建标签

一个公众号，最多可以创建100个标签

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html

POST https://api.weixin.qq.com/cgi-bin/tags/create?access_token=ACCESS_TOKEN
*/
func Create(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取公众号已创建的标签



See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html

GET https://api.weixin.qq.com/cgi-bin/tags/get?access_token=ACCESS_TOKEN
*/
func Get(ctx *offiaccount.OffiAccount) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGet)
}

/*
编辑标签



See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html

POST https://api.weixin.qq.com/cgi-bin/tags/update?access_token=ACCESS_TOKEN
*/
func Update(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdate, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除标签

请注意，当某个标签下的粉丝超过10w时，后台不可直接删除标签。此时，开发者可以对该标签下的openid列表，先进行取消标签的操作，直到粉丝数不超过10w后，才可直接删除该标签

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html

POST https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=ACCESS_TOKEN
*/
func Delete(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelete, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取标签下粉丝列表



See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html

POST https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=ACCESS_TOKEN
*/
func GetUsersByTag(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetUsersByTag, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量为用户打标签

标签功能目前支持公众号为用户打上最多20个标签

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html

POST https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token=ACCESS_TOKEN
*/
func BatchTagging(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBatchTagging, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
批量为用户取消标签

标签功能目前支持公众号为用户打上最多20个标签

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html

POST https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token=ACCESS_TOKEN
*/
func BatchUnTagging(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBatchUnTagging, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取用户身上的标签列表

标签功能目前支持公众号为用户打上最多20个标签

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html

POST https://api.weixin.qq.com/cgi-bin/tags/getidlist?access_token=ACCESS_TOKEN
*/
func GetTagIdList(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetTagIdList, bytes.NewReader(payload), "application/json;charset=utf-8")
}
