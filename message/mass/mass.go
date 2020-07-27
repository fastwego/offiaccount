package mass

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path"

	"github.com/fastwego/offiaccount"
)

const (
	apiUploadImg  = "/cgi-bin/media/uploadimg"
	apiUploadNews = "/cgi-bin/media/uploadnews"
	apiSendAll    = "/cgi-bin/message/mass/sendall"
	apiSend       = "/cgi-bin/message/mass/send"
	apiDelete     = "/cgi-bin/message/mass/delete"
	apiPreview    = "/cgi-bin/message/mass/preview"
	apiGet        = "/message/mass/get"
	apiSpeedGet   = "/cgi-bin/message/mass/speed/get"
	apiSpeedSet   = "/cgi-bin/message/mass/speed/set"
)

/*
上传图文消息内的图片获取URL

【订阅号与服务号认证后均可用】

请注意，本接口所上传的图片不占用公众号的素材库中图片数量的5000个的限制。图片仅支持jpg/png格式，大小必须在1MB以下

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html#0

POST https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN
*/
func UploadHeadImg(imagePath string) (resp []byte, err error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("media", path.Base(imagePath))
		if err != nil {
			return
		}
		file, err := os.Open(imagePath)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}
	}()
	return offiaccount.HTTPPost(apiUploadImg, r, m.FormDataContentType())
}

/*
上传图文消息素材

【订阅号与服务号认证后均可用】

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html#0

POST https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token=ACCESS_TOKEN
*/
func UploadNews(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUploadNews, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
根据标签进行群发

【订阅号与服务号认证后均可用】

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html#0

POST https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=ACCESS_TOKEN
*/
func SendAll(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSendAll, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
根据OpenID列表群发

【订阅号不可用，服务号认证后可用】

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html#0

POST https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=ACCESS_TOKEN
*/
func Send(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSend, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
删除群发

【订阅号与服务号认证后均可用】

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html#0

POST https://api.weixin.qq.com/cgi-bin/message/mass/delete?access_token=ACCESS_TOKEN
*/
func Delete(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelete, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
预览接口

【订阅号与服务号认证后均可用】

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html#0

POST https://api.weixin.qq.com/cgi-bin/message/mass/preview?access_token=ACCESS_TOKEN
*/
func Preview(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiPreview, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询群发消息发送状态

【订阅号与服务号认证后均可用】


See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html#0

 POST https://api.weixin.qq.com/cgi-bin/message/mass/get?access_token=ACCESS_TOKEN
*/
func Get(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取群发速度

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html#0

POST https://api.weixin.qq.com/cgi-bin/message/mass/speed/get?access_token=ACCESS_TOKEN
*/
func SpeedGet(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSpeedGet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
设置群发速度

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html#0

POST https://api.weixin.qq.com/cgi-bin/message/mass/speed/set?access_token=ACCESS_TOKEN
*/
func SpeedSet(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSpeedSet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
