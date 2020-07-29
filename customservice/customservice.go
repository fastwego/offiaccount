package customservice

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path"

	"github.com/fastwego/offiaccount"
)

const (
	apiKfaccountAdd         = "/customservice/kfaccount/add"
	apiKfaccountUpdate      = "/customservice/kfaccount/update"
	apiKfaccountDel         = "/customservice/kfaccount/del"
	apiUploadHeadImg        = "/customservice/kfaccount/uploadheadimg"
	apiGetKfList            = "/cgi-bin/customservice/getkflist"
	apiSend                 = "/cgi-bin/message/custom/send"
	apiTyping               = "/cgi-bin/message/custom/typing"
	apiGetOnlineKfList      = "/cgi-bin/customservice/getonlinekflist"
	apiInviteWorker         = "/customservice/kfaccount/inviteworker"
	apiKfSessionCreate      = "/customservice/kfsession/create"
	apiKfSessionGet         = "/customservice/kfsession/getsession"
	apiKfSessionGetList     = "/customservice/kfsession/getsessionlist"
	apiKfSessionGetWaitCase = "/customservice/kfsession/getwaitcase"
	apiGetMsgList           = "/customservice/msgrecord/getmsglist"
)

/*
添加客服帐号

开发者可以通过本接口为公众号添加客服账号，每个公众号最多添加100个客服账号

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html

POST https://api.weixin.qq.com/customservice/kfaccount/add?access_token=ACCESS_TOKEN
*/
func KfaccountAdd(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiKfaccountAdd, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
修改客服帐号

开发者可以通过本接口为公众号修改客服账号

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html

POST https://api.weixin.qq.com/customservice/kfaccount/update?access_token=ACCESS_TOKEN
*/
func KfaccountUpdate(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiKfaccountUpdate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
删除客服帐号

开发者可以通过该接口为公众号删除客服帐号

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html

POST https://api.weixin.qq.com/customservice/kfaccount/del?access_token=ACCESS_TOKEN
*/
func KfaccountDel(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiKfaccountDel, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
设置客服帐号的头像

开发者可调用本接口来上传图片作为客服人员的头像，头像图片文件必须是jpg格式，推荐使用640*640大小的图片以达到最佳效果

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html

POST(@media) http://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=ACCESS_TOKEN&kf_account=KFACCOUNT
*/
func UploadHeadImg(media string, params url.Values) (resp []byte, err error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("media", path.Base(media))
		if err != nil {
			return
		}
		file, err := os.Open(media)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}

	}()
	return offiaccount.HTTPPost(apiUploadHeadImg+"?"+params.Encode(), r, m.FormDataContentType())
}

/*
获取所有客服账号

开发者通过本接口，获取公众号中所设置的客服基本信息，包括客服工号、客服昵称、客服登录账号

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html

GET https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=ACCESS_TOKEN
*/
func GetKfList() (resp []byte, err error) {
	return offiaccount.HTTPGet(apiGetKfList)
}

/*
发消息

发送文本/图片/语言/语音/视频/音乐/图文/菜单/卡券... 消息

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html

POST https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN
*/
func Send(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSend, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
客服输入状态

开发者可通过调用“客服输入状态”接口，返回客服当前输入状态给用户

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html

POST https://api.weixin.qq.com/cgi-bin/message/custom/typing?access_token=ACCESS_TOKEN
*/
func Typing(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiTyping, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取客服基本信息



See: https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html

GET https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist?access_token=ACCESS_TOKEN
*/
func GetOnlineKfList() (resp []byte, err error) {
	return offiaccount.HTTPGet(apiGetOnlineKfList)
}

/*
邀请绑定客服帐号

新添加的客服帐号是不能直接使用的，只有客服人员用微信号绑定了客服账号后，方可登录Web客服进行操作。此接口发起一个绑定邀请到客服人员微信号，客服人员需要在微信客户端上用该微信号确认后帐号才可用。尚未绑定微信号的帐号可以进行绑定邀请操作，邀请未失效时不能对该帐号进行再次绑定微信号邀请

See: https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html

POST https://api.weixin.qq.com/customservice/kfaccount/inviteworker?access_token=ACCESS_TOKEN
*/
func InviteWorker(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiInviteWorker, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
创建会话

此接口在客服和用户之间创建一个会话，如果该客服和用户会话已存在，则直接返回0。指定的客服帐号必须已经绑定微信号且在线

See: https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html

POST https://api.weixin.qq.com/customservice/kfsession/create?access_token=ACCESS_TOKEN
*/
func KfSessionCreate(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiKfSessionCreate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取客户会话状态



See: https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html

GET https://api.weixin.qq.com/customservice/kfsession/getsession?access_token=ACCESS_TOKEN&openid=OPENID
*/
func KfSessionGet(params url.Values) (resp []byte, err error) {
	return offiaccount.HTTPGet(apiKfSessionGet + "?" + params.Encode())
}

/*
获取客服会话列表



See: https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html

GET https://api.weixin.qq.com/customservice/kfsession/getsessionlist?access_token=ACCESS_TOKEN&kf_account=KFACCOUNT
*/
func KfSessionGetList(params url.Values) (resp []byte, err error) {
	return offiaccount.HTTPGet(apiKfSessionGetList + "?" + params.Encode())
}

/*
获取未接入会话列表



See: https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html

GET https://api.weixin.qq.com/customservice/kfsession/getwaitcase?access_token=ACCESS_TOKEN
*/
func KfSessionGetWaitCase() (resp []byte, err error) {
	return offiaccount.HTTPGet(apiKfSessionGetWaitCase)
}

/*
获取聊天记录

此接口返回的聊天记录中，对于图片、语音、视频，分别展示成文本格式的[image]、[voice]、[video]

See: https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Obtain_chat_transcript.html

POST https://api.weixin.qq.com/customservice/msgrecord/getmsglist
*/
func GetMsgList(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetMsgList, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
