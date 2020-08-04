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

// Package job 服务号对话能力（原微信导购助手）/群发任务
package job

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiAddGuideMassendJob     = "/cgi-bin/guide/addguidemassendjob"
	apiGetGuideMassendJobList = "/cgi-bin/guide/getguidemassendjoblist"
	apiGetGuideMassendJob     = "/cgi-bin/guide/getguidemassendjob"
	apiUpdateGuideMassendJob  = "/cgi-bin/guide/updateguidemassendjob"
	apiCancelGuideMassendJob  = "/cgi-bin/guide/cancelguidemassendjob"
)

/*
添加群发任务



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/task-account/shopping-guide.addGuideMassendJob.html

POST https://api.weixin.qq.com/cgi-bin/guide/addguidemassendjob?access_token=ACCESS_TOKEN
*/
func AddGuideMassendJob(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddGuideMassendJob, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取群发任务列表



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/task-account/shopping-guide.getGuideMassendJobList.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidemassendjoblist?access_token=ACCESS_TOKEN
*/
func GetGuideMassendJobList(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGuideMassendJobList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取指定群发任务信息



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/task-account/shopping-guide.getGuideMassendJob.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidemassendjob?access_token=ACCESS_TOKEN
*/
func GetGuideMassendJob(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGuideMassendJob, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
修改群发任务



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/task-account/shopping-guide.updateGuideMassendJob.html

POST https://api.weixin.qq.com/cgi-bin/guide/updateguidemassendjob?access_token=ACCESS_TOKEN
*/
func UpdateGuideMassendJob(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdateGuideMassendJob, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
取消群发任务



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/task-account/shopping-guide.cancelGuideMassendJob.html

POST https://api.weixin.qq.com/cgi-bin/guide/cancelguidemassendjob?access_token=ACCESS_TOKEN
*/
func CancelGuideMassendJob(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCancelGuideMassendJob, bytes.NewReader(payload), "application/json;charset=utf-8")
}
