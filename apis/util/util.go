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

// Package util 获取微信服务器 ip
package util

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiGetCallbackIp  = "/cgi-bin/getcallbackip"
	apiGetApiDomainIp = "/cgi-bin/get_api_domain_ip"
	apiCallbackCheck  = "/cgi-bin/callback/check"
	apiClearQuota     = "/cgi-bin/clear_quota"
)

/*
获取微信服务器IP地址

如果公众号基于安全等考虑，需要获知微信服务器的IP地址列表，以便进行相关限制，可以通过该接口获得微信服务器IP地址列表或者IP网段信息

See: https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html

GET https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=ACCESS_TOKEN
*/
func GetCallbackIp(ctx *offiaccount.OffiAccount) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetCallbackIp)
}

/*
获取微信API接口 IP地址

即api.weixin.qq.com的解析地址，由开发者调用微信侧的接入IP

See: https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_the_WeChat_server_IP_address.html

GET https://api.weixin.qq.com/cgi-bin/get_api_domain_ip?access_token=ACCESS_TOKEN
*/
func GetApiDomainIp(ctx *offiaccount.OffiAccount) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetApiDomainIp)
}

/*
网络检测

为了帮助开发者排查回调连接失败的问题，提供这个网络检测的API。它可以对开发者URL做域名解析，然后对所有IP进行一次ping操作，得到丢包率和耗时

See: https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Network_Detection.html

POST https://api.weixin.qq.com/cgi-bin/callback/check?access_token=ACCESS_TOKEN
*/
func CallbackCheck(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCallbackCheck, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
对公众号的所有api调用次数进行清零



See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/API_Call_Limits.html

POST https://api.weixin.qq.com/cgi-bin/clear_quota?access_token=ACCESS_TOKEN
*/
func ClearQuota(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiClearQuota, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}
