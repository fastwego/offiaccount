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
	return ctx.Client.HTTPPost(apiMeetingTicketUpdateUser, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新电影票

领取电影票后通过调用“更新电影票”接口update电影信息及用户选座信息

See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Special_ticket.html

POST https://api.weixin.qq.com/card/movieticket/updateuser?access_token=TOKEN
*/
func MovieTicketUpdateUser(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiMovieTicketUpdateUser, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
更新飞机票信息



See: https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Special_ticket.html

POST https://api.weixin.qq.com/card/boardingpass/checkin?access_token=TOKEN
*/
func BoardingPassCheckin(ctx *offiaccount.OffiAccount, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiBoardingPassCheckin, bytes.NewReader(payload), "application/json;charset=utf-8")
}
