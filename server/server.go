package server

import (
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
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
	signature := fmt.Sprintf("% x", h.Sum(nil))

	echoStr := req.URL.Query().Get("echostr")
	if echoStr != "" && signature == req.URL.Query().Get("signature") {
		_, _ = io.WriteString(writer, echoStr)
	}

	_, _ = io.WriteString(writer, SUCCESS)
}

func GetMessage(req *http.Request) (m interface{}, err error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	message := Message{}
	err = xml.Unmarshal(body, &message)
	if err != nil {
		return
	}

	switch message.MsgType {
	case MsgTypeText:
		msg := TextMessage{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case MsgTypeImage:
		msg := ImageMessage{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case MsgTypeVoice:
		msg := VoiceMessage{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case MsgTypeVideo:
		msg := VideoMessage{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case MsgTypeShortVideo:
		msg := ShortVideoMessage{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case MsgTypeLocation:
		msg := LocationMessage{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case MsgTypeLink:
		msg := LinkMessage{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case MsgTypeEvent:
		msg := EventMessage{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	}
	return
}
