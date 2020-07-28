package job

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiAddGuideMassendJob     = "/cgi-bin/guide/addguidemassendjob"
	apiGetGuideMassendJobList = "/cgi-bin/guide/getguidemassendjoblist"
	apiGetGuideMassendJob     = "/cgi-bin/guide/getguidemassendjob"
	apiUpdateGuideMassendJob  = "/cgi-bin/guide/updateguidemassendjob"
	apiCancelGuideMassendJob  = "/cgi-bin/guide/cancelguidemassendjob"
)

/*
添加群发任务



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/task-account/shopping-guide.addGuideMassendJob.html

POST https://api.weixin.qq.com/cgi-bin/guide/addguidemassendjob?access_token=ACCESS_TOKEN
*/
func AddGuideMassendJob(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAddGuideMassendJob, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取群发任务列表



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/task-account/shopping-guide.getGuideMassendJobList.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidemassendjoblist?access_token=ACCESS_TOKEN
*/
func GetGuideMassendJobList(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetGuideMassendJobList, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取指定群发任务信息



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/task-account/shopping-guide.getGuideMassendJob.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidemassendjob?access_token=ACCESS_TOKEN
*/
func GetGuideMassendJob(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetGuideMassendJob, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
修改群发任务



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/task-account/shopping-guide.updateGuideMassendJob.html

POST https://api.weixin.qq.com/cgi-bin/guide/updateguidemassendjob?access_token=ACCESS_TOKEN
*/
func UpdateGuideMassendJob(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdateGuideMassendJob, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
取消群发任务



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/task-account/shopping-guide.cancelGuideMassendJob.html

POST https://api.weixin.qq.com/cgi-bin/guide/cancelguidemassendjob?access_token=ACCESS_TOKEN
*/
func CancelGuideMassendJob(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiCancelGuideMassendJob, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
