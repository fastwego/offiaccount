package mass

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiMediaUploadNews  = "/cgi-bin/media/uploadnews"
	apiSendAll          = "/cgi-bin/message/mass/sendall"
	apiMediaUploadVideo = "/cgi-bin/media/uploadvideo"
	apiSend             = "/cgi-bin/message/mass/send"
	apiDelete           = "/cgi-bin/message/mass/delete"
	apiPreview          = "/cgi-bin/message/mass/preview"
	apiGet              = "/cgi-bin/message/mass/get"
	apiSpeedGet         = "/cgi-bin/message/mass/speed/get"
	apiSpeedSet         = "/cgi-bin/message/mass/speed/set"
)

/*
上传图文消息素材



See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html

POST https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token=ACCESS_TOKEN
*/
func MediaUploadNews(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiMediaUploadNews, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
根据标签进行群发



See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html

POST https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=ACCESS_TOKEN
*/
func SendAll(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSendAll, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
上传视频素材

群发 视频 的 media_id 需通过此接口特别地得到

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html

POST https://api.weixin.qq.com/cgi-bin/media/uploadvideo?access_token=ACCESS_TOKEN
*/
func MediaUploadVideo(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiMediaUploadVideo, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
根据OpenID列表群发



See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html

POST https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=ACCESS_TOKEN
*/
func Send(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSend, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
删除群发



See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html

POST https://api.weixin.qq.com/cgi-bin/message/mass/delete?access_token=ACCESS_TOKEN
*/
func Delete(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelete, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
预览

开发者可通过该接口发送消息给指定用户，在手机端查看消息的样式和排版。为了满足第三方平台开发者的需求，在保留对openID预览能力的同时，增加了对指定微信号发送预览的能力，但该能力每日调用次数有限制（100次），请勿滥用

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html

POST https://api.weixin.qq.com/cgi-bin/message/mass/preview?access_token=ACCESS_TOKEN
*/
func Preview(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiPreview, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询群发消息发送状态



See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html

POST https://api.weixin.qq.com/cgi-bin/message/mass/get?access_token=ACCESS_TOKEN
*/
func Get(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取群发速度



See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html

POST https://api.weixin.qq.com/cgi-bin/message/mass/speed/get?access_token=ACCESS_TOKEN
*/
func SpeedGet(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSpeedGet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
设置群发速度



See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html

POST https://api.weixin.qq.com/cgi-bin/message/mass/speed/set?access_token=ACCESS_TOKEN
*/
func SpeedSet(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSpeedSet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
