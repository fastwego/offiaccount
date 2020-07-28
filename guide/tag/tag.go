package tags

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiNewGuideTagOption       = "/cgi-bin/guide/newguidetagoption"
	apiDelguidetagoption       = "/cgi-bin/guide/delguidetagoption"
	apiAddGuideTagOption       = "/cgi-bin/guide/addguidetagoption"
	apiGetGuideTagOption       = "/cgi-bin/guide/getguidetagoption"
	apiAddGuideBuyerTag        = "/cgi-bin/guide/addguidebuyertag"
	apiGetGuideBuyerTag        = "/cgi-bin/guide/getguidebuyertag"
	apiQueryGuideBuyerByTag    = "/cgi-bin/guide/queryguidebuyerbytag"
	apiDelGuideBuyerTag        = "/cgi-bin/guide/delguidebuyertag"
	apiAddGuideBuyerDisplayTag = "/cgi-bin/guide/addguidebuyerdisplaytag"
	apiGetGuideBuyerDisplayTag = "/cgi-bin/guide/getguidebuyerdisplaytag"
)

/*
/cgi-bin/guide/newguidetagoption



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.newGuideTagOption.html

POST https://api.weixin.qq.com/cgi-bin/guide/newguidetagoption?access_token=ACCESS_TOKEN
*/
func NewGuideTagOption(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiNewGuideTagOption, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
删除指定标签类型



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.delguidetagoption.html

POST https://api.weixin.qq.com/cgi-bin/guide/delguidetagoption?access_token=ACCESS_TOKEN
*/
func Delguidetagoption(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelguidetagoption, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
为标签添加可选值



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.addGuideTagOption.html

POST https://api.weixin.qq.com/cgi-bin/guide/addguidetagoption?access_token=ACCESS_TOKEN
*/
func AddGuideTagOption(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAddGuideTagOption, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取标签和可选值



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.getGuideTagOption.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidetagoption?access_token=ACCESS_TOKEN
*/
func GetGuideTagOption(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetGuideTagOption, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
为客户设置标签



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.addGuideBuyerTag.html

POST https://api.weixin.qq.com/cgi-bin/guide/addguidebuyertag?access_token=ACCESS_TOKEN
*/
func AddGuideBuyerTag(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAddGuideBuyerTag, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询客户标签



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.getGuideBuyerTag.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidebuyertag?access_token=ACCESS_TOKEN
*/
func GetGuideBuyerTag(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetGuideBuyerTag, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
根据标签值筛选客户



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.queryGuideBuyerByTag.html

POST https://api.weixin.qq.com/cgi-bin/guide/queryguidebuyerbytag?access_token=ACCESS_TOKEN
*/
func QueryGuideBuyerByTag(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiQueryGuideBuyerByTag, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
删除客户标签



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.delGuideBuyerTag.html

POST https://api.weixin.qq.com/cgi-bin/guide/delguidebuyertag?access_token=ACCESS_TOKEN
*/
func DelGuideBuyerTag(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelGuideBuyerTag, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
设置自定义客户信息



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.addGuideBuyerDisplayTag.html

POST https://api.weixin.qq.com/cgi-bin/guide/addguidebuyerdisplaytag?access_token=ACCESS_TOKEN
*/
func AddGuideBuyerDisplayTag(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAddGuideBuyerDisplayTag, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取自定义客户信息



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.getGuideBuyerDisplayTag.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidebuyerdisplaytag?access_token=ACCESS_TOKEN
*/
func GetGuideBuyerDisplayTag(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetGuideBuyerDisplayTag, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
