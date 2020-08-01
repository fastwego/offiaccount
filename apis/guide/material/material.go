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

// Package material 素材管理
package material

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiSetGuideCardMaterial  = "/cgi-bin/guide/setguidecardmaterial"
	apiGetGuideCardMaterial  = "/cgi-bin/guide/getguidecardmaterial"
	apiDelGuideCardMaterial  = "/cgi-bin/guide/delguidecardmaterial"
	apiSetGuideImageMaterial = "/cgi-bin/guide/setguideimagematerial"
	apiGetGuideImageMaterial = "/cgi-bin/guide/getguideimagematerial"
	apiDelGuideImageMaterial = "/cgi-bin/guide/delguideimagematerial"
	apiSetGuideWordMaterial  = "/cgi-bin/guide/setguidewordmaterial"
	apiGetGuideWordMaterial  = "/cgi-bin/guide/getguidewordmaterial"
	apiDelGuideWordMaterial  = "/cgi-bin/guide/delguidewordmaterial"
)

/*
添加小程序卡片素材



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.setGuideCardMaterial.html

POST https://api.weixin.qq.com/cgi-bin/guide/setguidecardmaterial?access_token=ACCESS_TOKEN
*/
func SetGuideCardMaterial(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetGuideCardMaterial, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
查询小程序卡片素材



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.getGuideCardMaterial.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidecardmaterial?access_token=ACCESS_TOKEN
*/
func GetGuideCardMaterial(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGuideCardMaterial, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
删除小程序卡片素材



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.delGuideCardMaterial.html

POST https://api.weixin.qq.com/cgi-bin/guide/delguidecardmaterial?access_token=ACCESS_TOKEN
*/
func DelGuideCardMaterial(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelGuideCardMaterial, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
添加图片素材



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.setGuideImageMaterial.html

POST https://api.weixin.qq.com/cgi-bin/guide/setguideimagematerial?access_token=ACCESS_TOKEN
*/
func SetGuideImageMaterial(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetGuideImageMaterial, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
查询图片素材



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.getGuideImageMaterial.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguideimagematerial?access_token=ACCESS_TOKEN
*/
func GetGuideImageMaterial(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGuideImageMaterial, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
删除图片素材



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.delGuideImageMaterial.html

POST https://api.weixin.qq.com/cgi-bin/guide/delguideimagematerial?access_token=ACCESS_TOKEN
*/
func DelGuideImageMaterial(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelGuideImageMaterial, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
添加文字素材



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.setGuideWordMaterial.html

POST https://api.weixin.qq.com/cgi-bin/guide/setguidewordmaterial?access_token=ACCESS_TOKEN
*/
func SetGuideWordMaterial(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetGuideWordMaterial, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
查询文字素材



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.getGuideWordMaterial.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidewordmaterial?access_token=ACCESS_TOKEN
*/
func GetGuideWordMaterial(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGuideWordMaterial, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
删除文字素材



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.delGuideWordMaterial.html

POST https://api.weixin.qq.com/cgi-bin/guide/delguidewordmaterial?access_token=ACCESS_TOKEN
*/
func DelGuideWordMaterial(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelGuideWordMaterial, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}
