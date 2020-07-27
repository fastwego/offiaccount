package menu

import (
	"bytes"
	"encoding/json"

	"github.com/fastwego/offiaccount"
)

const (
	apiCreate                 = "/cgi-bin/menu/create"
	apiGet                    = "/cgi-bin/menu/get"
	apiGetCurrentSelfMenuInfo = "/cgi-bin/get_current_selfmenu_info"
	apiDelete                 = "/cgi-bin/menu/delete"
	apiAddConditional         = "/cgi-bin/menu/addconditional"
	apiDelConditional         = "/cgi-bin/menu/delconditional"
	apiTryMatch               = "/cgi-bin/menu/trymatch"
)

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

	return offiaccount.HTTPPost(apiCreate, bytes.NewBuffer(data), offiaccount.ContentTypeApplicationJson)
}

/*
获取 接口

See: https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Getting_Custom_Menu_Configurations.html

*/
func Get() (resp []byte, err error) {
	return offiaccount.HTTPGet(apiGet)
}

/**
自定义菜单获取接口

See: https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Querying_Custom_Menus.html

GET https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=ACCESS_TOKEN

*/
func GetCurrentSelfMenuInfo() (resp []byte, err error) {
	return offiaccount.HTTPGet(apiGetCurrentSelfMenuInfo)
}

/**
自定义菜单删除接口

See: https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Deleting_Custom-Defined_Menu.html：

GET https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN

*/
func Delete() (resp []byte, err error) {
	return offiaccount.HTTPGet(apiDelete)
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
	return offiaccount.HTTPPost(apiAddConditional, bytes.NewReader(data), offiaccount.ContentTypeApplicationJson)
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
	return offiaccount.HTTPPost(apiDelConditional, bytes.NewBuffer(data), offiaccount.ContentTypeApplicationJson)
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
	return offiaccount.HTTPPost(apiTryMatch, bytes.NewBuffer(data), offiaccount.ContentTypeApplicationJson)
}
