package subscribe

import (
	"bytes"
	"fmt"

	"github.com/fastwego/offiaccount"
)

const subscribeServerUrl = "https://mp.weixin.qq.com/mp/subscribemsg"
const (
	apiSubscribe = "/cgi-bin/message/template/subscribe"
)

/*
获取一次给用户推送一条订阅模板消息的机会

在确保微信公众帐号拥有订阅消息授权的权限的前提下（已认证的公众号即有权限，可登录公众平台在接口权限列表处查看），引导用户在微信客户端打开如下链接：

https://mp.weixin.qq.com/mp/subscribemsg?action=get_confirm&appid=wxaba38c7f163da69b&scene=1000&template_id=1uDxHNXwYQfBmXOfPJcjAS3FynHArD8aWMEFNRGSbCc&redirect_url=http%3a%2f%2fsupport.qq.com&reserved=test#wechat_redirect

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/One-time_subscription_info.html

*/
func GetSubscribeUrl(appid, scene, templateId, redirectUrl, reserved string) (url string) {
	return fmt.Sprintf("%s?action=get_confirm&appid=%s&scene=%s&template_id=%s&redirect_url=%s&reserved=%s#wechat_redirect", subscribeServerUrl, appid, scene, templateId, redirectUrl, reserved)
}

/*
推送订阅模板消息给到授权微信用户

See: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/One-time_subscription_info.html

POST https://api.weixin.qq.com/cgi-bin/message/template/subscribe?access_token=ACCESS_TOKEN
*/
func Subscribe(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSubscribe, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
