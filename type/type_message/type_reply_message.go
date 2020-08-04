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

package type_message

import "encoding/xml"

type CDATA string

func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

const (
	ReplyMsgTypeText                    = "text"
	ReplyMsgTypeImage                   = "image"
	ReplyMsgTypeVoice                   = "voice"
	ReplyMsgTypeVideo                   = "video"
	ReplyMsgTypeMusic                   = "music"
	ReplyMsgTypeNews                    = "news"
	ReplyMsgTypeTransferCustomerService = "transfer_customer_service" // 消息转发到(指定)客服
)

type ReplyMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   string
	MsgType      CDATA
}

/*
加密处理后 的 回复 消息体
<xml>
<Encrypt></Encrypt>
<MsgSignature></MsgSignature>
<TimeStamp></TimeStamp>
<Nonce></Nonce>
</xml>
*/
type ReplyEncryptMessage struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      string
	MsgSignature string
	TimeStamp    string
	Nonce        string
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>12345678</CreateTime>
  <MsgType><![CDATA[text]]></MsgType>
  <Content><![CDATA[你好]]></Content>
</xml>
*/
type ReplyMessageText struct {
	ReplyMessage
	Content CDATA
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>12345678</CreateTime>
  <MsgType><![CDATA[image]]></MsgType>
  <Image>
    <MediaId><![CDATA[media_id]]></MediaId>
  </Image>
</xml>
*/
type ReplyMessageImage struct {
	ReplyMessage
	Image struct {
		MediaId CDATA
	}
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>12345678</CreateTime>
  <MsgType><![CDATA[voice]]></MsgType>
  <Voice>
    <MediaId><![CDATA[media_id]]></MediaId>
  </Voice>
</xml>
*/
type ReplyMessageVoice struct {
	ReplyMessage
	Voice struct {
		MediaId CDATA
	}
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>12345678</CreateTime>
  <MsgType><![CDATA[video]]></MsgType>
  <Video>
    <MediaId><![CDATA[media_id]]></MediaId>
    <Title><![CDATA[title]]></Title>
    <Description><![CDATA[description]]></Description>
  </Video>
</xml>
*/
type ReplyMessageVideo struct {
	ReplyMessage
	Video struct {
		MediaId     CDATA
		Title       CDATA
		Description CDATA
	}
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>12345678</CreateTime>
  <MsgType><![CDATA[music]]></MsgType>
  <Music>
    <Title><![CDATA[TITLE]]></Title>
    <Description><![CDATA[DESCRIPTION]]></Description>
    <MusicUrl><![CDATA[MUSIC_Url]]></MusicUrl>
    <HQMusicUrl><![CDATA[HQ_MUSIC_Url]]></HQMusicUrl>
    <ThumbMediaId><![CDATA[media_id]]></ThumbMediaId>
  </Music>
</xml>
*/
type ReplyMessageMusic struct {
	ReplyMessage
	Music struct {
		MediaId      CDATA
		Title        CDATA
		Description  CDATA
		MusicUrl     CDATA
		HQMusicUrl   CDATA
		ThumbMediaId CDATA
	}
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>12345678</CreateTime>
  <MsgType><![CDATA[news]]></MsgType>
  <ArticleCount>1</ArticleCount>
  <Articles>
    <item>
      <Title><![CDATA[title1]]></Title>
      <Description><![CDATA[description1]]></Description>
      <PicUrl><![CDATA[picurl]]></PicUrl>
      <Url><![CDATA[url]]></Url>
    </item>
  </Articles>
</xml>

*/
type ReplyMessageNews struct {
	ReplyMessage
	ArticleCount string
	Articles     struct {
		Item []struct {
			Title       CDATA
			Description CDATA
			PicUrl      CDATA
			URL         CDATA
		}
	}
}

/*
<xml>
  <ToUserName><![CDATA[touser]]></ToUserName>
  <FromUserName><![CDATA[fromuser]]></FromUserName>
  <CreateTime>1399197672</CreateTime>
  <MsgType><![CDATA[transfer_customer_service]]></MsgType>
  <TransInfo>
    <KfAccount><![CDATA[test1@test]]></KfAccount>
  </TransInfo>
</xml>
*/

type ReplyMessageTransferCustomerService struct {
	ReplyMessage
	TransInfo struct {
		KfAccount CDATA
	}
}
