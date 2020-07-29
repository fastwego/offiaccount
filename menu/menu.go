package menu

import (
	"bytes"

	"github.com/fastwego/offiaccount"
)

const (
	apiCreate                 = "/cgi-bin/menu/create"
	apiGetCurrentSelfmenuInfo = "/cgi-bin/get_current_selfmenu_info"
	apiDelete                 = "/cgi-bin/menu/delete"
	apiAddConditional         = "/cgi-bin/menu/addconditional"
	apiDelConditional         = "/cgi-bin/menu/delconditional"
	apiTryMatch               = "/cgi-bin/menu/trymatch"
	apiGet                    = "/cgi-bin/menu/get"
)

/*
创建

创建菜单

See: https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Creating_Custom-Defined_Menu.html

POST https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN
*/
func Create(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiCreate, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
查询

本接口将会提供公众号当前使用的自定义菜单的配置，如果公众号是通过API调用设置的菜单，则返回菜单的开发配置，而如果公众号是在公众平台官网通过网站功能发布菜单，则本接口返回运营者设置的菜单配置

See: https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Querying_Custom_Menus.html

GET https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=ACCESS_TOKEN
*/
func GetCurrentSelfmenuInfo() (resp []byte, err error) {
	return offiaccount.HTTPGet(apiGetCurrentSelfmenuInfo)
}

/*
删除

使用接口创建自定义菜单后，开发者还可使用接口删除当前使用的自定义菜单。另请注意，在个性化菜单时，调用此接口会删除默认菜单及全部个性化菜单

See: https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Deleting_Custom-Defined_Menu.html

GET https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN
*/
func Delete() (resp []byte, err error) {
	return offiaccount.HTTPGet(apiDelete)
}

/*
创建个性化菜单



See: https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Personalized_menu_interface.html

POST https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token=ACCESS_TOKEN
*/
func AddConditional(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAddConditional, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
删除个性化菜单



See: https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Personalized_menu_interface.html

POST https://api.weixin.qq.com/cgi-bin/menu/delconditional?access_token=ACCESS_TOKEN
*/
func DelConditional(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelConditional, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
测试个性化菜单匹配结果



See: https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Personalized_menu_interface.html

POST https://api.weixin.qq.com/cgi-bin/menu/trymatch?access_token=ACCESS_TOKEN
*/
func TryMatch(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiTryMatch, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取自定义菜单配置

使用接口创建自定义菜单后，开发者还可使用接口查询自定义菜单的结构。另外请注意，在设置了个性化菜单后，使用本自定义菜单查询接口可以获取默认菜单和全部个性化菜单信息

See: https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Getting_Started_Guide.html

POST https://api.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN
*/
func Get(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGet, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
