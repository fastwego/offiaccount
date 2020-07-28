package tag

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiCreate         = "/cgi-bin/tags/create"
	apiGet            = "/cgi-bin/tags/get"
	apiUpdate         = "/cgi-bin/tags/update"
	apiDelete         = "/cgi-bin/tags/delete"
	apiGetUsers       = "/cgi-bin/user/tag/get"
	apiBatchTagging   = "/cgi-bin/tags/members/batchtagging"
	apiBatchUntagging = "/cgi-bin/tags/members/batchuntagging"
	apiGetIdList      = "/cgi-bin/tags/getidlist"
)

/*
创建标签

一个公众号，最多可以创建100个标签

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html

POST https://api.weixin.qq.com/cgi-bin/tags/create?access_token=ACCESS_TOKEN
*/
func Create(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiCreate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取公众号已创建的标签



See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html

GET https://api.weixin.qq.com/cgi-bin/tags/get?access_token=ACCESS_TOKEN
*/
func Get() (resp []byte, err error) {
	return offiaccount.HTTPGet(apiGet)
}

/*
编辑标签



See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html

POST https://api.weixin.qq.com/cgi-bin/tags/update?access_token=ACCESS_TOKEN
*/
func Update(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
删除标签

请注意，当某个标签下的粉丝超过10w时，后台不可直接删除标签。此时，开发者可以对该标签下的openid列表，先进行取消标签的操作，直到粉丝数不超过10w后，才可直接删除该标签

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html

POST https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=ACCESS_TOKEN
*/
func Delete(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelete, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取标签下粉丝列表



See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html

GET https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=ACCESS_TOKEN
*/
func GetUsers() (resp []byte, err error) {
	return offiaccount.HTTPGet(apiGetUsers)
}

/*
批量为用户打标签

标签功能目前支持公众号为用户打上最多20个标签

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html

POST https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token=ACCESS_TOKEN
*/
func BatchTagging(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiBatchTagging, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
批量为用户取消标签

标签功能目前支持公众号为用户打上最多20个标签

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html

POST https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token=ACCESS_TOKEN
*/
func BatchUntagging(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiBatchUntagging, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取用户身上的标签列表

标签功能目前支持公众号为用户打上最多20个标签

See: https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html

POST https://api.weixin.qq.com/cgi-bin/tags/getidlist?access_token=ACCESS_TOKEN
*/
func GetIdList(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetIdList, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
