package user

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiUpdateremark     = "/cgi-bin/user/info/updateremark"
	apiInfo             = "/cgi-bin/user/info"
	apiBatchget         = "/cgi-bin/user/info/batchget"
	apiGet              = "/cgi-bin/user/get"
	apiGetblacklist     = "/cgi-bin/tags/members/getblacklist"
	apiBatchblacklist   = "/cgi-bin/tags/members/batchblacklist"
	apiBatchunblacklist = "/cgi-bin/tags/members/batchunblacklist"
)

/*
设置用户备注名

开发者可以通过该接口对指定用户设置备注名，该接口暂时开放给微信认证的服务号

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/Configuring_user_notes.html

POST https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=ACCESS_TOKEN
*/
func Updateremark(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdateremark, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取用户基本信息(UnionID机制)

在关注者与公众号产生消息交互后，公众号可获得关注者的OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的。对于不同公众号，同一用户的openid不同）。公众号可通过本接口来根据OpenID获取用户基本信息，包括昵称、头像、性别、所在城市、语言和关注时间

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId

GET https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
*/
func Info() (resp []byte, err error) {
	return offiaccount.HTTPGet(apiInfo)
}

/*
批量获取用户基本信息

开发者可通过该接口来批量获取用户基本信息。最多支持一次拉取100条

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId

POST https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=ACCESS_TOKEN
*/
func Batchget(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiBatchget, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
批量用户列表

公众号可通过本接口来获取帐号的关注者列表，关注者列表由一串OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成。一次拉取调用最多拉取10000个关注者的OpenID，可以通过多次拉取的方式来满足需求

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/Getting_a_User_List.html

GET https://api.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&next_openid=NEXT_OPENID
*/
func Get() (resp []byte, err error) {
	return offiaccount.HTTPGet(apiGet)
}

/*
获取公众号的黑名单列表

公众号可通过该接口来获取帐号的黑名单列表，黑名单列表由一串 OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html

POST https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token=ACCESS_TOKEN
*/
func Getblacklist(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetblacklist, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
拉黑用户

公众号可通过该接口来拉黑一批用户，黑名单列表由一串 OpenID （加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html

POST https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist?access_token=ACCESS_TOKEN
*/
func Batchblacklist(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiBatchblacklist, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
取消拉黑用户

公众号可通过该接口来取消拉黑一批用户，黑名单列表由一串OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html

POST https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist?access_token=ACCESS_TOKEN
*/
func Batchunblacklist(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiBatchunblacklist, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
