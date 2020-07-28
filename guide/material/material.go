package materials

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiSetGuideCardMaterial  = "/cgi-bin/guide/setguidecardmaterial"
	apiGetGuideCardMaterial  = "/cgi-bin/guide/getguidecardmaterial"
	apiDelGuideCardMaterial  = "/cgi-bin/guide/delguidecardmaterial"
	apiSetGuideImageMaterial = "/cgi-bin/guide/setguideimagematerial"
	apiGetGuideImageMaterial = "/cgi-bin/guide/getguideimagematerial"
	apiDelGuideImageMaterial = "/cgi-bin/guide/delguideimagematerial"
	apiSetGuideWordMaterial  = "/cgi-bin/guide/setguidewordmaterial"
	apiGetGuideWordMaterial  = "/cgi-bin/guide/getguidewordmaterial"
	apiDelGuideWordMaterial  = "/cgi-bin/guide/delguidewordmaterial"
)

/*
添加小程序卡片素材



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.setGuideCardMaterial.html

POST https://api.weixin.qq.com/cgi-bin/guide/setguidecardmaterial?access_token=ACCESS_TOKEN
*/
func SetGuideCardMaterial(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSetGuideCardMaterial, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询小程序卡片素材



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.getGuideCardMaterial.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidecardmaterial?access_token=ACCESS_TOKEN
*/
func GetGuideCardMaterial(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetGuideCardMaterial, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
删除小程序卡片素材



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.delGuideCardMaterial.html

POST https://api.weixin.qq.com/cgi-bin/guide/delguidecardmaterial?access_token=ACCESS_TOKEN
*/
func DelGuideCardMaterial(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelGuideCardMaterial, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
添加图片素材



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.setGuideImageMaterial.html

POST https://api.weixin.qq.com/cgi-bin/guide/setguideimagematerial?access_token=ACCESS_TOKEN
*/
func SetGuideImageMaterial(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSetGuideImageMaterial, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询图片素材



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.getGuideImageMaterial.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguideimagematerial?access_token=ACCESS_TOKEN
*/
func GetGuideImageMaterial(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetGuideImageMaterial, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
删除图片素材



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.delGuideImageMaterial.html

POST https://api.weixin.qq.com/cgi-bin/guide/delguideimagematerial?access_token=ACCESS_TOKEN
*/
func DelGuideImageMaterial(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelGuideImageMaterial, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
添加文字素材



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.setGuideWordMaterial.html

POST https://api.weixin.qq.com/cgi-bin/guide/setguidewordmaterial?access_token=ACCESS_TOKEN
*/
func SetGuideWordMaterial(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSetGuideWordMaterial, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询文字素材



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.getGuideWordMaterial.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidewordmaterial?access_token=ACCESS_TOKEN
*/
func GetGuideWordMaterial(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetGuideWordMaterial, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
删除文字素材



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.delGuideWordMaterial.html

POST https://api.weixin.qq.com/cgi-bin/guide/delguidewordmaterial?access_token=ACCESS_TOKEN
*/
func DelGuideWordMaterial(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelGuideWordMaterial, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
