// Package ticket 特殊票券
package ticket

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiMeetingTicketUpdateUser = "/card/meetingticket/updateuser"
	apiMovieTicketUpdateUser   = "/card/movieticket/updateuser"
	apiBoardingPassCheckin     = "/card/boardingpass/checkin"
)

/*
更新会议门票

持调用“更新会议门票”接口update 入场时间、区域、座位等信息

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Special_ticket.html

POST https://api.weixin.qq.com/card/meetingticket/updateuser?access_token=TOKEN
*/
func MeetingTicketUpdateUser(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMeetingTicketUpdateUser, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
更新电影票

领取电影票后通过调用“更新电影票”接口update电影信息及用户选座信息

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Special_ticket.html

POST https://api.weixin.qq.com/card/movieticket/updateuser?access_token=TOKEN
*/
func MovieTicketUpdateUser(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMovieTicketUpdateUser, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}

/*
更新飞机票信息



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Special_ticket.html

POST https://api.weixin.qq.com/card/boardingpass/checkin?access_token=TOKEN
*/
func BoardingPassCheckin(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBoardingPassCheckin, bytes.NewBuffer(payload), "application/json;charset=utf-8")
}
