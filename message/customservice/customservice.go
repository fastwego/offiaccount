package customservice

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/fastwego/offiaccount"
)

const (
	apiAdd           = "/customservice/kfaccount/add"
	apiUpdate        = "/customservice/kfaccount/update"
	apiDel           = "/customservice/kfaccount/del"
	apiUploadHeadImg = "/customservice/kfaccount/uploadheadimg"
	apiGetKFList     = "/cgi-bin/customservice/getkflist"
	apiSend          = "/cgi-bin/message/custom/send"
	apiTyping        = "/cgi-bin/message/custom/typing"
)

/*
添加客服帐号

开发者可以通过本接口为公众号添加客服账号，每个公众号最多添加100个客服账号

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html

POST https://api.weixin.qq.com/customservice/kfaccount/add?access_token=ACCESS_TOKEN
*/
func Add(kf KF) (resp []byte, err error) {
	data, err := json.Marshal(kf)
	if err != nil {
		return
	}
	return offiaccount.HTTPPost(apiAdd, bytes.NewBuffer(data), offiaccount.ContentTypeApplicationJson)
}

/*
修改客服帐号

开发者可以通过本接口为公众号修改客服账号

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html

POST https://api.weixin.qq.com/customservice/kfaccount/update?access_token=ACCESS_TOKEN
*/
func Update(kf KF) (resp []byte, err error) {
	data, err := json.Marshal(kf)
	if err != nil {
		return
	}
	return offiaccount.HTTPPost(apiUpdate, bytes.NewBuffer(data), offiaccount.ContentTypeApplicationJson)
}

/*
删除客服帐号

开发者可以通过该接口为公众号删除客服帐号

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html

POST https://api.weixin.qq.com/customservice/kfaccount/del?access_token=ACCESS_TOKEN
*/
func Del(kf KF) (resp []byte, err error) {
	data, err := json.Marshal(kf)
	if err != nil {
		return
	}
	return offiaccount.HTTPPost(apiDel, bytes.NewBuffer(data), offiaccount.ContentTypeApplicationJson)
}

/*
设置客服帐号的头像

开发者可调用本接口来上传图片作为客服人员的头像，头像图片文件必须是jpg格式，推荐使用640*640大小的图片以达到最佳效果

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html

POST http://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=ACCESS_TOKEN&kf_account=KFACCOUNT

调用示例：使用curl命令，用FORM表单方式上传一个多媒体文件，curl命令的具体用法请自行了解
*/
func UploadHeadImg(kf KF, imagePath string) (resp []byte, err error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("image", path.Base(imagePath))
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
	return offiaccount.HTTPPost(apiUploadHeadImg+"?kf_account="+kf.KfAccount, r, m.FormDataContentType())
}

/*
获取所有客服账号

开发者通过本接口，获取公众号中所设置的客服基本信息，包括客服工号、客服昵称、客服登录账号

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html

GET https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=ACCESS_TOKEN
*/
func GetKFList(kf KF) (resp []byte, err error) {
	return offiaccount.HTTPGet(apiGetKFList)
}

/*
客服接口-发消息

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html

POST https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN
*/
func Send(message string) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSend, strings.NewReader(message), offiaccount.ContentTypeApplicationJson)
}

/*
客服输入状态

开发者可通过调用“客服输入状态”接口，返回客服当前输入状态给用户

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html

POST https://api.weixin.qq.com/cgi-bin/message/custom/typing?access_token=ACCESS_TOKEN
*/
func Typing(touser string, command string) (resp []byte, err error) {
	data, err := json.Marshal(struct {
		Touser  string `json:"touser"`
		Command string `json:"command"`
	}{Touser: touser, Command: command})
	if err != nil {
		return
	}
	return offiaccount.HTTPPost(apiTyping, bytes.NewBuffer(data), offiaccount.ContentTypeApplicationJson)
}
