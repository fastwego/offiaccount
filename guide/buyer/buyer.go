package buyer

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiAddGuideBuyerRelation        = "/cgi-bin/guide/addguidebuyerrelation"
	apiDelGuideBuyerRelation        = "/cgi-bin/guide/delguidebuyerrelation"
	apiRebindGuideAcctForBuyer      = "/cgi-bin/guide/rebindguideacctforbuyer"
	apiUpdateGuideBuyerRelation     = "/cgi-bin/guide/updateguidebuyerrelation"
	apiGetGuideBuyerRelationByBuyer = "/cgi-bin/guide/getguidebuyerrelationbybuyer"
	apiGetGuideBuyerRelation        = "/cgi-bin/guide/getguidebuyerrelation"
)

/*
为顾问分配客户



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.addGuideBuyerRelation.html

POST https://api.weixin.qq.com/cgi-bin/guide/addguidebuyerrelation?access_token=ACCESS_TOKEN
*/
func AddGuideBuyerRelation(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAddGuideBuyerRelation, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
为顾问移除客户



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.delGuideBuyerRelation.html

POST https://api.weixin.qq.com/cgi-bin/guide/delguidebuyerrelation?access_token=ACCESS_TOKEN
*/
func DelGuideBuyerRelation(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelGuideBuyerRelation, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
为客户更换顾问



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.rebindGuideAcctForBuyer.html

POST https://api.weixin.qq.com/cgi-bin/guide/rebindguideacctforbuyer?access_token=ACCESS_TOKEN
*/
func RebindGuideAcctForBuyer(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiRebindGuideAcctForBuyer, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
修改客户昵称



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.updateGuideBuyerRelation.html

POST https://api.weixin.qq.com/cgi-bin/guide/updateguidebuyerrelation?access_token=ACCESS_TOKEN
*/
func UpdateGuideBuyerRelation(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdateGuideBuyerRelation, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询客户所属顾问



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.getGuideBuyerRelationByBuyer.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidebuyerrelationbybuyer?access_token=ACCESS_TOKEN
*/
func GetGuideBuyerRelationByBuyer(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetGuideBuyerRelationByBuyer, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询指定顾问和客户的关系



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.getGuideBuyerRelation.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidebuyerrelation?access_token=ACCESS_TOKEN
*/
func GetGuideBuyerRelation(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetGuideBuyerRelation, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
