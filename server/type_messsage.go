package server

import "encoding/xml"

const (
	MsgTypeText       = "text"
	MsgTypeImage      = "image"
	MsgTypeVoice      = "voice"
	MsgTypeVideo      = "video"
	MsgTypeShortVideo = "shortvideo"
	MsgTypeLocation   = "location"
	MsgTypeLink       = "link"
	MsgTypeEvent      = "event"
)

type Message struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	MsgId        string
}

type MessageText struct {
	Message
	Content string
}

type MessageImage struct {
	Message
	PicUrl  string
	MediaId string
}

type MessageVoice struct {
	Message
	Format      string
	Recognition string
}

type MessageVideo struct {
	Message
	MediaId      string
	ThumbMediaId string
}

type MessageShortVideo struct {
	Message
	MediaId      string
	ThumbMediaId string
}

type MessageLocation struct {
	Message
	Location_X string
	Location_Y string
	Scale      string
	Label      string
}

type MessageLink struct {
	Message
	Title       string
	Description string
	Url         string
}

type MessageEvent struct {
	Message
	Event string
}
