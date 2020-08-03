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

const (
	EventTypeGuideQrcodeScan = "guide_qrcode_scan_event" // 微信用户扫顾问二维码后会触发事件推送
)

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[fromUser]]></FromUserName>
  <CreateTime>1546924844</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[guide_qrcode_scan_event]]></Event>
  <GuideScanEvent>
    <qrcode_guide_account>![CDATA[xxx]]</qrcode_guide_account>
    <qrcode_guide_openid>![CDATA[xxx]]</qrcode_guide_openid>
    <openid>![CDATA[xxx]]</openid>
    <action>11</action>
    <qrcode_info>![CDATA[xxx]]</qrcode_info>
  </GuideScanEvent>
</xml>
*/

type EventGuideQrcodeScan struct {
	Event
	GuideScanEvent struct {
		QrcodeGuideAccount string `xml:"qrcode_guide_account"`
		QrcodeGuideOpenid  string `xml:"qrcode_guide_openid"`
		Openid             string `xml:"openid"`
		Action             string `xml:"action"`
		QrcodeInfo         string `xml:"qrcode_info"`
	} `xml:"GuideScanEvent"`
}
