package datacube

import (
	"bytes"
	"net/url"

	"github.com/fastwego/offiaccount"
)

const (
	apiGetUserSummary          = "/datacube/getusersummary"
	apiGetUserCumulate         = "/datacube/getusercumulate"
	apiGetArticleSummary       = "/datacube/getarticlesummary"
	apiGetArticleTotal         = "/datacube/getarticletotal"
	apiGetUserRead             = "/datacube/getuserread"
	apiGetUserReadHour         = "/datacube/getuserreadhour"
	apiGetUserShare            = "/datacube/getusershare"
	apiGetUserShareHour        = "/datacube/getusersharehour"
	apiGetUpstreamMsg          = "/datacube/getupstreammsg"
	apiGetUpstreamMsgHour      = "/datacube/getupstreammsghour"
	apiGetUpstreamMsgWeek      = "/datacube/getupstreammsgweek"
	apiGetUpstreamMsgMonth     = "/datacube/getupstreammsgmonth"
	apiGetUpstreamMsgDist      = "/datacube/getupstreammsgdist"
	apiGetUpstreamMsgDistWeek  = "/datacube/getupstreammsgdistweek"
	apiGetUpstreamMsgDistMonth = "/datacube/getupstreammsgdistmonth"
	apiPublisherStat           = "/publisher/stat"
	apiGetInterfaceSummary     = "/datacube/getinterfacesummary"
	apiGetInterfaceSummaryHour = "/datacube/getinterfacesummaryhour"
)

/*
获取用户增减数据



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/User_Analysis_Data_Interface.html

POST https://api.weixin.qq.com/datacube/getusersummary?access_token=ACCESS_TOKEN
*/
func GetUserSummary(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetUserSummary, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取累计用户数据



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/User_Analysis_Data_Interface.html

POST https://api.weixin.qq.com/datacube/getusercumulate?access_token=ACCESS_TOKEN
*/
func GetUserCumulate(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetUserCumulate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取图文群发每日数据



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html

POST https://api.weixin.qq.com/datacube/getarticlesummary?access_token=ACCESS_TOKEN
*/
func GetArticleSummary(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetArticleSummary, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取图文群发总数据



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html

POST https://api.weixin.qq.com/datacube/getarticletotal?access_token=ACCESS_TOKEN
*/
func GetArticleTotal(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetArticleTotal, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取图文统计数据



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html

POST https://api.weixin.qq.com/datacube/getuserread?access_token=ACCESS_TOKEN
*/
func GetUserRead(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetUserRead, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取图文统计分时数据



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html

POST https://api.weixin.qq.com/datacube/getuserreadhour?access_token=ACCESS_TOKEN
*/
func GetUserReadHour(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetUserReadHour, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取图文分享转发数据



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html

POST https://api.weixin.qq.com/datacube/getusershare?access_token=ACCESS_TOKEN
*/
func GetUserShare(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetUserShare, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取图文分享转发分时数据



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html

POST https://api.weixin.qq.com/datacube/getusersharehour?access_token=ACCESS_TOKEN
*/
func GetUserShareHour(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetUserShareHour, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取消息发送概况数据



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html

POST https://api.weixin.qq.com/datacube/getupstreammsg?access_token=ACCESS_TOKEN
*/
func GetUpstreamMsg(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetUpstreamMsg, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取消息分送分时数据



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html

POST https://api.weixin.qq.com/datacube/getupstreammsghour?access_token=ACCESS_TOKEN
*/
func GetUpstreamMsgHour(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetUpstreamMsgHour, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取消息发送周数据



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html

POST https://api.weixin.qq.com/datacube/getupstreammsgweek?access_token=ACCESS_TOKEN
*/
func GetUpstreamMsgWeek(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetUpstreamMsgWeek, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取消息发送月数据



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html

POST https://api.weixin.qq.com/datacube/getupstreammsgmonth?access_token=ACCESS_TOKEN
*/
func GetUpstreamMsgMonth(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetUpstreamMsgMonth, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取消息发送分布数据



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html

POST https://api.weixin.qq.com/datacube/getupstreammsgdist?access_token=ACCESS_TOKEN
*/
func GetUpstreamMsgDist(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetUpstreamMsgDist, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取消息发送分布周数据



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html

POST https://api.weixin.qq.com/datacube/getupstreammsgdistweek?access_token=ACCESS_TOKEN
*/
func GetUpstreamMsgDistWeek(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetUpstreamMsgDistWeek, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取消息发送分布月数据



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html

POST https://api.weixin.qq.com/datacube/getupstreammsgdistmonth?access_token=ACCESS_TOKEN
*/
func GetUpstreamMsgDistMonth(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetUpstreamMsgDistMonth, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取公众号分广告位数据/公众号返佣商品数据/公众号结算收入数据及结算主体信息



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Ad_Analysis.html

GET https://api.weixin.qq.com/publisher/stat?action=publisher_adpos_general&access_token=ACCESS_TOKEN
*/
func PublisherStat(params url.Values) (resp []byte, err error) {
	return offiaccount.HTTPGet(apiPublisherStat + "?" + params.Encode())
}

/*
获取接口分析数据



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Analytics_API.html

POST https://api.weixin.qq.com/datacube/getinterfacesummary?access_token=ACCESS_TOKEN
*/
func GetInterfaceSummary(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetInterfaceSummary, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取接口分析分时数据



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Analytics_API.html

POST https://api.weixin.qq.com/datacube/getinterfacesummaryhour?access_token=ACCESS_TOKEN
*/
func GetInterfaceSummaryHour(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetInterfaceSummaryHour, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
