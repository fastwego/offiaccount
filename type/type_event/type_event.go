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

package type_event

import "github.com/fastwego/offiaccount/type/type_message"

const (
	EventTypeSubscribe                  = "subscribe"                    // 关注
	EventTypeUnsubscribe                = "unsubscribe"                  // 取关
	EventTypeScan                       = "SCAN"                         // 已关注用户 扫码
	EventTypeLocation                   = "LOCATION"                     // 上报位置
	EventTypeQualificationVerifySuccess = "qualification_verify_success" // 资质认证成功
	EventTypeQualificationVerifyFail    = "qualification_verify_fail"    // 资质认证失败
	EventTypeNamingVerifySuccess        = "naming_verify_success"        // 名称认证成功
	EventTypeNamingVerifyFail           = "naming_verify_fail"           // 名称认证失败
	EventTypeAnnualRenew                = "annual_renew"                 // 年审通知
	EventTypeVerifyExpired              = "verify_expired"               // 认证过期失效通知审通知
)

type Event type_message.MessageEvent

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[FromUser]]></FromUserName>
  <CreateTime>123456789</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[subscribe]]></Event>

  <EventKey><![CDATA[qrscene_123123]]></EventKey>
  <Ticket><![CDATA[TICKET]]></Ticket>
</xml>
*/
type EventSubscribe struct {
	Event
	// 扫码关注 附加参数
	EventKey string
	Ticket   string
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[FromUser]]></FromUserName>
  <CreateTime>123456789</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[SCAN]]></Event>
  <EventKey><![CDATA[SCENE_VALUE]]></EventKey>
  <Ticket><![CDATA[TICKET]]></Ticket>
</xml>
*/
type EventScan struct {
	EventKey string
	Ticket   string
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[FromUser]]></FromUserName>
  <CreateTime>123456789</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[unsubscribe]]></Event>
</xml>
*/
type EventUnsubscribe struct {
	Event
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>123456789</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[LOCATION]]></Event>
  <Latitude>23.137466</Latitude>
  <Longitude>113.352425</Longitude>
  <Precision>119.385040</Precision>
</xml>
*/
type EventLocation struct {
	Event
	Latitude  string
	Longitude string
	Precision string
}

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1442401156</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[qualification_verify_success]]></Event>
  <ExpiredTime>1442401156</ExpiredTime>
</xml>
*/
type EventQualificationVerifySuccess struct {
	Event
	ExpiredTime string
}

/*

<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1442401156</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[qualification_verify_fail]]></Event>
  <FailTime>1442401122</FailTime>
  <FailReason><![CDATA[by time]]></FailReason>
</xml>
*/
type EventQualificationVerifyFail struct {
	Event
	FailTime   string
	FailReason string
}

/*

<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1442401093</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[naming_verify_success]]></Event>
  <ExpiredTime>1442401093</ExpiredTime>
</xml>
*/
type EventNamingVerifySuccess struct {
	Event
	ExpiredTime string
}

/*
名称认证失败
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1442401061</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[naming_verify_fail]]></Event>
  <FailTime>1442401061</FailTime>
  <FailReason><![CDATA[by time]]></FailReason>
</xml>

*/
type EventNamingVerifyFail struct {
	Event
	FailTime   string
	FailReason string
}

/*
年审通知
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1442401004</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[annual_renew]]></Event>
  <ExpiredTime>1442401004</ExpiredTime>
</xml>

*/
type EventAnnualRenew struct {
	Event
	ExpiredTime string
}

/*
认证过期失效通知审通知
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1442400900</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[verify_expired]]></Event>
  <ExpiredTime>1442400900</ExpiredTime>
</xml>
*/
type EventVerifyExpired struct {
	Event
	ExpiredTime string
}
