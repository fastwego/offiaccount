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
	Cache                 cachego.Cache
	GetAccessTokenHandler GetAccessTokenFunc
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
			Cache:                 file.New(os.TempDir()),
			GetAccessTokenHandler: GetAccessToken,
		},
	}

	instance.Client = Client{Ctx: &instance}
	instance.Server = Server{Ctx: &instance}

	instance.Logger = log.New(os.Stdout, "[offiaccount] ", log.LstdFlags)

	return &instance
}

/*
SetAccessTokenCacheDriver 设置 AccessToken 缓存器 默认为文件缓存：目录 os.TempDir()

驱动接口类型 为 cachego.Cache
*/
func (offiAccount *OffiAccount) SetAccessTokenCacheDriver(driver cachego.Cache) {
	offiAccount.AccessToken.Cache = driver
}

/*
SetGetAccessTokenHandler 设置 AccessToken 获取方法。默认 从本地缓存获取（过期从微信接口刷新）

如果有多实例服务，可以设置为 Redis 或 RPC 等中控服务器 获取 就可以避免 AccessToken 刷新冲突
*/
func (offiAccount *OffiAccount) SetGetAccessTokenHandler(f GetAccessTokenFunc) {
	offiAccount.AccessToken.GetAccessTokenHandler = f
}

/*
SetLogger 日志记录 默认输出到 os.Stdout

可以新建 logger 输出到指定文件

如果不想开启日志，可以 SetLogger(nil)
*/
func (offiAccount *OffiAccount) SetLogger(logger *log.Logger) {
	offiAccount.Logger = logger
}
