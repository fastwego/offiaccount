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

// Package mass 群发消息
package mass

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiMediaUploadNews  = "/cgi-bin/media/uploadnews"
	apiSendAll          = "/cgi-bin/message/mass/sendall"
	apiMediaUploadVideo = "/cgi-bin/media/uploadvideo"
	apiSend             = "/cgi-bin/message/mass/send"
	apiDelete           = "/cgi-bin/message/mass/delete"
	apiPreview          = "/cgi-bin/message/mass/preview"
	apiGet              = "/cgi-bin/message/mass/get"
	apiSpeedGet         = "/cgi-bin/message/mass/speed/get"
	apiSpeedSet         = "/cgi-bin/message/mass/speed/set"
)

/*
上传图文消息素材



See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html

POST https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token=ACCESS_TOKEN
*/
func MediaUploadNews(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMediaUploadNews, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
根据标签进行群发



See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html

POST https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=ACCESS_TOKEN
*/
func SendAll(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSendAll, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
上传视频素材

群发 视频 的 media_id 需通过此接口特别地得到

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html

POST https://api.weixin.qq.com/cgi-bin/media/uploadvideo?access_token=ACCESS_TOKEN
*/
func MediaUploadVideo(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMediaUploadVideo, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
根据OpenID列表群发



See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html

POST https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=ACCESS_TOKEN
*/
func Send(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSend, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
删除群发



See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html

POST https://api.weixin.qq.com/cgi-bin/message/mass/delete?access_token=ACCESS_TOKEN
*/
func Delete(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelete, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
预览

开发者可通过该接口发送消息给指定用户，在手机端查看消息的样式和排版。为了满足第三方平台开发者的需求，在保留对openID预览能力的同时，增加了对指定微信号发送预览的能力，但该能力每日调用次数有限制（100次），请勿滥用

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html

POST https://api.weixin.qq.com/cgi-bin/message/mass/preview?access_token=ACCESS_TOKEN
*/
func Preview(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiPreview, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
查询群发消息发送状态



See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html

POST https://api.weixin.qq.com/cgi-bin/message/mass/get?access_token=ACCESS_TOKEN
*/
func Get(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGet, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取群发速度



See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html

POST https://api.weixin.qq.com/cgi-bin/message/mass/speed/get?access_token=ACCESS_TOKEN
*/
func SpeedGet(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSpeedGet, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
设置群发速度



See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html

POST https://api.weixin.qq.com/cgi-bin/message/mass/speed/set?access_token=ACCESS_TOKEN
*/
func SpeedSet(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSpeedSet, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}
