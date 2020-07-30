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
	"os"

	"github.com/faabiosr/cachego"
	"github.com/faabiosr/cachego/file"
)

type RefreshAccessTokenFunc func(appid string, secret string) (accessToken string, expiresIn int, err error)

type OffiAccount struct {
	Config      OffiAccountConfig
	AccessToken AccessTokne
	Client      Client
	Server      Server
}

type AccessTokne struct {
	Cache          cachego.Cache
	RefreshHandler RefreshAccessTokenFunc
}

type OffiAccountConfig struct {
	Appid          string
	Secret         string
	Token          string
	EncodingAESKey string
}

func New(config OffiAccountConfig) (offiAccount *OffiAccount) {
	instance := OffiAccount{
		Config: config,
		AccessToken: AccessTokne{
			Cache:          file.New(os.TempDir()),
			RefreshHandler: RefreshAccessTokenFromWXServer,
		},
	}

	instance.Client = Client{Ctx: &instance}
	instance.Server = Server{Ctx: &instance}

	return &instance
}

func (offiAccount *OffiAccount) SetAccessTokenCache(cacheDriver cachego.Cache) {
	offiAccount.AccessToken.Cache = cacheDriver
}

func (offiAccount *OffiAccount) SetRefreshAccessTokenHandler(f RefreshAccessTokenFunc) {
	offiAccount.AccessToken.RefreshHandler = f
}

func (offiAccount *OffiAccount) DeleteAccessTokenCache() {
	offiAccount.AccessToken.Cache.Delete(offiAccount.Config.Appid)
}
