package card

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiCreate                  = "/card/create"
	apiSet                     = "/card/paycell/set"
	apiSet                     = "/card/selfconsumecell/set"
	apiCreate                  = "/card/qrcode/create"
	apiCreate                  = "/card/landingpage/create"
	apiGethtml                 = "/card/mpnews/gethtml"
	apiSet                     = "/card/testwhitelist/set"
	apiGet                     = "/card/code/get"
	apiConsume                 = "/card/code/consume"
	apiDecrypt                 = "/card/code/decrypt"
	apiGetcardlist             = "/card/user/getcardlist"
	apiGet                     = "/card/get"
	apiBatchget                = "/card/batchget"
	apiUpdate                  = "/card/update"
	apiModifystock             = "/card/modifystock"
	apiUpdate                  = "/card/code/update"
	apiDelete                  = "/card/delete"
	apiUnavailable             = "/card/code/unavailable"
	apiGetcardbizuininfo       = "/datacube/getcardbizuininfo"
	apiGetcardcardinfo         = "/datacube/getcardcardinfo"
	apiGetcardmembercardinfo   = "/datacube/getcardmembercardinfo"
	apiGetcardmembercarddetail = "/datacube/getcardmembercarddetail"
	apiGeturl                  = "/card/membercard/activate/geturl"
	apiGet                     = "/card/membercard/activatetempinfo/get"
	apiActivate                = "/card/membercard/activate"
	apiAdd                     = "/card/giftcard/page/add"
	apiGet                     = "/card/giftcard/page/get"
	apiUpdate                  = "/card/giftcard/page/update"
	apiBatchget                = "/card/giftcard/page/batchget"
	apiSet                     = "/card/giftcard/maintain/set"
	apiAdd                     = "/card/giftcard/pay/whitelist/add"
	apiBind                    = "/card/giftcard/pay/submch/bind"
	apiSet                     = "/card/giftcard/wxa/set"
	apiSet                     = "/card/giftcard/wxa/set"
	apiGet                     = "/card/giftcard/order/get"
	apiBatchget                = "/card/giftcard/order/batchget"
	apiUpdateuser              = "/card/generalcard/updateuser"
	apiRefund                  = "/card/giftcard/order/refund"
	apiSetbizattr              = "/card/invoice/setbizattr"
	apiGetauthdata             = "/card/invoice/getauthdata"
	apiGetauthdata             = "/card/invoice/getauthdata"
	apiSet                     = "/card/membercard/activateuserform/set"
	apiGet                     = "/card/membercard/userinfo/get"
	apiUpdateuser              = "/card/membercard/updateuser"
	apiAdd                     = "/card/paygiftcard/add"
	apiDelete                  = "/card/paygiftcard/delete"
	apiGetbyid                 = "/card/paygiftcard/getbyid"
	apiBatchget                = "/card/paygiftcard/batchget"
	apiUpdateuser              = "/card/meetingticket/updateuser"
	apiUpdateuser              = "/card/movieticket/updateuser"
	apiCheckin                 = "/card/boardingpass/checkin"
	apiSubmit                  = "/card/submerchant/submit"
	apiGetapplyprotocol        = "/card/getapplyprotocol"
	apiUpdate                  = "/card/submerchant/update"
	apiGet                     = "/card/submerchant/get"
	apiBatchget                = "/card/submerchant/batchget"
)

/*
/card/create



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Create_a_Coupon_Voucher_or_Card.html

POST https://api.weixin.qq.com/card/create?access_token=ACCESS_TOKEN
*/
func Create(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiCreate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/paycell/set



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Create_a_Coupon_Voucher_or_Card.html

POST https://api.weixin.qq.com/card/paycell/set?access_token=TOKEN
*/
func Set(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/selfconsumecell/set



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Create_a_Coupon_Voucher_or_Card.html

POST https://api.weixin.qq.com/card/selfconsumecell/set?access_token=TOKEN
*/
func Set(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/qrcode/create



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Distributing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/qrcode/create?access_token=TOKEN
*/
func Create(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiCreate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/landingpage/create



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Distributing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/landingpage/create?access_token=$TOKEN
*/
func Create(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiCreate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/mpnews/gethtml



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Distributing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/mpnews/gethtml?access_token=TOKEN
*/
func Gethtml(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGethtml, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/testwhitelist/set



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Distributing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/testwhitelist/set?access_token=TOKEN
*/
func Set(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/code/get



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Redeeming_a_coupon_voucher_or_card.html

POST https://api.weixin.qq.com/card/code/get?access_token=TOKEN
*/
func Get(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/code/consume



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Redeeming_a_coupon_voucher_or_card.html

POST https://api.weixin.qq.com/card/code/consume?access_token=TOKEN
*/
func Consume(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiConsume, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/code/decrypt



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Redeeming_a_coupon_voucher_or_card.html

POST https://api.weixin.qq.com/card/code/decrypt?access_token=TOKEN
*/
func Decrypt(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDecrypt, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/user/getcardlist



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/user/getcardlist?access_token=TOKEN
*/
func Getcardlist(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetcardlist, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/get



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/get?access_token=TOKEN
*/
func Get(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/batchget



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/batchget?access_token=TOKEN
*/
func Batchget(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiBatchget, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/update



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/update?access_token=TOKEN
*/
func Update(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/modifystock



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/modifystock?access_token=TOKEN
*/
func Modifystock(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiModifystock, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/code/update



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/code/update?access_token=TOKEN
*/
func Update(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/delete



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/delete?access_token=TOKEN
*/
func Delete(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelete, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/code/unavailable



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/card/code/unavailable?access_token=TOKEN
*/
func Unavailable(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUnavailable, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/datacube/getcardbizuininfo



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/datacube/getcardbizuininfo?access_token=ACCESS_TOKEN
*/
func Getcardbizuininfo(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetcardbizuininfo, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/datacube/getcardcardinfo



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/datacube/getcardcardinfo?access_token=ACCESS_TOKEN
*/
func Getcardcardinfo(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetcardcardinfo, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/datacube/getcardmembercardinfo



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/datacube/getcardmembercardinfo?access_token=ACCESS_TOKEN
*/
func Getcardmembercardinfo(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetcardmembercardinfo, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/datacube/getcardmembercarddetail



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html

POST https://api.weixin.qq.com/datacube/getcardmembercarddetail?access_token=ACCESS_TOKEN
*/
func Getcardmembercarddetail(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetcardmembercarddetail, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/membercard/activate/geturl



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Coupons-Mini_Program_Start_Up.html

POST https://api.weixin.qq.com/card/membercard/activate/geturl?access_token=
*/
func Geturl(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGeturl, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/membercard/activatetempinfo/get



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Coupons-Mini_Program_Start_Up.html

POST https://api.weixin.qq.com/card/membercard/activatetempinfo/get?access_token=TOKEN
*/
func Get(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/membercard/activate



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Coupons-Mini_Program_Start_Up.html

POST https://api.weixin.qq.com/card/membercard/activate?access_token=TOKEN
*/
func Activate(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiActivate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/giftcard/page/add



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/page/add?access_token=ACCESS_TOKEN
*/
func Add(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAdd, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/giftcard/page/get



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/page/get?access_token=ACCESS_TOKEN
*/
func Get(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/giftcard/page/update



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/page/update?access_token=ACCESS_TOKEN
*/
func Update(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/giftcard/page/batchget



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/page/batchget?access_token=ACCESS_TOKEN
*/
func Batchget(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiBatchget, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/giftcard/maintain/set



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/maintain/set?access_token=ACCESS_TOKEN
*/
func Set(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/giftcard/pay/whitelist/add



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/pay/whitelist/add?access_token=TOKEN
*/
func Add(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAdd, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/giftcard/pay/submch/bind



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/pay/submch/bind?access_token=TOKEN
*/
func Bind(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiBind, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/giftcard/wxa/set



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/wxa/set
*/
func Set(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/giftcard/wxa/set<span></span></a><em>



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/wxa/set
*/
func Set(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/giftcard/order/get



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/order/get?access_token=ACCESS_TOKEN
*/
func Get(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/giftcard/order/batchget



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/order/batchget?access_token=ACCESS_TOKEN
*/
func Batchget(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiBatchget, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/generalcard/updateuser



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/generalcard/updateuser?access_token=TOKEN
*/
func Updateuser(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdateuser, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/giftcard/order/refund



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/giftcard/order/refund?access_token=ACCESS_TOKEN
*/
func Refund(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiRefund, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/setbizattr



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/invoice/setbizattr?action=set_pay_mch&amp;access_token={access_token}
*/
func Setbizattr(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSetbizattr, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/getauthdata



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/invoice/getauthdata
*/
func Getauthdata(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetauthdata, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/invoice/getauthdata<span></span></a>



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html

POST https://api.weixin.qq.com/card/invoice/getauthdata
*/
func Getauthdata(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetauthdata, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/membercard/activateuserform/set



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Create_a_membership_card.html

POST https://api.weixin.qq.com/card/membercard/activateuserform/set?access_token=TOKEN
*/
func Set(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/membercard/userinfo/get



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Create_a_membership_card.html

POST https://api.weixin.qq.com/card/membercard/userinfo/get?access_token=TOKEN
*/
func Get(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/membercard/updateuser



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Create_a_membership_card.html

POST https://api.weixin.qq.com/card/membercard/updateuser?access_token=TOKEN
*/
func Updateuser(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdateuser, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/paygiftcard/add



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Manage_Member_Card.html

POST https://api.weixin.qq.com/card/paygiftcard/add?access_token=TOKEN
*/
func Add(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAdd, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/paygiftcard/delete



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Manage_Member_Card.html

POST https://api.weixin.qq.com/card/paygiftcard/delete?access_token=TOKEN
*/
func Delete(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelete, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/paygiftcard/getbyid



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Manage_Member_Card.html

POST https://api.weixin.qq.com/card/paygiftcard/getbyid?access_token=TOKEN
*/
func Getbyid(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetbyid, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/paygiftcard/batchget



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Manage_Member_Card.html

POST https://api.weixin.qq.com/card/paygiftcard/batchget?access_token=TOKEN
*/
func Batchget(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiBatchget, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/meetingticket/updateuser



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Special_ticket.html

POST https://api.weixin.qq.com/card/meetingticket/updateuser?access_token=TOKEN
*/
func Updateuser(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdateuser, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/movieticket/updateuser



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Special_ticket.html

POST https://api.weixin.qq.com/card/movieticket/updateuser?access_token=TOKEN
*/
func Updateuser(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdateuser, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/boardingpass/checkin



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Special_ticket.html

POST https://api.weixin.qq.com/card/boardingpass/checkin?access_token=TOKEN
*/
func Checkin(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiCheckin, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/submerchant/submit



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html

POST https://api.weixin.qq.com/card/submerchant/submit?access_token=TOKEN
*/
func Submit(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiSubmit, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/getapplyprotocol



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html

POST https://api.weixin.qq.com/card/getapplyprotocol?access_token=TOKEN
*/
func Getapplyprotocol(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetapplyprotocol, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/submerchant/update



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html

POST https://api.weixin.qq.com/card/submerchant/update?access_token=TOKEN
*/
func Update(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/submerchant/get



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html

POST https://api.weixin.qq.com/card/submerchant/get?access_token=TOKEN
*/
func Get(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
/card/submerchant/batchget



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html

POST https://api.weixin.qq.com/card/submerchant/batchget?access_token=TOKEN
*/
func Batchget(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiBatchget, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
