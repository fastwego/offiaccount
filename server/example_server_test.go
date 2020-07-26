package server_test

import (
	"fmt"
	"net/http"

	"github.com/fastwego/offiaccount/server"
)

func Example() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		//server.EchoStr(writer,request)

		message, err := server.GetMessage(request)
		if err != nil {
			return
		}
		switch message.(type) {
		case server.TextMessage:
			fmt.Println(message)
		default:
			fmt.Println(message)
		}

	})
	_ = http.ListenAndServe(":8080", nil)
}
