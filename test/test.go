package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"sync"

	filecache "github.com/faabiosr/cachego/file"
	"github.com/fastwego/offiaccount"
)

var MockSvr *httptest.Server
var MockSvrHandler *http.ServeMux
var onceSetup sync.Once

func Setup() {
	onceSetup.Do(func() {
		offiaccount.Appid = "mock"
		offiaccount.Secret = "mock"
		offiaccount.SetAccessTokenCache(filecache.New(os.TempDir()))

		// Mock Server
		MockSvrHandler = http.NewServeMux()
		MockSvr = httptest.NewServer(MockSvrHandler)
		offiaccount.WXServerUrl = MockSvr.URL // 拦截发往微信服务器的请求

		// Mock access token
		MockSvrHandler.HandleFunc("/cgi-bin/token", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"access_token":"ACCESS_TOKEN","expires_in":7200}`))
		})
	})
}
