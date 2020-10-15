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

/*
微信公众平台开发 SDK

See: https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Overview.html
*/
package offiaccount

import (
	"log"
	"os"

	"github.com/faabiosr/cachego"
	"github.com/faabiosr/cachego/file"
)

// GetAccessTokenFunc 获取 access_token 方法接口
type GetAccessTokenFunc func(ctx *OffiAccount) (accessToken string, err error)

// NoticeAccessTokenExpireFunc 通知中控 刷新 access_token
type NoticeAccessTokenExpireFunc func(ctx *OffiAccount) (err error)

/*
OffiAccount 公众号实例
*/
type OffiAccount struct {
	Config      Config
	AccessToken AccessToken
	Client      Client
	Server      Server
	Logger      *log.Logger
}

/*
AccessToken 管理器 处理缓存 和 刷新 逻辑
*/
type AccessToken struct {
	Cache                          cachego.Cache
	GetAccessTokenHandler          GetAccessTokenFunc
	NoticeAccessTokenExpireHandler NoticeAccessTokenExpireFunc
}

/*
公众号配置
*/
type Config struct {
	Appid          string
	Secret         string
	Token          string
	EncodingAESKey string
}

/*
创建公众号实例
*/
func New(config Config) (offiAccount *OffiAccount) {
	instance := OffiAccount{
		Config: config,
		AccessToken: AccessToken{
			Cache:                          file.New(os.TempDir()),
			GetAccessTokenHandler:          GetAccessToken,
			NoticeAccessTokenExpireHandler: NoticeAccessTokenExpire,
		},
	}

	instance.Client = Client{Ctx: &instance}
	instance.Server = Server{Ctx: &instance}

	instance.Logger = log.New(os.Stdout, "[fastwego/offiaccount] ", log.LstdFlags|log.Llongfile)

	return &instance
}
