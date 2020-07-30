// Package card 微信卡券
package card

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiCreate              = "/card/create"
	apiSetPayCell          = "/card/paycell/set"
	apiSetSelfConsumeCell  = "/card/selfconsumecell/set"
	apiCreateQRCode        = "/card/qrcode/create"
	apiCreateLandingPage   = "/card/landingpage/create"
	apiMpnewsGetHtml       = "/card/mpnews/gethtml"
	apiSetTestWhitelist    = "/card/testwhitelist/set"
	apiGetCode             = "/card/code/get"
	apiConsumeCode         = "/card/code/consume"
	apiDecryptCode         = "/card/code/decrypt"
	apiGetUserCardList     = "/card/user/getcardlist"
	apiGet                 = "/card/get"
	apiBatchGet            = "/card/batchget"
	apiUpdate              = "/card/update"
	apiModifyStock         = "/card/modifystock"
	apiUpdateCode          = "/card/code/update"
	apiDelete              = "/card/delete"
	apiUnavailableCoed     = "/card/code/unavailable"
	apiGetCardBizUinInfo   = "/datacube/getcardbizuininfo"
	apiGetCardInfo         = "/datacube/getcardcardinfo"
	apiGetMemberCardInfo   = "/datacube/getcardmembercardinfo"
	apiGetMemberCardDetail = "/datacube/getcardmembercarddetail"
)

/*
创建卡券

创建卡券接口是微信卡券的基础接口，用于创建一类新的卡券，获取card_id，创建成功并通过审核后，商家可以通过文档提供的其他接口将卡券下发给用户，每次成功领取，库存数量相应扣除

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Create_a_Coupon_Voucher_or_Card.html

POST https://api.weixin.qq.com/card/create?access_token=ACCESS_TOKEN
*/
func Create(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreate, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
设置买单

创建卡券之后，开发者可以通过设置微信买单接口设置该card_id支持微信买单功能。值得开发者注意的是，设置买单的card_id必须已经配置了门店，否则会报错

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Create_a_Coupon_Voucher_or_Card.html

POST https://api.weixin.qq.com/card/paycell/set?access_token=TOKEN
*/
func SetPayCell(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetPayCell, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
设置自助核销



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Create_a_Coupon_Voucher_or_Card.html

POST https://api.weixin.qq.com/card/selfconsumecell/set?access_token=TOKEN
*/
func SetSelfConsumeCell(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetSelfConsumeCell, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
创建二维码

开发者可调用该接口生成一张卡券二维码供用户扫码后添加卡券到卡包

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Distributing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/qrcode/create?access_token=TOKEN
*/
func CreateQRCode(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreateQRCode, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
创建货架

开发者需调用该接口创建货架链接，用于卡券投放。创建货架时需填写投放路径的场景字段

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Distributing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/landingpage/create?access_token=$TOKEN
*/
func CreateLandingPage(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCreateLandingPage, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
图文消息群发卡券

支持开发者调用该接口获取卡券嵌入图文消息的标准格式代码，将返回代码填入 新增临时素材中content字段，即可获取嵌入卡券的图文消息素材

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Distributing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/mpnews/gethtml?access_token=TOKEN
*/
func MpnewsGetHtml(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMpnewsGetHtml, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
设置测试白名单

由于卡券有审核要求，为方便公众号调试，可以设置一些测试帐号，这些帐号可领取未通过审核的卡券，体验整个流程

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Distributing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/testwhitelist/set?access_token=TOKEN
*/
func SetTestWhitelist(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiSetTestWhitelist, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
线下核销-查询 Code

我们强烈建议开发者在调用核销code接口之前调用查询code接口，并在核销之前对非法状态的code(如转赠中、已删除、已核销等)做出处理

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Redeeming_a_coupon_voucher_or_card.html

POST https://api.weixin.qq.com/card/code/get?access_token=TOKEN
*/
func GetCode(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetCode, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
核销 Code

消耗code接口是核销卡券的唯一接口,开发者可以调用当前接口将用户的优惠券进行核销，该过程不可逆

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Redeeming_a_coupon_voucher_or_card.html

POST https://api.weixin.qq.com/card/code/consume?access_token=TOKEN
*/
func ConsumeCode(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiConsumeCode, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
线上核销 - Code解码



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Redeeming_a_coupon_voucher_or_card.html

POST https://api.weixin.qq.com/card/code/decrypt?access_token=TOKEN
*/
func DecryptCode(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDecryptCode, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取用户已领取卡券

用于获取用户卡包里的，属于该appid下所有可用卡券，包括正常状态和异常状态

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/user/getcardlist?access_token=TOKEN
*/
func GetUserCardList(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetUserCardList, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
查看卡券详情

开发者可以调用该接口查询某个card_id的创建信息、审核状态以及库存数量

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/get?access_token=TOKEN
*/
func Get(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGet, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
批量查询卡券列表



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/batchget?access_token=TOKEN
*/
func BatchGet(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBatchGet, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
更改卡券信息

支持更新所有卡券类型的部分通用字段及特殊卡券（会员卡、飞机票、电影票、会议门票）中特定字段的信息

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/update?access_token=TOKEN
*/
func Update(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdate, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
修改库存

调用修改库存接口增减某张卡券的库存

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/modifystock?access_token=TOKEN
*/
func ModifyStock(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiModifyStock, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
更改Code

为确保转赠后的安全性，微信允许自定义Code的商户对已下发的code进行更改。 注：为避免用户疑惑，建议仅在发生转赠行为后（发生转赠后，微信会通过事件推送的方式告知商户被转赠的卡券Code）对用户的Code进行更改

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/code/update?access_token=TOKEN
*/
func UpdateCode(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUpdateCode, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
删除卡券

删除卡券接口允许商户删除任意一类卡券。删除卡券后，该卡券对应已生成的领取用二维码、添加到卡包JS API均会失效。 注意：如用户在商家删除卡券前已领取一张或多张该卡券依旧有效。即删除卡券不能删除已被用户领取，保存在微信客户端中的卡券

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/delete?access_token=TOKEN
*/
func Delete(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiDelete, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
设置卡券失效

为满足改票、退款等异常情况，可调用卡券失效接口将用户的卡券设置为失效状态

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/code/unavailable?access_token=TOKEN
*/
func UnavailableCoed(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiUnavailableCoed, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
拉取卡券概况数据

支持调用该接口拉取本商户的总体数据情况，包括时间区间内的各指标总量

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/datacube/getcardbizuininfo?access_token=ACCESS_TOKEN
*/
func GetCardBizUinInfo(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetCardBizUinInfo, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
获取免费券数据

支持开发者调用该接口拉取免费券（优惠券、团购券、折扣券、礼品券）在固定时间区间内的相关数据

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/datacube/getcardcardinfo?access_token=ACCESS_TOKEN
*/
func GetCardInfo(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetCardInfo, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
拉取会员卡概况数据

支持开发者调用该接口拉取公众平台创建的会员卡相关数据

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/datacube/getcardmembercardinfo?access_token=ACCESS_TOKEN
*/
func GetMemberCardInfo(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetMemberCardInfo, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
拉取单张会员卡数据

支持开发者调用该接口拉取API创建的会员卡数据情况

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/datacube/getcardmembercarddetail?access_token=ACCESS_TOKEN
*/
func GetMemberCardDetail(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetMemberCardDetail, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}
