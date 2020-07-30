// Package store 微信小店/功能接口
package store

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiUploadImg = "/merchant/common/upload_img"
)

/*
上传图片


See: https://developers.weixin.qq.com/doc/offiaccount/Instant_Stores/WeChat_Store_Interface.html

POST https://api.weixin.qq.com/merchant/common/upload_img?access_token=ACCESS_TOKEN&filename=test.png
*/
func UploadImg(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUploadImg, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}
