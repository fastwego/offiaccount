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

// Package cps 返佣商品
package cps

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiProductAdd           = "/scan/product/v2/add"
	apiProductStatus        = "/scan/product/v2/status"
	apiProductGetInfo       = "/scan/product/v2/getinfo"
	apiProductGetInfoByPage = "/scan/product/v2/getinfobypage"
)

/*
导入/更新商品

每次调用支持批量导入不超过1000条的商品信息。每分钟单个商户全局调用次数不得超过200次。每天调用次数不得超过100万次。每次请求包大小不超过2M

See: https://mp.weixin.qq.com/cgi-bin/announce?action=getannouncement&key=11533749572M9ODP&version=1&lang=zh_CN&platform=2

POST https://api.weixin.qq.com/scan/product/v2/add?access_token=ACCESS_TOKEN
*/
func ProductAdd(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiProductAdd, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
导入或更新结果查询

用于查询导入或更新商品的结果，当导入或更新商品失败时，若为系统错误可进行重试；若为其他错误，请排查解决后进行重试

See: https://mp.weixin.qq.com/cgi-bin/announce?action=getannouncement&key=11533749572M9ODP&version=1&lang=zh_CN&platform=2

POST https://api.weixin.qq.com/scan/product/v2/status?access_token=ACCESS_TOKEN
*/
func ProductStatus(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiProductStatus, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
单个商品信息查询

使用该接口，商户可获取已导入的商品信息，供验证信息及抽查导入情况使用

See: https://mp.weixin.qq.com/cgi-bin/announce?action=getannouncement&key=11533749572M9ODP&version=1&lang=zh_CN&platform=2

POST https://api.weixin.qq.com/scan/product/v2/getinfo?access_token=ACCESS_TOKEN
*/
func ProductGetInfo(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiProductGetInfo, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
全量商品信息查询

使用该接口，商户可获取已导入的全量商品信息，供全量验证信息使用

See: https://mp.weixin.qq.com/cgi-bin/announce?action=getannouncement&key=11533749572M9ODP&version=1&lang=zh_CN&platform=2

POST https://api.weixin.qq.com/scan/product/v2/getinfobypage?access_token=ACCESS_TOKEN
*/
func ProductGetInfoByPage(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiProductGetInfoByPage, bytes.NewReader(payload), "application/json;charset=utf-8")
}
