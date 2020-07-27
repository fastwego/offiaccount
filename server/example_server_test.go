package server_test

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fastwego/offiaccount/server"
)

func Example() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		//server.EchoStr(writer,request)

		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return
		}
		message, err := server.ParseMessage(body)
		if err != nil {
			return
		}
		switch message.(type) {
		case server.MessageText:
			fmt.Println(message)

			// 回复文本消息
			msg := server.ReplyMessageText{
				ReplyMessage: server.ReplyMessage{
					ToUserName:   "to_user",
					FromUserName: "from",
					CreateTime:   "1525464546",
					MsgType:      server.ReplyMsgTypeText,
				},
				Content: "Hello",
			}
			data, err := xml.Marshal(msg)
			if err != nil {
				return
			}
			writer.Write(data)
		default:
			fmt.Println(message)
		}

	})
	_ = http.ListenAndServe(":8080", nil)
}
