package template

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiSetIndustry           = "/cgi-bin/template/api_set_industry"
	apiGetIndustry           = "/cgi-bin/template/get_industry"
	apiAddTemplate           = "/cgi-bin/template/api_add_template"
	apiGetAllPrivateTemplate = "/cgi-bin/template/get_all_private_template"
	apiDelPrivateTemplate    = "/cgi-bin/template/del_private_template"
	apiSend                  = "/cgi-bin/message/template/send"
)

/*
设置所属行业

设置行业可在微信公众平台后台完成，每月可修改行业1次，帐号仅可使用所属行业中相关的模板，为方便第三方开发者，提供通过接口调用的方式来修改账号所属行业

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html

POST https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token=ACCESS_TOKEN
*/
func SetIndustry(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSetIndustry, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取设置的行业信息

获取帐号设置的行业信息。可登录微信公众平台，在公众号后台中查看行业信息。为方便第三方开发者，提供通过接口调用的方式来获取帐号所设置的行业信息

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html

GET https://api.weixin.qq.com/cgi-bin/template/get_industry?access_token=ACCESS_TOKEN
*/
func GetIndustry() (resp []byte, err error) {
	return offiaccount.HTTPGet(apiGetIndustry)
}

/*
获得模板ID

从行业模板库选择模板到帐号后台，获得模板ID的过程可在微信公众平台后台完成。为方便第三方开发者，提供通过接口调用的方式来获取模板ID

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html

POST https://api.weixin.qq.com/cgi-bin/template/api_add_template?access_token=ACCESS_TOKEN
*/
func AddTemplate(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAddTemplate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取模板列表

获取已添加至帐号下所有模板列表，可在微信公众平台后台中查看模板列表信息。为方便第三方开发者，提供通过接口调用的方式来获取帐号下所有模板信息

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html

GET https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=ACCESS_TOKEN
*/
func GetAllPrivateTemplate() (resp []byte, err error) {
	return offiaccount.HTTPGet(apiGetAllPrivateTemplate)
}

/*
删除模板

删除模板可在微信公众平台后台完成，为方便第三方开发者，提供通过接口调用的方式来删除某帐号下的模板

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html

POST https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token=ACCESS_TOKEN
*/
func DelPrivateTemplate(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelPrivateTemplate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
发送模板消息

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html

POST https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=ACCESS_TOKEN
*/
func Send(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSend, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
