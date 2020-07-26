package menu

import (
	"bytes"
	"encoding/json"
	"net/url"

	"github.com/fastwego/offiaccount"
)

const (
	//click：点击推事件用户点击click类型按钮后，微信服务器会通过消息接口推送消息类型为event的结构给开发者（参考消息接口指南），并且带上按钮中开发者填写的key值，开发者可以通过自定义的key值与用户进行交互；
	ButtonTypeClick = "click"

	//view：跳转URL用户点击view类型按钮后，微信客户端将会打开开发者在按钮中填写的网页URL，可与网页授权获取用户基本信息接口结合，获得用户基本信息。
	ButtonTypeView = "view"

	//scancode_push：扫码推事件用户点击按钮后，微信客户端将调起扫一扫工具，完成扫码操作后显示扫描结果（如果是URL，将进入URL），且会将扫码的结果传给开发者，开发者可以下发消息。
	ButtonTypeScancodePush = "scancode_push"

	//scancode_waitmsg：扫码推事件且弹出“消息接收中”提示框用户点击按钮后，微信客户端将调起扫一扫工具，完成扫码操作后，将扫码的结果传给开发者，同时收起扫一扫工具，然后弹出“消息接收中”提示框，随后可能会收到开发者下发的消息。
	ButtonTypeScancodeWaitmsg = "scancode_waitmsg"

	//pic_sysphoto：弹出系统拍照发图用户点击按钮后，微信客户端将调起系统相机，完成拍照操作后，会将拍摄的相片发送给开发者，并推送事件给开发者，同时收起系统相机，随后可能会收到开发者下发的消息。
	ButtonTypeSysPhoto = "pic_sysphoto"

	//pic_photo_or_album：弹出拍照或者相册发图用户点击按钮后，微信客户端将弹出选择器供用户选择“拍照”或者“从手机相册选择”。用户选择后即走其他两种流程。
	ButtonTypePicPhotoOrAlbum = "pic_photo_or_album"

	//pic_weixin：弹出微信相册发图器用户点击按钮后，微信客户端将调起微信相册，完成选择操作后，将选择的相片发送给开发者的服务器，并推送事件给开发者，同时收起相册，随后可能会收到开发者下发的消息。
	ButtonTypePicWeixin = "pic_weixin"

	//location_select：弹出地理位置选择器用户点击按钮后，微信客户端将调起地理位置选择工具，完成选择操作后，将选择的地理位置发送给开发者的服务器，同时收起位置选择工具，随后可能会收到开发者下发的消息。
	ButtonTypeLocationSelect = "location_select"

	//media_id：下发消息（除文本消息）用户点击media_id类型按钮后，微信服务器会将开发者填写的永久素材id对应的素材下发给用户，永久素材类型可以是图片、音频、视频、图文消息。请注意：永久素材id必须是在“素材管理/新增永久素材”接口上传后获得的合法id。
	ButtonTypeMediaId = "media_id"

	//view_limited：跳转图文消息URL用户点击view_limited类型按钮后，微信客户端将打开开发者在按钮中填写的永久素材id对应的图文消息URL，永久素材类型只支持图文消息。请注意：永久素材id必须是在“素材管理/新增永久素材”接口上传后获得的合法id。​
	ButtonTypeViewLimited = "view_limited"

	// miniprogram 小程序
	ButtonTypeMiniProgram = "miniprogram"
)

type Button struct {
	Type       string   `json:"type,omitempty"`       // 菜单的响应动作类型，view表示网页类型，click表示点击类型，miniprogram表示小程序类型
	Name       string   `json:"name,omitempty"`       // 菜单标题，不超过16个字节，子菜单不超过40个字节
	Key        string   `json:"key,omitempty"`        // 菜单KEY值，用于消息接口推送，不超过128字节
	URL        string   `json:"url,omitempty"`        // 网页链接，用户点击菜单可打开链接，不超过1024字节。当type为miniprogram时，不支持小程序的老版本客户端将打开本url
	MediaID    string   `json:"media_id,omitempty"`   // 调用新增永久素材接口返回的合法media_id
	AppID      string   `json:"appid,omitempty"`      // 小程序的appid
	PagePath   string   `json:"pagepath,omitempty"`   // 小程序的页面路径
	SubButtons []Button `json:"sub_button,omitempty"` // 二级菜单数组，个数应为1~5个
}

/*
创建 接口

自定义菜单最多包括3个一级菜单，每个一级菜单最多包含5个二级菜单。

一级菜单最多4个汉字，二级菜单最多7个汉字，多出来的部分将会以“...”代替。

创建自定义菜单后，菜单的刷新策略是，在用户进入公众号会话页或公众号profile页时，如果发现上一次拉取菜单的请求在5分钟以前，就会拉取一下菜单，如果菜单有更新，就会刷新客户端的菜单。测试时可以尝试取消关注公众账号后再次关注，则可以看到创建后的效果。​

See: https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Creating_Custom-Defined_Menu.html：

POST https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN

*/
func Create(buttons []Button) (resp []byte, err error) {
	data, _ := json.Marshal(struct {
		Buttons []Button `json:"button"`
	}{Buttons: buttons})

	return offiaccount.HTTPPost("/cgi-bin/menu/create", url.Values{}, bytes.NewBuffer(data), offiaccount.ContentTypeApplicationJson)
}

/*
获取 接口

See: https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Getting_Custom_Menu_Configurations.html

*/
func Get() (resp []byte, err error) {
	return offiaccount.HTTPGet("/cgi-bin/menu/get", url.Values{})
}

/**
自定义菜单获取接口

See: https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Querying_Custom_Menus.html

GET https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=ACCESS_TOKEN

*/
func GetCurrentSelfMenuInfo() (resp []byte, err error) {
	return offiaccount.HTTPGet("/cgi-bin/get_current_selfmenu_info", url.Values{})
}

/**
自定义菜单删除接口

See: https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Deleting_Custom-Defined_Menu.html：

GET https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN

*/
func Delete() (resp []byte, err error) {
	return offiaccount.HTTPGet("/cgi-bin/menu/delete", url.Values{})
}

type MatchRule struct {
	GroupID            string `json:"group_id,omitempty"`
	Sex                string `json:"sex,omitempty"`
	Country            string `json:"country,omitempty"`
	Province           string `json:"province,omitempty"`
	City               string `json:"city,omitempty"`
	ClientPlatformType string `json:"client_platform_type,omitempty"`
	Language           string `json:"language,omitempty"`
}

/*
创建个性化菜单 接口

See: https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Personalized_menu_interface.html

POST https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token=ACCESS_TOKEN

*/
func AddConditional(buttons []Button, matchRule MatchRule) (resp []byte, err error) {
	data, _ := json.Marshal(struct {
		Buttons   []Button  `json:"button"`
		MatchRule MatchRule `json:"matchrule"`
	}{
		Buttons:   buttons,
		MatchRule: matchRule,
	})
	return offiaccount.HTTPPost("/cgi-bin/menu/addconditional", url.Values{}, bytes.NewReader(data), offiaccount.ContentTypeApplicationJson)
}

/*
删除个性化菜单 接口

See: https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Personalized_menu_interface.html#1

POST https://api.weixin.qq.com/cgi-bin/menu/delconditional?access_token=ACCESS_TOKEN

*/
func DelConditional(menuId string) (resp []byte, err error) {
	data, _ := json.Marshal(struct {
		MenuId string `json:"menuid"`
	}{MenuId: menuId})
	return offiaccount.HTTPPost("/cgi-bin/menu/delconditional", url.Values{}, bytes.NewBuffer(data), offiaccount.ContentTypeApplicationJson)
}

/*
测试个性化菜单匹配结果 接口

See: https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Personalized_menu_interface.html#1

POST https://api.weixin.qq.com/cgi-bin/menu/trymatch?access_token=ACCESS_TOKEN

*/
func TryMatch(userId string) (resp []byte, err error) {
	data, _ := json.Marshal(struct {
		UserId string `json:"user_id"`
	}{UserId: userId})
	return offiaccount.HTTPPost("/cgi-bin/menu/trymatch", url.Values{}, bytes.NewBuffer(data), offiaccount.ContentTypeApplicationJson)
}
