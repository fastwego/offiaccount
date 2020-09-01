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
	EventTypeMenuClick              = "CLICK"              // 点击菜单拉取消息时的事件推送
	EventTypeMenuView               = "VIEW"               // 点击菜单跳转链接时的事件推送
	EventTypeMenuScanCodePush       = "scancode_push"      // 扫码推事件的事件推送
	EventTypeMenuScanCodeWaitMsg    = "scancode_waitmsg"   // 扫码推事件且弹出“消息接收中”提示框的事件推送
	EventTypeMenuPicSysPhoto        = "pic_sysphoto"       // 弹出系统拍照发图的事件推送
	EventTypeMenuPicSysPhotoOrAlbum = "pic_photo_or_album" // 弹出拍照或者相册发图的事件推送
	EventTypeMenuPicWeixin          = "pic_weixin"         // 弹出微信相册发图器的事件推送
	EventTypeMenuLocationSelect     = "location_select"    // 弹出地理位置选择器的事件推送
	EventTypeMenuViewMiniprogram    = "view_miniprogram"   // 点击菜单跳转小程序的事件推送
)

/*
<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[FromUser]]></FromUserName>
  <CreateTime>123456789</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[CLICK]]></Event>
  <EventKey><![CDATA[EVENTKEY]]></EventKey>
</xml>
*/
type EventMenuClick struct {
	Event
	EventKey string
}

/*
<xml>
	<ToUserName><![CDATA[toUser]]></ToUserName>
	<FromUserName><![CDATA[FromUser]]></FromUserName>
	<CreateTime>123456789</CreateTime>
	<MsgType><![CDATA[event]]></MsgType>
	<Event><![CDATA[VIEW]]></Event>
	<EventKey><![CDATA[www.qq.com]]></EventKey>
	<MenuId>MENUID</MenuId>
</xml>
*/
type EventMenuView struct {
	Event
	EventKey string
	MenuId   string
}

/*
<xml><ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName>
<FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName>
<CreateTime>1408090502</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[scancode_push]]></Event>
<EventKey><![CDATA[6]]></EventKey>
<ScanCodeInfo><ScanType><![CDATA[qrcode]]></ScanType>
<ScanResult><![CDATA[1]]></ScanResult>
</ScanCodeInfo>
</xml>
*/
type EventMenuScanCodePush struct {
	Event
	EventKey     string
	ScanCodeInfo struct {
		ScanType   string
		ScanResult string
	}
}

/*
<xml><ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName>
<FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName>
<CreateTime>1408090606</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[scancode_waitmsg]]></Event>
<EventKey><![CDATA[6]]></EventKey>
<ScanCodeInfo><ScanType><![CDATA[qrcode]]></ScanType>
<ScanResult><![CDATA[2]]></ScanResult>
</ScanCodeInfo>
</xml>
*/
type EventMenuScanCodeWaitMsg struct {
	Event
	EventKey     string
	ScanCodeInfo struct {
		ScanType   string
		ScanResult string
	}
}

type Item struct {
	PicMd5Sum string
}

/*
<xml><ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName>
<FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName>
<CreateTime>1408090651</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[pic_sysphoto]]></Event>
<EventKey><![CDATA[6]]></EventKey>
<SendPicsInfo><Count>1</Count>
<PicList><item><PicMd5Sum><![CDATA[1b5f7c23b5bf75682a53e7b6d163e185]]></PicMd5Sum>
</item>
</PicList>
</SendPicsInfo>
</xml>
*/
type EventMenuPicSysPhoto struct {
	Event
	EventKey     string
	SendPicsInfo struct {
		Count   string
		PicList struct {
			Item []Item `xml:"item"`
		}
	}
}

/*
<xml><ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName>
<FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName>
<CreateTime>1408090816</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[pic_photo_or_album]]></Event>
<EventKey><![CDATA[6]]></EventKey>
<SendPicsInfo><Count>1</Count>
<PicList><item><PicMd5Sum><![CDATA[5a75aaca956d97be686719218f275c6b]]></PicMd5Sum>
</item>
</PicList>
</SendPicsInfo>
</xml>
*/
type EventMenuPicSysPhotoOrAlbum struct {
	Event
	EventKey     string
	SendPicsInfo struct {
		Count   string
		PicList struct {
			Item []Item `xml:"item"`
		}
	}
}

/*
<xml><ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName>
<FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName>
<CreateTime>1408090816</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[pic_weixin]]></Event>
<EventKey><![CDATA[6]]></EventKey>
<SendPicsInfo><Count>1</Count>
<PicList><item><PicMd5Sum><![CDATA[5a75aaca956d97be686719218f275c6b]]></PicMd5Sum>
</item>
</PicList>
</SendPicsInfo>
</xml>
*/
type EventMenuPicWeixin struct {
	Event
	EventKey     string
	SendPicsInfo struct {
		Count   string
		PicList struct {
			Item []Item `xml:"item"`
		}
	}
}

/*
<xml><ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName>
<FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName>
<CreateTime>1408091189</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[location_select]]></Event>
<EventKey><![CDATA[6]]></EventKey>
<SendLocationInfo><Location_X><![CDATA[23]]></Location_X>
<Location_Y><![CDATA[113]]></Location_Y>
<Scale><![CDATA[15]]></Scale>
<Label><![CDATA[ 广州市海珠区客村艺苑路 106号]]></Label>
<Poiname><![CDATA[]]></Poiname>
</SendLocationInfo>
</xml>
*/
type EventMenuLocationSelect struct {
	Event
	EventKey         string
	SendLocationInfo struct {
		Location_X string
		Location_Y string
		Scale      string
		Label      string
		Poiname    string
	}
}

/*
<xml>
<ToUserName><![CDATA[toUser]]></ToUserName>
<FromUserName><![CDATA[FromUser]]></FromUserName>
<CreateTime>123456789</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[view_miniprogram]]></Event>
<EventKey><![CDATA[pages/index/index]]></EventKey>
<MenuId>MENUID</MenuId>
</xml>
*/
type EventMenuViewMiniprogram struct {
	Event
	EventKey string
	MenuID   string
}
