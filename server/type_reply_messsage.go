package server

import "encoding/xml"

type CDATA string

func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		string `xml:",cdata"`
	}{string(c)}, start)
}

const (
	ReplyMsgTypeText  = "text"
	ReplyMsgTypeImage = "image"
	ReplyMsgTypeVoice = "voice"
	ReplyMsgTypeVideo = "video"
	ReplyMsgTypeMusic = "music"
	ReplyMsgTypeNews  = "news"
)

type ReplyMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   string
	MsgType      CDATA
}

type ReplyMessageText struct {
	ReplyMessage
	Content CDATA
}

type ReplyMessageImage struct {
	ReplyMessage
	Image struct {
		MediaId CDATA
	}
}

type ReplyMessageVoice struct {
	ReplyMessage
	Voice struct {
		MediaId CDATA
	}
}

type ReplyMessageVideo struct {
	ReplyMessage
	Video struct {
		MediaId     CDATA
		Title       CDATA
		Description CDATA
	}
}
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
