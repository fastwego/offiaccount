// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package buyer 服务号对话能力（原微信导购助手）/客户管理
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
func AddGuideBuyerRelation(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiAddGuideBuyerRelation, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
为顾问移除客户



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.delGuideBuyerRelation.html

POST https://api.weixin.qq.com/cgi-bin/guide/delguidebuyerrelation?access_token=ACCESS_TOKEN
*/
func DelGuideBuyerRelation(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelGuideBuyerRelation, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
为客户更换顾问



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.rebindGuideAcctForBuyer.html

POST https://api.weixin.qq.com/cgi-bin/guide/rebindguideacctforbuyer?access_token=ACCESS_TOKEN
*/
func RebindGuideAcctForBuyer(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiRebindGuideAcctForBuyer, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
修改客户昵称



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.updateGuideBuyerRelation.html

POST https://api.weixin.qq.com/cgi-bin/guide/updateguidebuyerrelation?access_token=ACCESS_TOKEN
*/
func UpdateGuideBuyerRelation(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdateGuideBuyerRelation, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询客户所属顾问



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.getGuideBuyerRelationByBuyer.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidebuyerrelationbybuyer?access_token=ACCESS_TOKEN
*/
func GetGuideBuyerRelationByBuyer(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGuideBuyerRelationByBuyer, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
查询指定顾问和客户的关系



See: https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.getGuideBuyerRelation.html

POST https://api.weixin.qq.com/cgi-bin/guide/getguidebuyerrelation?access_token=ACCESS_TOKEN
*/
func GetGuideBuyerRelation(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetGuideBuyerRelation, bytes.NewReader(payload), "application/json;charset=utf-8")
}
