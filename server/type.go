package server

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
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	MsgId        string
}

type TextMessage struct {
	Message
	Content string
}

type ImageMessage struct {
	Message
	PicUrl  string
	MediaId string
}

type VoiceMessage struct {
	Message
	Format      string
	Recognition string
}

type VideoMessage struct {
	Message
	MediaId      string
	ThumbMediaId string
}
type ShortVideoMessage struct {
	Message
	MediaId      string
	ThumbMediaId string
}

type LocationMessage struct {
	Message
	Location_X string
	Location_Y string
	Scale      string
	Label      string
}

type LinkMessage struct {
	Message
	Title       string
	Description string
	Url         string
}

type EventMessage struct {
	Message
	Event string
}
