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

type RefreshAccessTokenFunc func(appid string, secret string) (accessToken string, expiresIn int, err error)

/*
OffiAccount 公众号实例
*/
type OffiAccount struct {
	Config      OffiAccountConfig
	AccessToken AccessToken
	Client      Client
	Server      Server
	Logger      *log.Logger
}

/*
AccessToken 管理器 处理缓存 和 刷新 逻辑
*/
type AccessToken struct {
	Cache          cachego.Cache
	RefreshHandler RefreshAccessTokenFunc
}

/*
公众号配置
*/
type OffiAccountConfig struct {
	Appid          string
	Secret         string
	Token          string
	EncodingAESKey string
}

/*
创建公众号实例
*/
func New(config OffiAccountConfig) (offiAccount *OffiAccount) {
	instance := OffiAccount{
		Config: config,
		AccessToken: AccessToken{
			Cache:          file.New(os.TempDir()),
			RefreshHandler: refreshAccessTokenFromWXServer,
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
SetRefreshAccessTokenHandler 设置 AccessToken 获取方法。默认 从微信服务器获取

如果有多实例服务，可以设置为 Redis 或 数据库 等中控服务器 获取 就可以避免 AccessToken 刷新冲突
*/
func (offiAccount *OffiAccount) SetRefreshAccessTokenHandler(f RefreshAccessTokenFunc) {
	offiAccount.AccessToken.RefreshHandler = f
}

// DeleteAccessTokenCache 删除本地 AccessToken ，从而强制刷新
func (offiAccount *OffiAccount) DeleteAccessTokenCache() {
	offiAccount.AccessToken.Cache.Delete(offiAccount.Config.Appid)
}

/*
SetLogger 日志记录 默认输出到 os.Stdout

可以新建 logger 输出到指定文件

如果不想开启日志，可以输出到 /dev/null log.SetOutput(ioutil.Discard)
*/
func (offiAccount *OffiAccount) SetLogger(logger *log.Logger) {
	offiAccount.Logger = logger
}
