package server

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
)

const SUCCESS = "SUCCESS"

var Token string
var EncodingAESKey string

func EchoStr(writer http.ResponseWriter, req *http.Request) {
	strs := []string{
		req.URL.Query().Get("timestamp"),
		req.URL.Query().Get("nonce"),
		Token,
	}
	sort.Strings(strs)

	h := sha1.New()
	_, _ = io.WriteString(h, strings.Join(strs, ""))
	signature := fmt.Sprintf("%x", h.Sum(nil))

	echoStr := req.URL.Query().Get("echostr")
	if echoStr != "" && signature == req.URL.Query().Get("signature") {
		_, _ = io.WriteString(writer, echoStr)
		return
	}

	_, _ = io.WriteString(writer, SUCCESS)
}

func ParseMessage(body []byte) (m interface{}, err error) {
	message := Message{}
	err = xml.Unmarshal(body, &message)
	if err != nil {
		return
	}

	switch message.MsgType {
	case MsgTypeText:
		msg := MessageText{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case MsgTypeImage:
		msg := MessageImage{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case MsgTypeVoice:
		msg := MessageVoice{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case MsgTypeVideo:
		msg := MessageVideo{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case MsgTypeShortVideo:
		msg := MessageShortVideo{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case MsgTypeLocation:
		msg := MessageLocation{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case MsgTypeLink:
		msg := MessageLink{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case MsgTypeEvent:
		msg := MessageEvent{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	}
	return
}
