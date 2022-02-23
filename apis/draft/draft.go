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

// Package draft 草稿箱
package draft

import (
	"bytes"
	"github.com/fastwego/offiaccount"
)

const (
	apiAddDraft          = "/cgi-bin/draft/add"
	apiGetDraft          = "/cgi-bin/draft/get"
	apiDeleteDraft       = "/cgi-bin/draft/delete"
	apiUpdateDraft       = "/cgi-bin/draft/update"
	apiCountDrafts       = "/cgi-bin/draft/count"
	apiGetDraftList      = "/cgi-bin/draft/batchget"
	apiTemporaryMPSwitch = "/cgi-bin/draft/switch"
)

/*
新建草稿

开发者可新增常用的素材到草稿箱中进行使用。上传到草稿箱中的素材被群发或发布后，该素材将从草稿箱中移除。新增草稿可在公众平台官网-草稿箱中查看和管理。

See: https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Add_draft.html

POST https://api.weixin.qq.com/cgi-bin/draft/add?access_token=ACCESS_TOKEN
*/
func AddDraft(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddDraft, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取草稿

新增草稿后，开发者可以根据草稿指定的字段来下载草稿。

See: https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Get_draft.html

POST https://api.weixin.qq.com/cgi-bin/draft/get?access_token=ACCESS_TOKEN
*/
func GetDraft(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetDraft, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
删除草稿

新增草稿后，开发者可以根据本接口来删除不再需要的草稿，节省空间。此操作无法撤销，请谨慎操作。

See: https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Delete_draft.html

POST https://api.weixin.qq.com/cgi-bin/draft/delete?access_token=ACCESS_TOKEN
*/
func DeleteDraft(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDeleteDraft, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
修改草稿

开发者可通过本接口对草稿进行修改。

See: https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Update_draft.html

POST https://api.weixin.qq.com/cgi-bin/draft/update?access_token=ACCESS_TOKEN
*/
func UpdateDraft(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdateDraft, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取草稿总次数

开发者可以根据本接口来获取草稿的总数。此接口只统计数量，不返回草稿的具体内容。

See: https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Count_drafts.html

GET https://api.weixin.qq.com/cgi-bin/draft/count?access_token=ACCESS_TOKEN
*/
func CountDrafts(ctx *offiaccount.OffiAccount) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiCountDrafts)
}

/*
获取草稿列表

新增草稿之后，开发者可以获取草稿的列表。

See: https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Get_draft_list.html

POST https://api.weixin.qq.com/cgi-bin/draft/batchget?access_token=ACCESS_TOKEN
*/
func GetDraftList(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetDraftList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
MP端开关（仅内测期间使用）

由于草稿箱和发布功能仍处于内测阶段，若公众号没有被灰度覆盖，可能无法体验草稿箱和发布功能。为了解决这个问题，我们在上述API接口的基础上，设了这样一个开关：当一个公众号选择开启后，该帐号在微信公众平台后台（mp.weixin.qq.com)上的图文素材库将升级为草稿箱，并可以在微信公众平台后台使用发布功能。

See: https://developers.weixin.qq.com/doc/offiaccount/Draft_Box/Temporary_MP_Switch.html

POST https://api.weixin.qq.com/cgi-bin/draft/switch?access_token=ACCESS_TOKEN
*/
func TemporaryMPSwitch(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiTemporaryMPSwitch, bytes.NewReader(payload), "application/json;charset=utf-8")
}
