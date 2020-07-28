package message

import "github.com/fastwego/offiaccount"

const (
	apiGetCurrentAutoReplyInfo = "/cgi-bin/get_current_autoreply_info"
)

/*
获取公众号的自动回复规则

开发者可以通过该接口，获取公众号当前使用的自动回复规则，包括关注后自动回复、消息自动回复（60分钟内触发一次）、关键词自动回复。

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Getting_Rules_for_Auto_Replies.html

GET https://api.weixin.qq.com/cgi-bin/get_current_autoreply_info?access_token=ACCESS_TOKEN
*/
func GetCurrentAutoReplyInfo() (resp []byte, err error) {
	return offiaccount.HTTPGet(apiGetCurrentAutoReplyInfo)
}
