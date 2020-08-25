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

package oauth

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/fastwego/offiaccount"
)

var MockSvr *httptest.Server
var MockSvrHandler *http.ServeMux

func TestMain(m *testing.M) {
	// Mock Server
	MockSvrHandler = http.NewServeMux()
	MockSvr = httptest.NewServer(MockSvrHandler)
	offiaccount.WXServerUrl = MockSvr.URL // 拦截发往微信服务器的请求

	os.Exit(m.Run())
}

func TestAuth(t *testing.T) {
	// Mock
	MockSvrHandler.HandleFunc(apiAuth, func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{ "errcode":0,"errmsg":"ok"}`))
	})

	type args struct {
		access_token string
		openid       string
	}
	tests := []struct {
		name        string
		args        args
		wantIsValid bool
		wantErr     bool
	}{
		{name: "case1", args: args{access_token: "", openid: ""}, wantIsValid: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsValid, err := Auth(tt.args.access_token, tt.args.openid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Auth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIsValid != tt.wantIsValid {
				t.Errorf("Auth() gotIsValid = %v, want %v", gotIsValid, tt.wantIsValid)
			}
		})
	}
}

func TestGetAccessToken(t *testing.T) {
	// Mock
	MockSvrHandler.HandleFunc(apiAccessToken, func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{
		  "access_token":"ACCESS_TOKEN",
		  "expires_in":7200,
		  "refresh_token":"REFRESH_TOKEN",
		  "openid":"OPENID",
		  "scope":"SCOPE" 
		}`))
	})

	type args struct {
		appid  string
		secret string
		code   string
	}
	tests := []struct {
		name                 string
		args                 args
		wantOauthAccessToken OauthAccessToken
		wantErr              bool
	}{
		{name: "case1", args: args{appid: "", secret: "", code: ""}, wantOauthAccessToken: OauthAccessToken{
			AccessToken:  "ACCESS_TOKEN",
			ExpiresIn:    7200,
			RefreshToken: "REFRESH_TOKEN",
			Openid:       "OPENID",
			Scope:        "SCOPE",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOauthAccessToken, err := GetAccessToken(tt.args.appid, tt.args.secret, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOauthAccessToken, tt.wantOauthAccessToken) {
				t.Errorf("GetAccessToken() gotOauthAccessToken = %v, want %v", gotOauthAccessToken, tt.wantOauthAccessToken)
			}
		})
	}
}

func TestGetAuthorizeUrl(t *testing.T) {
	type args struct {
		appid       string
		redirectUri string
		scope       string
		state       string
	}
	tests := []struct {
		name             string
		args             args
		wantAuthorizeUrl string
	}{
		{name: "case1", args: args{appid: "appid", redirectUri: "https://fastwego.dev/api/weixin/oauth", scope: ScopeSnsapiUserinfo, state: "STATE"}, wantAuthorizeUrl: "https://open.weixin.qq.com/connect/oauth2/authorize?appid=appid&redirect_uri=https%3A%2F%2Ffastwego.dev%2Fapi%2Fweixin%2Foauth&response_type=code&scope=snsapi_userinfo&state=STATE"},
		{name: "case2", args: args{appid: "appid", redirectUri: "https://fastwego.dev/api/weixin/oauth", scope: ScopeSnsapiBase, state: "STATE"}, wantAuthorizeUrl: "https://open.weixin.qq.com/connect/oauth2/authorize?appid=appid&redirect_uri=https%3A%2F%2Ffastwego.dev%2Fapi%2Fweixin%2Foauth&response_type=code&scope=snsapi_base&state=STATE"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAuthorizeUrl := GetAuthorizeUrl(tt.args.appid, tt.args.redirectUri, tt.args.scope, tt.args.state)
			fmt.Println(gotAuthorizeUrl)
			if gotAuthorizeUrl != tt.wantAuthorizeUrl {
				t.Errorf("GetAuthorizeUrl() = %v \n want %v", gotAuthorizeUrl, tt.wantAuthorizeUrl)
			}
		})
	}
}

func TestGetUserInfo(t *testing.T) {
	// Mock
	MockSvrHandler.HandleFunc(apiUserInfo, func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{   
          "openid":"OPENID",
		  "nickname": "NICKNAME",
		  "sex":1,
		  "province":"PROVINCE",
		  "city":"CITY",
		  "country":"COUNTRY",
		  "headimgurl":"http://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
		  "privilege":[ "PRIVILEGE1","PRIVILEGE2"     ],
		  "unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL"
		}`))
	})

	type args struct {
		access_token string
		openid       string
		lang         string
	}
	tests := []struct {
		name              string
		args              args
		wantOauthUserInfo OauthUserInfo
		wantErr           bool
	}{
		{name: "case1", args: args{access_token: "", openid: "", lang: LANG_zh_CN}, wantOauthUserInfo: OauthUserInfo{
			Openid:     "OPENID",
			Nickname:   "NICKNAME",
			Sex:        1,
			Province:   "PROVINCE",
			City:       "CITY",
			Country:    "COUNTRY",
			Headimgurl: "http://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
			Privilege:  []string{"PRIVILEGE1", "PRIVILEGE2"},
			Unionid:    "o6_bmasdasdsad6_2sgVt7hMZOPfL",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOauthUserInfo, err := GetUserInfo(tt.args.access_token, tt.args.openid, tt.args.lang)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOauthUserInfo, tt.wantOauthUserInfo) {
				t.Errorf("GetUserInfo() gotOauthUserInfo = \n%v,\n want \n%v", gotOauthUserInfo, tt.wantOauthUserInfo)
			}
		})
	}
}

func TestRefreshToken(t *testing.T) {
	// Mock
	MockSvrHandler.HandleFunc(apiRefreshToken, func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{
		  "access_token":"ACCESS_TOKEN",
		  "expires_in":7200,
		  "refresh_token":"REFRESH_TOKEN",
		  "openid":"OPENID",
		  "scope":"SCOPE" 
		}`))
	})
	type args struct {
		appid         string
		refresh_token string
	}
	tests := []struct {
		name                 string
		args                 args
		wantOauthAccessToken OauthAccessToken
		wantErr              bool
	}{
		{name: "case1", args: args{appid: "", refresh_token: ""}, wantOauthAccessToken: OauthAccessToken{
			AccessToken:  "ACCESS_TOKEN",
			ExpiresIn:    7200,
			RefreshToken: "REFRESH_TOKEN",
			Openid:       "OPENID",
			Scope:        "SCOPE",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOauthAccessToken, err := RefreshToken(tt.args.appid, tt.args.refresh_token)
			if (err != nil) != tt.wantErr {
				t.Errorf("RefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOauthAccessToken, tt.wantOauthAccessToken) {
				t.Errorf("RefreshToken() gotOauthAccessToken = %v, want %v", gotOauthAccessToken, tt.wantOauthAccessToken)
			}
		})
	}
}
