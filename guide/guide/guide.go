package guide

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiGetGuideAcct              = "/cgi-bin/guide/getguideacct"
	apiAddGuideAcct              = "/cgi-bin/guide/addguideacct"
	apiUpdateGuideAcct           = "/cgi-bin/guide/updateguideacct"
	apiDelGuideAcct              = "/cgi-bin/guide/delguideacct"
	apiGetGuideAcctList          = "/cgi-bin/guide/getguideacctlist"
	apiGuideCreateQrCode         = "/cgi-bin/guide/guidecreateqrcode"
	apiGetGuideBuyerChatRecord   = "/cgi-bin/guide/getguidebuyerchatrecord"
	apiSetGuideConfig            = "/cgi-bin/guide/setguideconfig"
	apiGetGuideConfig            = "/cgi-bin/guide/getguideconfig"
	apiSetGuideAcctConfig        = "/cgi-bin/guide/setguideacctconfig"
	apiGetGuideAcctConfig        = "/cgi-bin/guide/getguideacctconfig"
	apiPushShowWxaPathMenu       = "/cgi-bin/guide/pushshowwxapathmenu"
	apiNewGuideGroup             = "/cgi-bin/guide/newguidegroup"
	apiGetGuideGroupList         = "/cgi-bin/guide/getguidegrouplist"
	apiGetGroupInfo              = "/cgi-bin/guide/getgroupinfo"
	apiAddGuide2GuideGroup       = "/cgi-bin/guide/addguide2guidegroup"
	apiDelGuide2GuideGroup       = "/cgi-bin/guide/delguide2guidegroup"
	apiGetGroupByGuide           = "/cgi-bin/guide/getgroupbyguide"
	apiDelGuideGroup             = "/cgi-bin/guide/delguidegroup"
	apiGetGuideBuyerRelationList = "/cgi-bin/guide/getguidebuyerrelationlist"
)

/*
获取顾问信息



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideAcct.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguideacct?access_token=ACCESS_TOKEN
*/
func GetGuideAcct(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetGuideAcct, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
为服务号添加顾问



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.addGuideAcct.html

POST https://api.weixin.qq.com/cgi-bin/guide/addguideacct?access_token=ACCESS_TOKEN
*/
func AddGuideAcct(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAddGuideAcct, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
修改顾问的昵称或头像



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.updateGuideAcct.html

POST https://api.weixin.qq.com/cgi-bin/guide/updateguideacct?access_token=ACCESS_TOKEN
*/
func UpdateGuideAcct(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdateGuideAcct, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
删除顾问



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.delGuideAcct.html

POST https://api.weixin.qq.com/cgi-bin/guide/delguideacct?access_token=ACCESS_TOKEN
*/
func DelGuideAcct(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelGuideAcct, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取服务号顾问列表



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideAcctList.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguideacctlist?access_token=ACCESS_TOKEN
*/
func GetGuideAcctList(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetGuideAcctList, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
生成顾问二维码



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.guideCreateQrCode.html

POST https://api.weixin.qq.com/cgi-bin/guide/guidecreateqrcode?access_token=ACCESS_TOKEN
*/
func GuideCreateQrCode(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGuideCreateQrCode, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/cgi-bin/guide/getguidebuyerchatrecord



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideBuyerChatRecord.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidebuyerchatrecord?access_token=ACCESS_TOKEN
*/
func GetGuideBuyerChatRecord(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetGuideBuyerChatRecord, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
设置快捷回复与关注自动回复



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.setGuideConfig.html

POST https://api.weixin.qq.com/cgi-bin/guide/setguideconfig?access_token=ACCESS_TOKEN
*/
func SetGuideConfig(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSetGuideConfig, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取快捷回复与关注自动回复



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideConfig.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguideconfig?access_token=ACCESS_TOKEN
*/
func GetGuideConfig(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetGuideConfig, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
为服务号设置敏感词与离线自动回复



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.setGuideAcctConfig.html

POST https://api.weixin.qq.com/cgi-bin/guide/setguideacctconfig?access_token=ACCESS_TOKEN
*/
func SetGuideAcctConfig(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSetGuideAcctConfig, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取离线自动回复与敏感词



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideAcctConfig.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguideacctconfig?access_token=ACCESS_TOKEN
*/
func GetGuideAcctConfig(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetGuideAcctConfig, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
允许微信用户复制小程序页面路径



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.pushShowWxaPathMenu.html

POST https://api.weixin.qq.com/cgi-bin/guide/pushshowwxapathmenu?access_token=ACCESS_TOKEN
*/
func PushShowWxaPathMenu(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiPushShowWxaPathMenu, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
新建顾问分组



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.newGuideGroup.html

POST https://api.weixin.qq.com/cgi-bin/guide/newguidegroup?access_token=ACCESS_TOKEN
*/
func NewGuideGroup(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiNewGuideGroup, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取服务号下所有顾问分组的列表



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideGroupList.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidegrouplist?access_token=ACCESS_TOKEN
*/
func GetGuideGroupList(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetGuideGroupList, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取指定顾问分组信息，以及分组内顾问信息



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGroupInfo.html

POST https://api.weixin.qq.com/cgi-bin/guide/getgroupinfo?access_token=ACCESS_TOKEN
*/
func GetGroupInfo(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetGroupInfo, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
分组内添加顾问



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.addGuide2GuideGroup.html

POST https://api.weixin.qq.com/cgi-bin/guide/addguide2guidegroup?access_token=ACCESS_TOKEN
*/
func AddGuide2GuideGroup(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAddGuide2GuideGroup, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
分组内删除顾问



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.delGuide2GuideGroup.html

POST https://api.weixin.qq.com/cgi-bin/guide/delguide2guidegroup?access_token=ACCESS_TOKEN
*/
func DelGuide2GuideGroup(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelGuide2GuideGroup, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取顾问所在分组



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGroupByGuide.html

POST https://api.weixin.qq.com/cgi-bin/guide/getgroupbyguide?access_token=ACCESS_TOKEN
*/
func GetGroupByGuide(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetGroupByGuide, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
删除指定顾问分组



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.delGuideGroup.html

POST https://api.weixin.qq.com/cgi-bin/guide/delguidegroup?access_token=ACCESS_TOKEN
*/
func DelGuideGroup(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelGuideGroup, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取顾问的客户列表



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.getGuideBuyerRelationList.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidebuyerrelationlist?access_token=ACCESS_TOKEN
*/
func GetGuideBuyerRelationList(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetGuideBuyerRelationList, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
