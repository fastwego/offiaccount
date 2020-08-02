package offiaccount

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_getAccessToken(t *testing.T) {

	var MockOffiAccount = New(OffiAccountConfig{
		Appid:  "TestClient_getAccessToken",
		Secret: "SECRET",
	})

	// Mock Server
	var MockSvrHandler = http.NewServeMux()
	var MockSvr = httptest.NewServer(MockSvrHandler)
	WXServerUrl = MockSvr.URL // 拦截发往微信服务器的请求

	mockResp := map[string][]byte{
		"case1": []byte(`{"access_token":"ACCESS_TOKEN","expires_in":3}`),
		"case2": []byte(`{"errcode":40013,"errmsg":"invalid appid"}`),
	}
	var resp []byte

	// Mock access token
	MockSvrHandler.HandleFunc("/cgi-bin/token", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(resp)
	})
	client := &Client{
		Ctx: MockOffiAccount,
	}

	tests := []struct {
		name            string
		wantAccessToken string
		wantErr         bool
	}{
		{name: "case1", wantAccessToken: "ACCESS_TOKEN", wantErr: false},
		{name: "case2", wantAccessToken: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client.Ctx.DeleteAccessTokenCache()
			resp = mockResp[tt.name]
			gotAccessToken, err := client.getAccessToken()
			fmt.Println(gotAccessToken, err)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAccessToken != tt.wantAccessToken {
				t.Errorf("getAccessToken() gotAccessToken = %v, want %v", gotAccessToken, tt.wantAccessToken)
			}
		})
	}
}
