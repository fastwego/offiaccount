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

const (
	MsgTypeText       = "text"
	MsgTypeImage      = "image"
	MsgTypeVoice      = "voice"
	MsgTypeVideo      = "video"
	MsgTypeShortVideo = "shortvideo"
	MsgTypeLocation   = "location"
	MsgTypeLink       = "link"
	MsgTypeFile       = "file"
	MsgTypeEvent      = "event"
)

type Message struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
}

/*
启用 加密模式 后 收到的 消息格式
<xml>
    <ToUserName><![CDATA[]]></ToUserName>
    <Encrypt><![CDATA[]]></Encrypt>
</xml>
*/
type EncryptMessage struct {
	XMLName    xml.Name `xml:"xml"`
	ToUserName string
	Encrypt    string
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1348831860</CreateTime>
  <MsgType><![CDATA[text]]></MsgType>
  <Content><![CDATA[this is a test]]></Content>
  <MsgId>1234567890123456</MsgId>
</xml>
*/
type MessageText struct {
	Message
	MsgId   string
	Content string
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1348831860</CreateTime>
  <MsgType><![CDATA[image]]></MsgType>
  <PicUrl><![CDATA[this is a url]]></PicUrl>
  <MediaId><![CDATA[media_id]]></MediaId>
  <MsgId>1234567890123456</MsgId>
</xml>
*/
type MessageImage struct {
	Message
	MsgId   string
	PicUrl  string
	MediaId string
}

/*
<xml>
  <ToUserName>< ![CDATA[toUser] ]></ToUserName>
  <FromUserName>< ![CDATA[fromUser] ]></FromUserName>
  <CreateTime>1357290913</CreateTime>
  <MsgType>< ![CDATA[voice] ]></MsgType>
  <MediaId>< ![CDATA[media_id] ]></MediaId>
  <Format>< ![CDATA[Format] ]></Format>
  <Recognition>< ![CDATA[腾讯微信团队] ]></Recognition>
  <MsgId>1234567890123456</MsgId>
</xml>

*/
type MessageVoice struct {
	Message
	MsgId       string
	Format      string
	Recognition string
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1357290913</CreateTime>
  <MsgType><![CDATA[video]]></MsgType>
  <MediaId><![CDATA[media_id]]></MediaId>
  <ThumbMediaId><![CDATA[thumb_media_id]]></ThumbMediaId>
  <MsgId>1234567890123456</MsgId>
</xml>

*/
type MessageVideo struct {
	Message
	MsgId        string
	MediaId      string
	ThumbMediaId string
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1357290913</CreateTime>
  <MsgType><![CDATA[shortvideo]]></MsgType>
  <MediaId><![CDATA[media_id]]></MediaId>
  <ThumbMediaId><![CDATA[thumb_media_id]]></ThumbMediaId>
  <MsgId>1234567890123456</MsgId>
</xml>
*/
type MessageShortVideo struct {
	Message
	MsgId        string
	MediaId      string
	ThumbMediaId string
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1351776360</CreateTime>
  <MsgType><![CDATA[location]]></MsgType>
  <Location_X>23.134521</Location_X>
  <Location_Y>113.358803</Location_Y>
  <Scale>20</Scale>
  <Label><![CDATA[位置信息]]></Label>
  <MsgId>1234567890123456</MsgId>
</xml>

*/
type MessageLocation struct {
	Message
	MsgId      string
	Location_X string
	Location_Y string
	Scale      string
	Label      string
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1351776360</CreateTime>
  <MsgType><![CDATA[link]]></MsgType>
  <Title><![CDATA[公众平台官网链接]]></Title>
  <Description><![CDATA[公众平台官网链接]]></Description>
  <Url><![CDATA[url]]></Url>
  <MsgId>1234567890123456</MsgId>
</xml>
*/
type MessageLink struct {
	Message
	MsgId       string
	Title       string
	Description string
	Url         string
}

/*
<xml>
	<ToUserName><![CDATA[]]></ToUserName>
	<FromUserName><![CDATA[]]></FromUserName>
	<CreateTime>1596098918</CreateTime>
	<MsgType><![CDATA[file]]></MsgType>
	<Title><![CDATA[2020-07-30]]></Title>
	<Description><![CDATA[2.5 KB]]></Description>
	<FileKey><![CDATA[]]></FileKey>
	<FileMd5><![CDATA[]]></FileMd5>
	<FileTotalLen>2515</FileTotalLen>
	<MsgId>22850641670526495</MsgId>
</xml>
*/
type MessageFile struct {
	Message
	MsgId        string
	Title        string
	Description  string
	FileKey      string
	FileMd5      string
	FileTotalLen string
}

type MessageEvent struct {
	Message
	Event string
}
