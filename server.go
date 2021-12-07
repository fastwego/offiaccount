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
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	eventtype "github.com/fastwego/offiaccount/type/type_event"
	messagetype "github.com/fastwego/offiaccount/type/type_message"

	"github.com/fastwego/offiaccount/util"
)

/*
响应微信请求 或 推送消息/事件 的服务器
*/
type Server struct {
	Ctx *OffiAccount
}

// EchoStr 服务器接口校验
func (s *Server) EchoStr(writer http.ResponseWriter, request *http.Request) {
	if s.IsSignatureValid(request) {
		echoStr := request.URL.Query().Get("echostr")
		if echoStr != "" {
			io.WriteString(writer, echoStr)
			if s.Ctx.Logger != nil {
				s.Ctx.Logger.Println("echostr ", echoStr)
			}
		}
	}
}

// IsSignatureValid 检查微信通知的签名
func (s *Server) IsSignatureValid(request *http.Request) bool {
	strs := []string{
		request.URL.Query().Get("timestamp"),
		request.URL.Query().Get("nonce"),
		s.Ctx.Config.Token,
	}
	sort.Strings(strs)

	h := sha1.New()
	_, _ = io.WriteString(h, strings.Join(strs, ""))
	signature := fmt.Sprintf("%x", h.Sum(nil))

	if signature == request.URL.Query().Get("signature") {
		return true
	} else {
		return false
	}
}

// ParseXML 解析微信推送过来的消息/事件
func (s *Server) ParseXML(body []byte) (m interface{}, err error) {

	if s.Ctx.Logger != nil {
		s.Ctx.Logger.Println(string(body))
	}

	// 是否加密消息
	encryptMsg := messagetype.EncryptMessage{}
	err = xml.Unmarshal(body, &encryptMsg)
	if err != nil {
		return
	}

	// 需要解密
	if encryptMsg.Encrypt != "" {
		var xmlMsg []byte
		_, xmlMsg, _, err = util.AESDecryptMsg(encryptMsg.Encrypt, s.Ctx.Config.EncodingAESKey)
		if err != nil {
			return
		}
		body = xmlMsg

		if s.Ctx.Logger != nil {
			s.Ctx.Logger.Println("AESDecryptMsg ", string(body))
		}

	}

	message := messagetype.Message{}
	err = xml.Unmarshal(body, &message)
	//fmt.Println(message)
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
	case messagetype.MsgTypeFile:
		msg := messagetype.MessageFile{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case messagetype.MsgTypeEvent:
		return parseEvent(body)
	}
	return
}

// parseEvent 解析微信推送过来的事件
func parseEvent(body []byte) (m interface{}, err error) {
	event := eventtype.Event{}
	err = xml.Unmarshal(body, &event)
	if err != nil {
		return
	}
	switch event.Event {

	// 关注事件
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

	// 菜单事件
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

	// 资质事件
	case eventtype.EventTypeQualificationVerifySuccess:
		msg := eventtype.EventQualificationVerifySuccess{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeQualificationVerifyFail:
		msg := eventtype.EventQualificationVerifyFail{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeNamingVerifySuccess:
		msg := eventtype.EventNamingVerifySuccess{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeNamingVerifyFail:
		msg := eventtype.EventNamingVerifyFail{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeAnnualRenew:
		msg := eventtype.EventAnnualRenew{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeVerifyExpired:
		msg := eventtype.EventVerifyExpired{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil

		// 卡券事件
	case eventtype.EventTypeCardPassChecke:
		msg := eventtype.EventCardPassChecke{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeCardNotPassChecke:
		msg := eventtype.EventCardNotPassChecke{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeUserGetCard:
		msg := eventtype.EventUserGetCard{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeUserGiftingCard:
		msg := eventtype.EventUserGiftingCard{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeUserDelCard:
		msg := eventtype.EventUserDelCard{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeUserConsumeCard:
		msg := eventtype.EventUserConsumeCard{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeUserPayFromPayCell:
		msg := eventtype.EventUserPayFromPayCell{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeUserViewCard:
		msg := eventtype.EventUserViewCard{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeUserEnterSessionFromCard:
		msg := eventtype.EventUserEnterSessionFromCard{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeUpdateMemberCard:
		msg := eventtype.EventUpdateMemberCard{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeCardSkuRemind:
		msg := eventtype.EventCardSkuRemind{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeCardPayOrder:
		msg := eventtype.EventCardPayOrder{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeSubmitMembercardUserInfo:
		msg := eventtype.EventSubmitMembercardUserInfo{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil

		// 导购事件
	case eventtype.EventTypeGuideQrcodeScan:
		msg := eventtype.EventGuideQrcodeScan{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil

		// 模版消息发送任务完成
	case eventtype.EventTypeTemplateSendJobFinish:
		msg := eventtype.EventTemplateSendJobFinish{}
		err = xml.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	}

	return
}

// Response 响应微信消息 (自动判断是否要加密)
func (s *Server) Response(writer http.ResponseWriter, request *http.Request, reply interface{}) (err error) {

	// 如果 开启加密，微信服务器 发过来的请求 带有 如下参数
	//signature=c44d29564aa1d57bd0e274c37baa92bd5b3da5bd
	//&timestamp=1596184957
	//&nonce=1250398014
	//&openid=oEnxesxpxWw-PKkz-vW5IMdfcQaE
	//&encrypt_type=aes
	//&msg_signature=cc24cc38467417603fc3689170e8b0fd3c9bf4a2

	output := []byte("success") // 默认回复
	if reply != nil {
		output, err = xml.Marshal(reply)
		if err != nil {
			return
		}

		// 加密
		if request.URL.Query().Get("encrypt_type") == "aes" {
			message := s.encryptReplyMessage(output)
			output, err = xml.Marshal(message)
			if err != nil {
				return
			}
		}
	}

	_, err = writer.Write(output)

	if s.Ctx.Logger != nil {
		s.Ctx.Logger.Println("Response: ", string(output))
	}

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
