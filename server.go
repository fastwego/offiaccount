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

package offiaccount

import (
	"crypto/sha1"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	eventtype "github.com/fastwego/offiaccount/type/type_event"
	messagetype "github.com/fastwego/offiaccount/type/type_message"

	"github.com/fastwego/offiaccount/util"
)

const SUCCESS = "success"

type Server struct {
	Ctx *OffiAccount
}

// EchoStr 服务器接口校验
func (s *Server) EchoStr(writer http.ResponseWriter, request *http.Request) {
	strs := []string{
		request.URL.Query().Get("timestamp"),
		request.URL.Query().Get("nonce"),
		s.Ctx.Config.Token,
	}
	sort.Strings(strs)

	h := sha1.New()
	_, _ = io.WriteString(h, strings.Join(strs, ""))
	signature := fmt.Sprintf("%x", h.Sum(nil))

	echoStr := request.URL.Query().Get("echostr")
	if echoStr != "" && signature == request.URL.Query().Get("signature") {
		io.WriteString(writer, echoStr)
	}
}

// ParseMessage 解析微信推送过来的消息
func (s *Server) ParseMessage(body []byte) (m interface{}, err error) {

	// 是否加密消息
	encryptMsg := messagetype.EncryptMessage{}
	err = xml.Unmarshal(body, &encryptMsg)
	if err != nil {
		return
	}

	if encryptMsg.Encrypt != "" {
		var xmlMsg []byte
		var appId []byte
		_, xmlMsg, appId, err = util.AESDecryptMsg(encryptMsg.Encrypt, s.Ctx.Config.EncodingAESKey)
		if err != nil {
			return
		}

		if string(appId) != s.Ctx.Config.Appid {
			err = errors.New("appid not match")
			return
		}

		body = xmlMsg
	}

	message := messagetype.Message{}
	err = xml.Unmarshal(body, &message)
	if err != nil {
		return
	}

	switch message.MsgType {
	case messagetype.MsgTypeText:
		msg := messagetype.MessageText{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case messagetype.MsgTypeImage:
		msg := messagetype.MessageImage{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case messagetype.MsgTypeVoice:
		msg := messagetype.MessageVoice{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case messagetype.MsgTypeVideo:
		msg := messagetype.MessageVideo{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case messagetype.MsgTypeShortVideo:
		msg := messagetype.MessageShortVideo{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case messagetype.MsgTypeLocation:
		msg := messagetype.MessageLocation{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case messagetype.MsgTypeLink:
		msg := messagetype.MessageLink{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case messagetype.MsgTypeEvent:
		msg := messagetype.MessageEvent{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	}
	return
}

// ParseEvent 解析微信推送过来的事件
func (s *Server) ParseEvent(body []byte) (m interface{}, err error) {
	event := eventtype.Event{}
	err = xml.Unmarshal(body, &event)
	if err != nil {
		return
	}

	switch event.Event {
	case eventtype.EventTypeSubscribe:
		msg := eventtype.EventSubscribe{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeUnsubscribe:
		msg := eventtype.EventUnsubscribe{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeScan:
		msg := eventtype.EventScan{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeLocation:
		msg := eventtype.EventLocation{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeMenuClick:
		msg := eventtype.EventMenuClick{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeMenuView:
		msg := eventtype.EventMenuView{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeMenuScanCodePush:
		msg := eventtype.EventMenuScanCodePush{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeMenuScanCodeWaitMsg:
		msg := eventtype.EventMenuScanCodeWaitMsg{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeMenuPicSysPhoto:
		msg := eventtype.EventMenuPicSysPhoto{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeMenuPicSysPhotoOrAlbum:
		msg := eventtype.EventMenuPicSysPhotoOrAlbum{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeMenuPicWeixin:
		msg := eventtype.EventMenuPicWeixin{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeMenuLocationSelect:
		msg := eventtype.EventMenuLocationSelect{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeMenuViewMiniprogram:
		msg := eventtype.EventMenuViewMiniprogram{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	}
	return
}

// Response 响应微信消息 (自动判断是否要加密)
func (s *Server) Response(writer http.ResponseWriter, request *http.Request, reply []byte) (err error) {

	// 如果 开启加密，微信服务器 发过来的请求 带有 如下参数
	//signature=c44d29564aa1d57bd0e274c37baa92bd5b3da5bd&timestamp=1596184957
	//&nonce=1250398014
	//&openid=oEnxesxpxWw-PKkz-vW5IMdfcQaE
	//&encrypt_type=aes
	//&msg_signature=cc24cc38467417603fc3689170e8b0fd3c9bf4a2

	// 所以 响应也要加密
	if request.URL.Query().Get("encrypt_type") == "aes" {
		// 加密
		message := s.encryptReplyMessage(reply)
		reply, err = xml.Marshal(message)
		if err != nil {
			return
		}
	}

	_, err = writer.Write(reply)

	log.Println("Response : ", string(reply))

	return
}

// encryptReplyMessage 加密回复消息
func (s *Server) encryptReplyMessage(rawXmlMsg []byte) (replyEncryptMessage messagetype.ReplyEncryptMessage) {
	cipherText := util.AESEncryptMsg([]byte(util.GetRandString(16)), rawXmlMsg, s.Ctx.Config.Appid, s.Ctx.Config.EncodingAESKey)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := util.GetRandString(6)

	strs := []string{
		timestamp,
		nonce,
		s.Ctx.Config.Token,
		cipherText,
	}
	sort.Strings(strs)
	h := sha1.New()
	_, _ = io.WriteString(h, strings.Join(strs, ""))
	signature := fmt.Sprintf("%x", h.Sum(nil))

	return messagetype.ReplyEncryptMessage{
		Encrypt:      cipherText,
		MsgSignature: signature,
		TimeStamp:    timestamp,
		Nonce:        nonce,
	}
}
