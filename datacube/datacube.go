package datacube

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiGetusersummary          = "/datacube/getusersummary"
	apiGetusercumulate         = "/datacube/getusercumulate"
	apiGetarticlesummary       = "/datacube/getarticlesummary"
	apiGetarticletotal         = "/datacube/getarticletotal"
	apiGetuserread             = "/datacube/getuserread"
	apiGetuserreadhour         = "/datacube/getuserreadhour"
	apiGetusershare            = "/datacube/getusershare"
	apiGetusersharehour        = "/datacube/getusersharehour"
	apiGetupstreammsg          = "/datacube/getupstreammsg"
	apiGetupstreammsghour      = "/datacube/getupstreammsghour"
	apiGetupstreammsgweek      = "/datacube/getupstreammsgweek"
	apiGetupstreammsgmonth     = "/datacube/getupstreammsgmonth"
	apiGetupstreammsgdist      = "/datacube/getupstreammsgdist"
	apiGetupstreammsgdistweek  = "/datacube/getupstreammsgdistweek"
	apiGetupstreammsgdistmonth = "/datacube/getupstreammsgdistmonth"
	apiStat                    = "/publisher/stat"
	apiGetinterfacesummary     = "/datacube/getinterfacesummary"
	apiGetinterfacesummaryhour = "/datacube/getinterfacesummaryhour"
)

/*
/datacube/getusersummary



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/User_Analysis_Data_Interface.html

POST https://api.weixin.qq.com/datacube/getusersummary?access_token=ACCESS_TOKEN
*/
func Getusersummary(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetusersummary, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/datacube/getusercumulate



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/User_Analysis_Data_Interface.html

POST https://api.weixin.qq.com/datacube/getusercumulate?access_token=ACCESS_TOKEN
*/
func Getusercumulate(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetusercumulate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/datacube/getarticlesummary



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html

POST https://api.weixin.qq.com/datacube/getarticlesummary?access_token=ACCESS_TOKEN
*/
func Getarticlesummary(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetarticlesummary, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/datacube/getarticletotal



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html

POST https://api.weixin.qq.com/datacube/getarticletotal?access_token=ACCESS_TOKEN
*/
func Getarticletotal(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetarticletotal, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/datacube/getuserread



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html

POST https://api.weixin.qq.com/datacube/getuserread?access_token=ACCESS_TOKEN
*/
func Getuserread(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetuserread, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/datacube/getuserreadhour



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html

POST https://api.weixin.qq.com/datacube/getuserreadhour?access_token=ACCESS_TOKEN
*/
func Getuserreadhour(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetuserreadhour, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/datacube/getusershare



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html

POST https://api.weixin.qq.com/datacube/getusershare?access_token=ACCESS_TOKEN
*/
func Getusershare(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetusershare, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/datacube/getusersharehour



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html

POST https://api.weixin.qq.com/datacube/getusersharehour?access_token=ACCESS_TOKEN
*/
func Getusersharehour(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetusersharehour, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/datacube/getupstreammsg



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html

POST https://api.weixin.qq.com/datacube/getupstreammsg?access_token=ACCESS_TOKEN
*/
func Getupstreammsg(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetupstreammsg, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/datacube/getupstreammsghour



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html

POST https://api.weixin.qq.com/datacube/getupstreammsghour?access_token=ACCESS_TOKEN
*/
func Getupstreammsghour(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetupstreammsghour, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/datacube/getupstreammsgweek



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html

POST https://api.weixin.qq.com/datacube/getupstreammsgweek?access_token=ACCESS_TOKEN
*/
func Getupstreammsgweek(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetupstreammsgweek, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/datacube/getupstreammsgmonth



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html

POST https://api.weixin.qq.com/datacube/getupstreammsgmonth?access_token=ACCESS_TOKEN
*/
func Getupstreammsgmonth(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetupstreammsgmonth, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/datacube/getupstreammsgdist



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html

POST https://api.weixin.qq.com/datacube/getupstreammsgdist?access_token=ACCESS_TOKEN
*/
func Getupstreammsgdist(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetupstreammsgdist, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/datacube/getupstreammsgdistweek



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html

POST https://api.weixin.qq.com/datacube/getupstreammsgdistweek?access_token=ACCESS_TOKEN
*/
func Getupstreammsgdistweek(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetupstreammsgdistweek, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/datacube/getupstreammsgdistmonth



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html

POST https://api.weixin.qq.com/datacube/getupstreammsgdistmonth?access_token=ACCESS_TOKEN
*/
func Getupstreammsgdistmonth(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetupstreammsgdistmonth, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/publisher/stat



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Ad_Analysis.html

POST https://api.weixin.qq.com/publisher/stat?action=publisher_adpos_general&amp;access_token=ACCESS_TOKEN
*/
func Stat(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiStat, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/datacube/getinterfacesummary



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Analytics_API.html

POST https://api.weixin.qq.com/datacube/getinterfacesummary?access_token=ACCESS_TOKEN
*/
func Getinterfacesummary(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetinterfacesummary, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/datacube/getinterfacesummaryhour



See: https://developers.weixin.qq.com/doc/offiaccount/Analytics/Analytics_API.html

POST https://api.weixin.qq.com/datacube/getinterfacesummaryhour?access_token=ACCESS_TOKEN
*/
func Getinterfacesummaryhour(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetinterfacesummaryhour, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
