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

// Package store 微信小店/功能接口
package store

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiUploadImg = "/merchant/common/upload_img"
)

/*
上传图片



See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/common/upload_img?access_token=ACCESS_TOKEN&filename=test.png
*/
func UploadImg(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUploadImg, bytes.NewReader(payload), "application/json;charset=utf-8")
}
