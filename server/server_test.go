package server

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"

	"github.com/fastwego/offiaccount/test"
)

func TestMain(m *testing.M) {
	test.Setup()
	os.Exit(m.Run())
}
func TestParseMessage(t *testing.T) {
	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		args    args
		wantM   interface{}
		wantErr bool
	}{
		{
			name: "case1",
			args: args{body: []byte(`
			<xml>
			  <ToUserName><![CDATA[toUser]]></ToUserName>
			  <FromUserName><![CDATA[fromUser]]></FromUserName>
			  <CreateTime>1348831860</CreateTime>
			  <MsgType><![CDATA[text]]></MsgType>
			  <Content><![CDATA[this is a test]]></Content>
			  <MsgId>1234567890123456</MsgId>
			</xml>
			`)},
			wantM: MessageText{
				Message: Message{
					ToUserName:   "toUser",
					FromUserName: "fromUser",
					CreateTime:   "1348831860",
					MsgType:      "text",
					MsgId:        "1234567890123456",
				},
				Content: "this is a test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotM, err := ParseMessage(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("ParseMessage() gotM = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func TestEchoStr(t *testing.T) {
	test.MockSvrHandler.HandleFunc("/echostr", EchoStr)

	tests := []struct {
		name     string
		args     url.Values
		wantEcho string
	}{
		{
			name: "case1",
			args: url.Values{
				"timestamp": []string{"1526545852"},
				"nonce":     []string{"nonce"},
				"echostr":   []string{"echostr"},
				"signature": []string{"ca1efcb1f2719b2d3265c68a8d12b13835ff91c7"},
			},
			wantEcho: "echostr",
		},
		{
			name: "case2",
			args: url.Values{
				"timestamp": []string{"1526545852"},
				"nonce":     []string{"nonce"},
				"echostr":   []string{"echostr"},
				"signature": []string{"123"},
			},
			wantEcho: SUCCESS,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r, _ := http.NewRequest(http.MethodGet, "/echostr?"+tt.args.Encode(), nil)
			w := httptest.NewRecorder()
			test.MockSvrHandler.ServeHTTP(w, r)
			resp := w.Result()
			if resp.StatusCode != http.StatusOK {
				t.Errorf("Response code is %v", resp.StatusCode)
			}

			echo := string(w.Body.Bytes())
			fmt.Println(echo)

			if tt.wantEcho != echo {
				t.Errorf("want %s but get %s", tt.wantEcho, echo)
			}
		})
	}
}

func TestReplyMessage(t *testing.T) {

	tests := []struct {
		name     string
		wantEcho string
	}{
		{
			name:     "case1",
			wantEcho: `<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>12345678</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[你好]]></Content></xml>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 回复文本消息
			msg := ReplyMessageText{
				ReplyMessage: ReplyMessage{
					ToUserName:   "toUser",
					FromUserName: "fromUser",
					CreateTime:   "12345678",
					MsgType:      ReplyMsgTypeText,
				},
				Content: "你好",
			}
			data, err := xml.Marshal(msg)
			fmt.Println(string(data), err)

			if tt.wantEcho != string(data) {
				t.Errorf("\nwant %s \nget %s", tt.wantEcho, string(data))
			}
		})
	}
}
