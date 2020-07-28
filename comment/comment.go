package comment

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiOpen        = "/cgi-bin/comment/open"
	apiClose       = "/cgi-bin/comment/close"
	apiList        = "/cgi-bin/comment/list"
	apiMarkelect   = "/cgi-bin/comment/markelect"
	apiUnmarkelect = "/cgi-bin/comment/unmarkelect"
	apiDelete      = "/cgi-bin/comment/delete"
	apiReplyAdd    = "/cgi-bin/comment/reply/add"
	apiReplyDelete = "/cgi-bin/comment/reply/delete"
)

/*
打开已群发文章评论


See: https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html

POST https://api.weixin.qq.com/cgi-bin/comment/open?access_token=ACCESS_TOKEN
*/
func Open(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiOpen, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
关闭已群发文章评论



See: https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html

POST https://api.weixin.qq.com/cgi-bin/comment/close?access_token=ACCESS_TOKEN
*/
func Close(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiClose, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查看指定文章的评论数据



See: https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html

POST https://api.weixin.qq.com/cgi-bin/comment/list?access_token=ACCESS_TOKEN
*/
func List(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiList, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
 将评论标记精选



See: https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html

POST https://api.weixin.qq.com/cgi-bin/comment/markelect?access_token=ACCESS_TOKEN
*/
func Markelect(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiMarkelect, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
 将评论取消精选



See: https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html

POST https://api.weixin.qq.com/cgi-bin/comment/unmarkelect?access_token=ACCESS_TOKEN
*/
func Unmarkelect(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUnmarkelect, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
 删除评论



See: https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html

POST https://api.weixin.qq.com/cgi-bin/comment/delete?access_token=ACCESS_TOKEN
*/
func Delete(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelete, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
 回复评论



See: https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html

POST https://api.weixin.qq.com/cgi-bin/comment/reply/add?access_token=ACCESS_TOKEN
*/
func ReplyAdd(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiReplyAdd, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
 删除回复



See: https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html

POST https://api.weixin.qq.com/cgi-bin/comment/reply/delete?access_token=ACCESS_TOKEN
*/
func ReplyDelete(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiReplyDelete, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
