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

// Package comment 图文消息留言管理
package comment

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiOpen        = "/cgi-bin/comment/open"
	apiClose       = "/cgi-bin/comment/close"
	apiList        = "/cgi-bin/comment/list"
	apiMarkElect   = "/cgi-bin/comment/markelect"
	apiUnMarkElect = "/cgi-bin/comment/unmarkelect"
	apiDelete      = "/cgi-bin/comment/delete"
	apiReplyAdd    = "/cgi-bin/comment/reply/add"
	apiReplyDelete = "/cgi-bin/comment/reply/delete"
)

/*
打开已群发文章评论



See: https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html

POST https://api.weixin.qq.com/cgi-bin/comment/open?access_token=ACCESS_TOKEN
*/
func Open(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiOpen, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
关闭已群发文章评论



See: https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html

POST https://api.weixin.qq.com/cgi-bin/comment/close?access_token=ACCESS_TOKEN
*/
func Close(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiClose, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查看指定文章的评论数据



See: https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html

POST https://api.weixin.qq.com/cgi-bin/comment/list?access_token=ACCESS_TOKEN
*/
func List(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
 将评论标记精选



See: https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html

POST https://api.weixin.qq.com/cgi-bin/comment/markelect?access_token=ACCESS_TOKEN
*/
func MarkElect(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMarkElect, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
 将评论取消精选



See: https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html

POST https://api.weixin.qq.com/cgi-bin/comment/unmarkelect?access_token=ACCESS_TOKEN
*/
func UnMarkElect(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUnMarkElect, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
 删除评论



See: https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html

POST https://api.weixin.qq.com/cgi-bin/comment/delete?access_token=ACCESS_TOKEN
*/
func Delete(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelete, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
 回复评论



See: https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html

POST https://api.weixin.qq.com/cgi-bin/comment/reply/add?access_token=ACCESS_TOKEN
*/
func ReplyAdd(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiReplyAdd, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
 删除回复



See: https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html

POST https://api.weixin.qq.com/cgi-bin/comment/reply/delete?access_token=ACCESS_TOKEN
*/
func ReplyDelete(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiReplyDelete, bytes.NewReader(payload), "application/json;charset=utf-8")
}
