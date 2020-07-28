package poi

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiAddpoi        = "/cgi-bin/poi/addpoi"
	apiGetpoi        = "/cgi-bin/poi/getpoi"
	apiGetpoilist    = "/cgi-bin/poi/getpoilist"
	apiUpdatepoi     = "/cgi-bin/poi/updatepoi"
	apiDelpoi        = "/cgi-bin/poi/delpoi"
	apiGetwxcategory = "/cgi-bin/poi/getwxcategory"
)

/*
/cgi-bin/poi/addpoi



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html

POST http://api.weixin.qq.com/cgi-bin/poi/addpoi?access_token=TOKEN
*/
func Addpoi(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAddpoi, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/cgi-bin/poi/getpoi



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html

POST http://api.weixin.qq.com/cgi-bin/poi/getpoi?access_token=TOKEN
*/
func Getpoi(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetpoi, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/cgi-bin/poi/getpoilist



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/cgi-bin/poi/getpoilist?access_token=TOKEN
*/
func Getpoilist(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetpoilist, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/cgi-bin/poi/updatepoi



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/cgi-bin/poi/updatepoi?access_token=TOKEN
*/
func Updatepoi(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdatepoi, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/cgi-bin/poi/delpoi



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/cgi-bin/poi/delpoi?access_token=TOKEN
*/
func Delpoi(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelpoi, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/cgi-bin/poi/getwxcategory



See: https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html

POST http://api.weixin.qq.com/cgi-bin/poi/getwxcategory?access_token=TOKEN
*/
func Getwxcategory(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetwxcategory, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
