package main

type Api struct {
	Name        string
	Description string
	Request     string
	See         string
	FuncName    string
}

type ApiGroup struct {
	Name    string
	Apis    []Api
	Package string
}

var apiConfig = []ApiGroup{
	{
		Name:    `自定义菜单`,
		Package: `menu`,
		Apis: []Api{
			{
				Name:        "创建",
				Description: "创建菜单",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Creating_Custom-Defined_Menu.html",
				FuncName:    "",
			},
			{
				Name:        "查询",
				Description: "本接口将会提供公众号当前使用的自定义菜单的配置，如果公众号是通过API调用设置的菜单，则返回菜单的开发配置，而如果公众号是在公众平台官网通过网站功能发布菜单，则本接口返回运营者设置的菜单配置",
				Request:     "GET https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Querying_Custom_Menus.html",
				FuncName:    "",
			},

			{
				Name:        "删除",
				Description: "使用接口创建自定义菜单后，开发者还可使用接口删除当前使用的自定义菜单。另请注意，在个性化菜单时，调用此接口会删除默认菜单及全部个性化菜单",
				Request:     "GET https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Deleting_Custom-Defined_Menu.html",
				FuncName:    "",
			},

			{
				Name:        "创建个性化菜单",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Personalized_menu_interface.html",
				FuncName:    "AddConditional",
			},
			{
				Name:        "删除个性化菜单",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/menu/delconditional?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Personalized_menu_interface.html",
				FuncName:    "DelConditional",
			},
			{
				Name:        "测试个性化菜单匹配结果",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/menu/trymatch?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Personalized_menu_interface.html",
				FuncName:    "TryMatch",
			},
			{
				Name:        "获取自定义菜单配置",
				Description: "使用接口创建自定义菜单后，开发者还可使用接口查询自定义菜单的结构。另外请注意，在设置了个性化菜单后，使用本自定义菜单查询接口可以获取默认菜单和全部个性化菜单信息",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Getting_Started/Getting_Started_Guide.html",
				FuncName:    "",
			},
		},
	},
	{
		Name:    `客服消息/功能`,
		Package: `customservice`,
		Apis: []Api{
			//客服消息
			{
				Name:        "添加客服帐号",
				Description: "开发者可以通过本接口为公众号添加客服账号，每个公众号最多添加100个客服账号",
				Request:     "POST https://api.weixin.qq.com/customservice/kfaccount/add?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html",
				FuncName:    "KfaccountAdd",
			},
			{
				Name:        "修改客服帐号",
				Description: "开发者可以通过本接口为公众号修改客服账号",
				Request:     "POST https://api.weixin.qq.com/customservice/kfaccount/update?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html",
				FuncName:    "KfaccountUpdate",
			},
			{
				Name:        "删除客服帐号",
				Description: "开发者可以通过该接口为公众号删除客服帐号",
				Request:     "POST https://api.weixin.qq.com/customservice/kfaccount/del?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html",
				FuncName:    "KfaccountDel",
			},
			{
				Name:        "设置客服帐号的头像",
				Description: "开发者可调用本接口来上传图片作为客服人员的头像，头像图片文件必须是jpg格式，推荐使用640*640大小的图片以达到最佳效果",
				Request:     "POST http://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=ACCESS_TOKEN&kf_account=KFACCOUNT",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/customservice/getkflist",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/message/custom/send",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/message/custom/typing",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/custom/typing?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html",
				FuncName:    "",
			},
			// 新版客服功能
			{
				Name:        "/cgi-bin/customservice/getonlinekflist",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html",
				FuncName:    "",
			},
			{
				Name:        "/customservice/kfaccount/inviteworker",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/customservice/kfaccount/inviteworker?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html",
				FuncName:    "",
			},

			{
				Name:        "/customservice/kfsession/create",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/customservice/kfsession/create?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html",
				FuncName:    "",
			},
			{
				Name:        "/customservice/kfsession/getsession",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/customservice/kfsession/getsession?access_token=ACCESS_TOKEN&openid=OPENID",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html",
				FuncName:    "",
			},
			{
				Name:        "/customservice/kfsession/getsessionlist",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/customservice/kfsession/getsessionlist?access_token=ACCESS_TOKEN&kf_account=KFACCOUNT",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html",
				FuncName:    "",
			},
			{
				Name:        "/customservice/kfsession/getwaitcase",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/customservice/kfsession/getwaitcase?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html",
				FuncName:    "",
			},

			{
				Name:        "/customservice/msgrecord/",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/customservice/msgrecord/",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Obtain_chat_transcript.html",
				FuncName:    "",
			},
		},
	},
	{
		Name:    `群发消息`,
		Package: `message/mass`,
		Apis: []Api{

			{
				Name:        "/cgi-bin/media/uploadimg",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/media/uploadnews",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/message/mass/sendall",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/media/uploadvideo",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/media/uploadvideo?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/message/mass/send",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/message/mass/delete",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/mass/delete?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/message/mass/preview",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/mass/preview?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/message/mass/get",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/mass/get?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/message/mass/speed/get",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/mass/speed/get?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/message/mass/speed/set",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/mass/speed/set?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html",
				FuncName:    "",
			},
		},
	},

	{
		Name:    `模板消息`,
		Package: `message/template`,
		Apis: []Api{

			{
				Name:        "/cgi-bin/template/api_set_industry",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/template/get_industry",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/template/get_industry?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/template/api_add_template",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/template/api_add_template?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/template/get_all_private_template",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/template/del_private_template",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/message/template/send",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html",
				FuncName:    "",
			},

			{
				Name:        "/cgi-bin/message/template/subscribe",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/template/subscribe?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/One-time_subscription_info.html",
				FuncName:    "",
			},
		},
	},
	{
		Name:    "素材管理",
		Package: "material",
		Apis: []Api{
			{
				Name:        "新增临时素材",
				Description: "公众号经常有需要用到一些临时性的多媒体素材的场景，例如在使用接口特别是发送消息时，对多媒体文件、多媒体消息的获取和调用等操作，是通过media_id来进行的。素材管理接口对所有认证的订阅号和服务号开放。通过本接口，公众号可以新增临时素材（即上传临时多媒体文件）",
				Request:     "POST(@media) https://api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html",
			},
			{
				Name:        "获取临时素材",
				Description: "公众号可以使用本接口获取临时素材（即下载临时的多媒体文件）",
				Request:     "GET https://api.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id=MEDIA_ID",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_temporary_materials.html",
			},
			{
				Name:        "高清语音素材获取接口",
				Description: "公众号可以使用本接口获取从JSSDK的uploadVoice接口上传的临时语音素材，格式为speex，16K采样率。该音频比上文的临时素材获取接口（格式为amr，8K采样率）更加清晰，适合用作语音识别等对音质要求较高的业务",
				Request:     "GET https://api.weixin.qq.com/cgi-bin/media/get/jssdk?access_token=ACCESS_TOKEN&media_id=MEDIA_ID",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_temporary_materials.html",
			},
			{
				Name:        "新增永久图文素材",
				Description: "对于常用的素材，开发者可通过本接口上传到微信服务器，永久使用。新增的永久素材也可以在公众平台官网素材管理模块中查询管理",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html",
			},
			{
				Name:        "上传图文消息内的图片获取URL",
				Description: "本接口所上传的图片不占用公众号的素材库中图片数量的100000个的限制。图片仅支持jpg/png格式，大小必须在1MB以下",
				Request:     "POST(@media) https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html",
			},
			{
				Name:        "新增其他类型永久素材",
				Description: "通过POST表单来调用接口，表单id为media，包含需要上传的素材内容，有filename、filelength、content-type等信息。请注意：图片素材将进入公众平台官网素材管理模块中的默认分组",
				Request:     "POST(@media|field=description) https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=ACCESS_TOKEN&type=TYPE",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html",
			},
			{
				Name:        "获取永久素材",
				Description: `在新增了永久素材后，开发者可以根据media_id通过本接口下载永久素材。公众号在公众平台官网素材管理模块中新建的永久素材，可通过"获取素材列表"获知素材的media_id。请注意：临时素材无法通过本接口获取`,
				Request:     "POST https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Getting_Permanent_Assets.html",
			},
			{
				Name:        "删除永久素材",
				Description: `在新增了永久素材后，开发者可以根据本接口来删除不再需要的永久素材，节省空间`,
				Request:     "POST https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Deleting_Permanent_Assets.html",
			},
			{
				Name:        "修改永久图文素材",
				Description: `开发者可以通过本接口对永久图文素材进行修改`,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/material/update_news?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Editing_Permanent_Rich_Media_Assets.html`,
			},
			{
				Name:        `获取素材总数`,
				Description: `1.永久素材的总数，也会计算公众平台官网素材管理中的素材 2.图片和图文消息素材（包括单图文和多图文）的总数上限为5000，其他素材的总数上限为1000`,
				Request:     `GET https://api.weixin.qq.com/cgi-bin/material/get_materialcount?access_token=ACCESS_TOKEN`,
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_the_total_of_all_materials.html",
			},
			{
				Name:        `获取素材列表`,
				Description: `在新增了永久素材后，开发者可以分类型获取永久素材的列表`,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=ACCESS_TOKEN`,
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_materials_list.html",
			},
		},
	},
	{
		Name:    "图文留言管理",
		Package: "comment",
		Apis: []Api{
			{
				Name:        `打开已群发文章评论`,
				Description: ``,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/comment/open?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html`,
			},
			{
				Name:        `关闭已群发文章评论`,
				Description: ``,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/comment/close?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html`,
			},
			{
				Name:        `查看指定文章的评论数据`,
				Description: ``,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/comment/list?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html`,
			},
			{
				Name:        ` 将评论标记精选`,
				Description: ``,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/comment/markelect?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html`,
			},
			{
				Name:        ` 将评论取消精选`,
				Description: ``,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/comment/unmarkelect?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html`,
			},
			{
				Name:        ` 删除评论`,
				Description: ``,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/comment/delete?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html`,
			},
			{
				Name:        ` 回复评论`,
				Description: ``,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/comment/reply/add?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html`,
				FuncName:    "ReplyAdd",
			},
			{
				Name:        ` 删除回复`,
				Description: ``,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/comment/reply/delete?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html`,
				FuncName:    "ReplyDelete",
			},
		},
	},
	{
		Name:    "用户标签管理",
		Package: "tag",
		Apis: []Api{
			{
				Name:        `创建标签`,
				Description: `一个公众号，最多可以创建100个标签`,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/tags/create?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html`,
			},
			{
				Name:        `获取公众号已创建的标签`,
				Description: ``,
				Request:     `GET https://api.weixin.qq.com/cgi-bin/tags/get?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html`,
			},
			{
				Name:        `编辑标签`,
				Description: ``,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/tags/update?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html`,
			},
			{
				Name:        `删除标签`,
				Description: `请注意，当某个标签下的粉丝超过10w时，后台不可直接删除标签。此时，开发者可以对该标签下的openid列表，先进行取消标签的操作，直到粉丝数不超过10w后，才可直接删除该标签`,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html`,
			},
			{
				Name:        `获取标签下粉丝列表`,
				Description: ``,
				Request:     `GET https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html`,
				FuncName:    "GetUsers",
			},
			{
				Name:        `批量为用户打标签`,
				Description: `标签功能目前支持公众号为用户打上最多20个标签`,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html`,
				FuncName:    "BatchTagging",
			},
			{
				Name:        `批量为用户取消标签`,
				Description: `标签功能目前支持公众号为用户打上最多20个标签`,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html`,
				FuncName:    "BatchUntagging",
			},
			{
				Name:        `获取用户身上的标签列表`,
				Description: `标签功能目前支持公众号为用户打上最多20个标签`,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/tags/getidlist?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html`,
				FuncName:    `GetIdList`,
			},
		},
	},
	{
		Name:    `用户管理`,
		Package: `user`,
		Apis: []Api{
			{
				Name:        "设置用户备注名",
				Description: `开发者可以通过该接口对指定用户设置备注名，该接口暂时开放给微信认证的服务号`,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/Configuring_user_notes.html`,
			},
			{
				Name:        "获取用户基本信息(UnionID机制)",
				Description: `在关注者与公众号产生消息交互后，公众号可获得关注者的OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的。对于不同公众号，同一用户的openid不同）。公众号可通过本接口来根据OpenID获取用户基本信息，包括昵称、头像、性别、所在城市、语言和关注时间`,
				Request:     `GET https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId`,
			},
			{
				Name:        `批量获取用户基本信息`,
				Description: `开发者可通过该接口来批量获取用户基本信息。最多支持一次拉取100条`,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId`,
			},
			{
				Name:        `批量用户列表`,
				Description: `公众号可通过本接口来获取帐号的关注者列表，关注者列表由一串OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成。一次拉取调用最多拉取10000个关注者的OpenID，可以通过多次拉取的方式来满足需求`,
				Request:     `GET https://api.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&next_openid=NEXT_OPENID`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/Getting_a_User_List.html`,
			},
			{
				Name:        `获取公众号的黑名单列表`,
				Description: `公众号可通过该接口来获取帐号的黑名单列表，黑名单列表由一串 OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成`,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html`,
			},
			{
				Name:        `拉黑用户`,
				Description: `公众号可通过该接口来拉黑一批用户，黑名单列表由一串 OpenID （加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成`,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html`,
			},
			{
				Name:        `取消拉黑用户`,
				Description: `公众号可通过该接口来取消拉黑一批用户，黑名单列表由一串OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成`,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html`,
			},
		},
	},
	{
		Name:    `账号管理`,
		Package: `account`,
		Apis: []Api{

			{
				Name:        "/cgi-bin/qrcode/create",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Account_Management/Generating_a_Parametric_QR_Code.html",
				FuncName:    "",
			},

			{
				Name:        "/cgi-bin/shorturl",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/shorturl?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Account_Management/URL_Shortener.html",
				FuncName:    "",
			},
		},
	},
	{
		Name:    `数据统计`,
		Package: `datacube`,
		Apis: []Api{

			{
				Name:        "/datacube/getusersummary",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getusersummary?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/User_Analysis_Data_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/datacube/getusercumulate",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getusercumulate?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/User_Analysis_Data_Interface.html",
				FuncName:    "",
			},

			{
				Name:        "/datacube/getarticlesummary",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getarticlesummary?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/datacube/getarticletotal",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getarticletotal?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/datacube/getuserread",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getuserread?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/datacube/getuserreadhour",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getuserreadhour?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/datacube/getusershare",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getusershare?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/datacube/getusersharehour",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getusersharehour?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html",
				FuncName:    "",
			},

			{
				Name:        "/datacube/getupstreammsg",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getupstreammsg?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html",
				FuncName:    "",
			},
			{
				Name:        "/datacube/getupstreammsghour",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getupstreammsghour?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html",
				FuncName:    "",
			},
			{
				Name:        "/datacube/getupstreammsgweek",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getupstreammsgweek?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html",
				FuncName:    "",
			},
			{
				Name:        "/datacube/getupstreammsgmonth",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getupstreammsgmonth?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html",
				FuncName:    "",
			},
			{
				Name:        "/datacube/getupstreammsgdist",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getupstreammsgdist?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html",
				FuncName:    "",
			},
			{
				Name:        "/datacube/getupstreammsgdistweek",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getupstreammsgdistweek?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html",
				FuncName:    "",
			},
			{
				Name:        "/datacube/getupstreammsgdistmonth",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getupstreammsgdistmonth?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html",
				FuncName:    "",
			},

			{
				Name:        "/publisher/stat",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/publisher/stat?action=publisher_adpos_general&access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Ad_Analysis.html",
				FuncName:    "",
			},

			{
				Name:        "/datacube/getinterfacesummary",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getinterfacesummary?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Analytics_API.html",
				FuncName:    "",
			},
			{
				Name:        "/datacube/getinterfacesummaryhour",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getinterfacesummaryhour?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Analytics_API.html",
				FuncName:    "",
			},
		},
	},
	{
		Name:    `微信卡券`,
		Package: `card`,
		Apis: []Api{

			{
				Name:        "/card/create",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/create?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Create_a_Coupon_Voucher_or_Card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/paycell/set",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/paycell/set?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Create_a_Coupon_Voucher_or_Card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/selfconsumecell/set",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/selfconsumecell/set?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Create_a_Coupon_Voucher_or_Card.html",
				FuncName:    "",
			},

			{
				Name:        "/card/qrcode/create",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/qrcode/create?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Distributing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "",
			},
			{
				Name:        "/card/landingpage/create",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/landingpage/create?access_token=$TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Distributing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "",
			},
			{
				Name:        "/card/mpnews/gethtml",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/mpnews/gethtml?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Distributing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "",
			},
			{
				Name:        "/card/testwhitelist/set",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/testwhitelist/set?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Distributing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "",
			},

			{
				Name:        "/card/code/get",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/code/get?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Redeeming_a_coupon_voucher_or_card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/code/consume",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/code/consume?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Redeeming_a_coupon_voucher_or_card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/code/decrypt",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/code/decrypt?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Redeeming_a_coupon_voucher_or_card.html",
				FuncName:    "",
			},

			{
				Name:        "/card/user/getcardlist",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/user/getcardlist?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "",
			},
			{
				Name:        "/card/get",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/get?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "",
			},
			{
				Name:        "/card/batchget",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/batchget?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "",
			},
			{
				Name:        "/card/update",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/update?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "",
			},
			{
				Name:        "/card/modifystock",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/modifystock?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "",
			},
			{
				Name:        "/card/code/update",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/code/update?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "",
			},
			{
				Name:        "/card/delete",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/delete?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "",
			},
			{
				Name:        "/card/code/unavailable",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/code/unavailable?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "",
			},
			{
				Name:        "/datacube/getcardbizuininfo",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getcardbizuininfo?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "",
			},
			{
				Name:        "/datacube/getcardcardinfo",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getcardcardinfo?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "",
			},
			{
				Name:        "/datacube/getcardmembercardinfo",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getcardmembercardinfo?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "",
			},
			{
				Name:        "/datacube/getcardmembercarddetail",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getcardmembercarddetail?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "",
			},

			{
				Name:        "/card/membercard/activate/geturl",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/membercard/activate/geturl?access_token=",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Coupons-Mini_Program_Start_Up.html",
				FuncName:    "",
			},
			{
				Name:        "/card/membercard/activatetempinfo/get",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/membercard/activatetempinfo/get?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Coupons-Mini_Program_Start_Up.html",
				FuncName:    "",
			},
			{
				Name:        "/card/membercard/activate",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/membercard/activate?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Coupons-Mini_Program_Start_Up.html",
				FuncName:    "",
			},

			{
				Name:        "/card/giftcard/page/add",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/page/add?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/giftcard/page/get",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/page/get?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/giftcard/page/update",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/page/update?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/giftcard/page/batchget",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/page/batchget?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/giftcard/maintain/set",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/maintain/set?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/giftcard/pay/whitelist/add",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/pay/whitelist/add?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/giftcard/pay/submch/bind",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/pay/submch/bind?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/giftcard/wxa/set",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/wxa/set",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/giftcard/wxa/set<span></span></a><em>",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/wxa/set",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/giftcard/order/get",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/order/get?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/giftcard/order/batchget",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/order/batchget?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/generalcard/updateuser",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/generalcard/updateuser?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/giftcard/order/refund",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/order/refund?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/invoice/setbizattr",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/setbizattr?action=set_pay_mch&access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/invoice/getauthdata",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/getauthdata",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/invoice/getauthdata<span></span></a>",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/getauthdata",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "",
			},

			{
				Name:        "/card/membercard/activateuserform/set",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/membercard/activateuserform/set?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Create_a_membership_card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/membercard/userinfo/get",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/membercard/userinfo/get?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Create_a_membership_card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/membercard/updateuser",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/membercard/updateuser?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Create_a_membership_card.html",
				FuncName:    "",
			},

			{
				Name:        "/card/paygiftcard/add",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/paygiftcard/add?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Manage_Member_Card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/paygiftcard/delete",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/paygiftcard/delete?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Manage_Member_Card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/paygiftcard/getbyid",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/paygiftcard/getbyid?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Manage_Member_Card.html",
				FuncName:    "",
			},
			{
				Name:        "/card/paygiftcard/batchget",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/paygiftcard/batchget?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Manage_Member_Card.html",
				FuncName:    "",
			},

			{
				Name:        "/card/meetingticket/updateuser",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/meetingticket/updateuser?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Special_ticket.html",
				FuncName:    "",
			},
			{
				Name:        "/card/movieticket/updateuser",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/movieticket/updateuser?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Special_ticket.html",
				FuncName:    "",
			},
			{
				Name:        "/card/boardingpass/checkin",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/boardingpass/checkin?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Special_ticket.html",
				FuncName:    "",
			},

			{
				Name:        "/card/submerchant/submit",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/submerchant/submit?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html",
				FuncName:    "",
			},
			{
				Name:        "/card/getapplyprotocol",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/getapplyprotocol?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html",
				FuncName:    "",
			},
			{
				Name:        "/card/submerchant/update",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/submerchant/update?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html",
				FuncName:    "",
			},
			{
				Name:        "/card/submerchant/get",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/submerchant/get?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html",
				FuncName:    "",
			},
			{
				Name:        "/card/submerchant/batchget",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/submerchant/batchget?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html",
				FuncName:    "",
			},
		},
	},
	{
		Name:    `微信门店`,
		Package: `poi`,
		Apis: []Api{

			{
				Name:        "/cgi-bin/poi/addpoi",
				Description: "",
				Request:     "POST http://api.weixin.qq.com/cgi-bin/poi/addpoi?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/poi/getpoi",
				Description: "",
				Request:     "POST http://api.weixin.qq.com/cgi-bin/poi/getpoi?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/poi/getpoilist",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/poi/getpoilist?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/poi/updatepoi",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/poi/updatepoi?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/poi/delpoi",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/poi/delpoi?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/poi/getwxcategory",
				Description: "",
				Request:     "POST http://api.weixin.qq.com/cgi-bin/poi/getwxcategory?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html",
				FuncName:    "",
			},
		},
	},
	{
		Name:    `微信门店小程序`,
		Package: `poi/miniprogram`,
		Apis: []Api{

			{
				Name:        "/wxa/getwxastorecatelist",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/getwxastorecatelist?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/wxa/get_merchant_category",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/get_merchant_category?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/wxa/applywxastore ",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/applywxastore%20?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/wxa/apply_merchant",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/apply_merchant?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/wxa/get_merchant_audit_info",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/get_merchant_audit_info?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/wxa/modify_merchant",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/modify_merchant?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/wxa/get_district",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/get_district?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/wxa/search_map_poi",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/search_map_poi?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/wxa/create_map_poi",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/create_map_poi?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/wxa/add_store",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/add_store?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/wxa/update_store",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/update_store?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/card/storewxa/get",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/storewxa/get?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/wxa/get_store_info",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/get_store_info?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/wxa/get_store_list",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/get_store_list?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "/wxa/del_store",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/del_store?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
		},
	},
	{
		Name:    `智能接口`,
		Package: `ai`,
		Apis: []Api{

			{
				Name:        "/semantic/semproxy/search",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/semantic/semproxy/search?access_token=YOUR_ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Natural_Language_Processing.html",
				FuncName:    "",
			},

			{
				Name:        "/cgi-bin/media/voice/addvoicetorecofortext",
				Description: "",
				Request:     "POST http://api.weixin.qq.com/cgi-bin/media/voice/addvoicetorecofortext?access_token=ACCESS_TOKEN&format=&voice_id=xxxxxx&lang=zh_CN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/AI_Open_API.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/media/voice/queryrecoresultfortext",
				Description: "",
				Request:     "POST http://api.weixin.qq.com/cgi-bin/media/voice/queryrecoresultfortext?access_token=ACCESS_TOKEN&voice_id=xxxxxx&lang=zh_CN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/AI_Open_API.html",
				FuncName:    "",
			},
			{
				Name:        "/cgi-bin/media/voice/translatecontent",
				Description: "",
				Request:     "POST http://api.weixin.qq.com/cgi-bin/media/voice/translatecontent?access_token=ACCESS_TOKEN&lfrom=xxx&lto=xxx",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/AI_Open_API.html",
				FuncName:    "",
			},

			{
				Name:        "/cv/ocr/idcard",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cv/ocr/idcard?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html",
				FuncName:    "",
			},
			{
				Name:        "/cv/ocr/",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cv/ocr/",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html",
				FuncName:    "",
			},
			{
				Name:        "/cv/ocr/drivinglicense",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cv/ocr/drivinglicense?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html",
				FuncName:    "",
			},
			{
				Name:        "/cv/ocr/bizlicense",
				Description: "",
				Request:     "POST http://api.weixin.qq.com/cv/ocr/bizlicense?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html",
				FuncName:    "",
			},
			{
				Name:        "/cv/ocr/comm",
				Description: "",
				Request:     "POST http://api.weixin.qq.com/cv/ocr/comm?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html",
				FuncName:    "",
			},

			{
				Name:        "/cv/img/qrcode",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cv/img/qrcode?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Img_Proc.html",
				FuncName:    "",
			},
			{
				Name:        "/cv/img/superresolution",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cv/img/superresolution?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Img_Proc.html",
				FuncName:    "",
			},
			{
				Name:        "/cv/img/aicrop",
				Description: "",
				Request:     "POST http://api.weixin.qq.com/cv/img/aicrop?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Img_Proc.html",
				FuncName:    "",
			},
		},
	},
	// 导购
	{
		Name:    `顾问管理`,
		Package: `guide/guide`,
		Apis: []Api{

			{
				Name:        "获取顾问信息",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/getguideacct?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideAcct.html",
				FuncName:    "GetGuideAcct",
			},

			{
				Name:        "为服务号添加顾问",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/addguideacct?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.addGuideAcct.html",
				FuncName:    "AddGuideAcct",
			},

			{
				Name:        "修改顾问的昵称或头像",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/updateguideacct?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.updateGuideAcct.html",
				FuncName:    "UpdateGuideAcct",
			},

			{
				Name:        "删除顾问",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/delguideacct?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.delGuideAcct.html",
				FuncName:    "DelGuideAcct",
			},

			{
				Name:        "获取服务号顾问列表",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/getguideacctlist?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideAcctList.html",
				FuncName:    "GetGuideAcctList",
			},

			{
				Name:        "生成顾问二维码",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/guidecreateqrcode?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.guideCreateQrCode.html",
				FuncName:    "GuideCreateQrCode",
			},

			{
				Name:        "/cgi-bin/guide/getguidebuyerchatrecord",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/getguidebuyerchatrecord?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideBuyerChatRecord.html",
				FuncName:    "GetGuideBuyerChatRecord",
			},

			{
				Name:        "设置快捷回复与关注自动回复",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/setguideconfig?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.setGuideConfig.html",
				FuncName:    "SetGuideConfig",
			},

			{
				Name:        "获取快捷回复与关注自动回复",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/getguideconfig?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideConfig.html",
				FuncName:    "GetGuideConfig",
			},

			{
				Name:        "为服务号设置敏感词与离线自动回复",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/setguideacctconfig?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.setGuideAcctConfig.html",
				FuncName:    "SetGuideAcctConfig",
			},

			{
				Name:        "获取离线自动回复与敏感词",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/getguideacctconfig?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideAcctConfig.html",
				FuncName:    "GetGuideAcctConfig",
			},

			{
				Name:        "允许微信用户复制小程序页面路径",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/pushshowwxapathmenu?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.pushShowWxaPathMenu.html",
				FuncName:    "PushShowWxaPathMenu",
			},

			{
				Name:        "新建顾问分组",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/newguidegroup?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.newGuideGroup.html",
				FuncName:    "NewGuideGroup",
			},

			{
				Name:        "获取服务号下所有顾问分组的列表",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/getguidegrouplist?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideGroupList.html",
				FuncName:    "GetGuideGroupList",
			},

			{
				Name:        "获取指定顾问分组信息，以及分组内顾问信息",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/getgroupinfo?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGroupInfo.html",
				FuncName:    "GetGroupInfo",
			},

			{
				Name:        "分组内添加顾问",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/addguide2guidegroup?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.addGuide2GuideGroup.html",
				FuncName:    "AddGuide2GuideGroup",
			},

			{
				Name:        "分组内删除顾问",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/delguide2guidegroup?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.delGuide2GuideGroup.html",
				FuncName:    "DelGuide2GuideGroup",
			},

			{
				Name:        "获取顾问所在分组",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/getgroupbyguide?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGroupByGuide.html",
				FuncName:    "GetGroupByGuide",
			},

			{
				Name:        "删除指定顾问分组",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/delguidegroup?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.delGuideGroup.html",
				FuncName:    "DelGuideGroup",
			},

			{
				Name:        "获取顾问的客户列表",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/getguidebuyerrelationlist?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.getGuideBuyerRelationList.html",
				FuncName:    "GetGuideBuyerRelationList",
			},
		},
	},
	{
		Name:    `客户管理`,
		Package: `guide/buyer`,
		Apis: []Api{

			{
				Name:        "为顾问分配客户",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/addguidebuyerrelation?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.addGuideBuyerRelation.html",
				FuncName:    "AddGuideBuyerRelation",
			},

			{
				Name:        "为顾问移除客户",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/delguidebuyerrelation?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.delGuideBuyerRelation.html",
				FuncName:    "DelGuideBuyerRelation",
			},

			{
				Name:        "为客户更换顾问",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/rebindguideacctforbuyer?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.rebindGuideAcctForBuyer.html",
				FuncName:    "RebindGuideAcctForBuyer",
			},

			{
				Name:        "修改客户昵称",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/updateguidebuyerrelation?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.updateGuideBuyerRelation.html",
				FuncName:    "UpdateGuideBuyerRelation",
			},

			{
				Name:        "查询客户所属顾问",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/getguidebuyerrelationbybuyer?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.getGuideBuyerRelationByBuyer.html",
				FuncName:    "GetGuideBuyerRelationByBuyer",
			},

			{
				Name:        "查询指定顾问和客户的关系",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/getguidebuyerrelation?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.getGuideBuyerRelation.html",
				FuncName:    "GetGuideBuyerRelation",
			},
		},
	},
	{
		Name:    `标签管理`,
		Package: `guide/tags`,
		Apis: []Api{

			{
				Name:        "/cgi-bin/guide/newguidetagoption",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/newguidetagoption?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.newGuideTagOption.html",
				FuncName:    "NewGuideTagOption",
			},

			{
				Name:        "删除指定标签类型",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/delguidetagoption?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.delguidetagoption.html",
				FuncName:    "Delguidetagoption",
			},

			{
				Name:        "为标签添加可选值",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/addguidetagoption?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.addGuideTagOption.html",
				FuncName:    "AddGuideTagOption",
			},

			{
				Name:        "获取标签和可选值",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/getguidetagoption?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.getGuideTagOption.html",
				FuncName:    "GetGuideTagOption",
			},

			{
				Name:        "为客户设置标签",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/addguidebuyertag?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.addGuideBuyerTag.html",
				FuncName:    "AddGuideBuyerTag",
			},

			{
				Name:        "查询客户标签",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/getguidebuyertag?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.getGuideBuyerTag.html",
				FuncName:    "GetGuideBuyerTag",
			},

			{
				Name:        "根据标签值筛选客户",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/queryguidebuyerbytag?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.queryGuideBuyerByTag.html",
				FuncName:    "QueryGuideBuyerByTag",
			},

			{
				Name:        "删除客户标签",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/delguidebuyertag?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.delGuideBuyerTag.html",
				FuncName:    "DelGuideBuyerTag",
			},

			{
				Name:        "设置自定义客户信息",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/addguidebuyerdisplaytag?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.addGuideBuyerDisplayTag.html",
				FuncName:    "AddGuideBuyerDisplayTag",
			},

			{
				Name:        "获取自定义客户信息",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/getguidebuyerdisplaytag?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.getGuideBuyerDisplayTag.html",
				FuncName:    "GetGuideBuyerDisplayTag",
			},
		},
	},
	{
		Name:    `素材管理`,
		Package: `guide/material`,
		Apis: []Api{

			{
				Name:        "添加小程序卡片素材",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/setguidecardmaterial?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.setGuideCardMaterial.html",
				FuncName:    "SetGuideCardMaterial",
			},

			{
				Name:        "查询小程序卡片素材",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/getguidecardmaterial?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.getGuideCardMaterial.html",
				FuncName:    "GetGuideCardMaterial",
			},

			{
				Name:        "删除小程序卡片素材",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/delguidecardmaterial?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.delGuideCardMaterial.html",
				FuncName:    "DelGuideCardMaterial",
			},

			{
				Name:        "添加图片素材",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/setguideimagematerial?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.setGuideImageMaterial.html",
				FuncName:    "SetGuideImageMaterial",
			},

			{
				Name:        "查询图片素材",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/getguideimagematerial?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.getGuideImageMaterial.html",
				FuncName:    "GetGuideImageMaterial",
			},

			{
				Name:        "删除图片素材",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/delguideimagematerial?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.delGuideImageMaterial.html",
				FuncName:    "DelGuideImageMaterial",
			},

			{
				Name:        "添加文字素材",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/setguidewordmaterial?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.setGuideWordMaterial.html",
				FuncName:    "SetGuideWordMaterial",
			},

			{
				Name:        "查询文字素材",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/getguidewordmaterial?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.getGuideWordMaterial.html",
				FuncName:    "GetGuideWordMaterial",
			},

			{
				Name:        "删除文字素材",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/delguidewordmaterial?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.delGuideWordMaterial.html",
				FuncName:    "DelGuideWordMaterial",
			},
		},
	},
	{
		Name:    `群发任务`,
		Package: `guide/job`,
		Apis: []Api{

			{
				Name:        "添加群发任务",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/addguidemassendjob?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/task-account/shopping-guide.addGuideMassendJob.html",
				FuncName:    "AddGuideMassendJob",
			},

			{
				Name:        "获取群发任务列表",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/getguidemassendjoblist?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/task-account/shopping-guide.getGuideMassendJobList.html",
				FuncName:    "GetGuideMassendJobList",
			},

			{
				Name:        "获取指定群发任务信息",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/getguidemassendjob?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/task-account/shopping-guide.getGuideMassendJob.html",
				FuncName:    "GetGuideMassendJob",
			},

			{
				Name:        "修改群发任务",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/updateguidemassendjob?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/task-account/shopping-guide.updateGuideMassendJob.html",
				FuncName:    "UpdateGuideMassendJob",
			},

			{
				Name:        "取消群发任务",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/guide/cancelguidemassendjob?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/task-account/shopping-guide.cancelGuideMassendJob.html",
				FuncName:    "CancelGuideMassendJob",
			},
		},
	},
	{
		Name:    `一物一码`,
		Package: `iot`,
		Apis: []Api{

			{
				Name:        "/intp/marketcode/applycode",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/intp/marketcode/applycode?access_token=ACCESSTOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html",
				FuncName:    "",
			},
			{
				Name:        "/intp/marketcode/applycodequery",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/intp/marketcode/applycodequery?access_token=ACCESSTOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html",
				FuncName:    "",
			},
			{
				Name:        "/intp/marketcode/applycodedownload",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/intp/marketcode/applycodedownload?access_token=ACCESSTOKE",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html",
				FuncName:    "",
			},
			{
				Name:        "/intp/marketcode/codeactive",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/intp/marketcode/codeactive?access_token=ACCESSTOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html",
				FuncName:    "",
			},
			{
				Name:        "/intp/marketcode/codeactivequery",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/intp/marketcode/codeactivequery?access_token=ACCESSTOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html",
				FuncName:    "",
			},
			{
				Name:        "/intp/marketcode/tickettocode",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/intp/marketcode/tickettocode?access_token=ACCESSTOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html",
				FuncName:    "",
			},
		},
	},
	{
		Name:    `微信发票`,
		Package: `invoice`,
		Apis: []Api{

			{
				Name:        "/card/invoice/getauthurl",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/getauthurl?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html",
				FuncName:    "",
			},
			{
				Name:        "/card/invoice/getauthdata",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/getauthdata?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html",
				FuncName:    "",
			},
			{
				Name:        "/card/invoice/rejectinsert",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/rejectinsert?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html",
				FuncName:    "",
			},
			{
				Name:        "/card/invoice/makeoutinvoice",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/makeoutinvoice?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html",
				FuncName:    "",
			},
			{
				Name:        "/card/invoice/clearoutinvoice",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/clearoutinvoice?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html",
				FuncName:    "",
			},
			{
				Name:        "/card/invoice/queryinvoceinfo",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/queryinvoceinfo?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html",
				FuncName:    "",
			},

			{
				Name:        "/card/invoice/seturl",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/seturl?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html",
				FuncName:    "",
			},
			{
				Name:        "/card/invoice/platform/createcard",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/platform/createcard?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html",
				FuncName:    "",
			},
			{
				Name:        "/card/invoice/platform/setpdf",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/platform/setpdf?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html",
				FuncName:    "",
			},
			{
				Name:        "/card/invoice/platform/getpdf",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/platform/getpdf?action=get_url&access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html",
				FuncName:    "",
			},
			{
				Name:        "/card/invoice/insert",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/insert?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html",
				FuncName:    "",
			},
			{
				Name:        "/card/invoice/platform/updatestatus",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/platform/updatestatus?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html",
				FuncName:    "",
			},

			{
				Name:        "/card/invoice/reimburse/getinvoiceinfo",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/reimburse/getinvoiceinfo?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Reimburser_API_List.html",
				FuncName:    "",
			},
			{
				Name:        "/card/invoice/reimburse/getinvoicebatch",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/reimburse/getinvoicebatch?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Reimburser_API_List.html",
				FuncName:    "",
			},
			{
				Name:        "/card/invoice/reimburse/updateinvoicestatus",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/reimburse/updateinvoicestatus?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Reimburser_API_List.html",
				FuncName:    "",
			},
			{
				Name:        "/card/invoice/reimburse/updatestatusbatch",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/reimburse/updatestatusbatch?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Reimburser_API_List.html",
				FuncName:    "",
			},

			{
				Name:        "/card/invoice/biz/getusertitleurl",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/biz/getusertitleurl?access_token={access_token",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Quick_issuing/Interface_Instructions.html",
				FuncName:    "",
			},
			{
				Name:        "/card/invoice/getauthdata)card/invoice/biz/getselecttitleurl",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/getauthdata)card/invoice/biz/getselecttitleurl?access_token={access_token/",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Quick_issuing/Interface_Instructions.html",
				FuncName:    "",
			},
			{
				Name:        "/card/invoice/biz/getselecttitleurl",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/biz/getselecttitleurl?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Quick_issuing/Interface_Instructions.html",
				FuncName:    "",
			},
			{
				Name:        "/card/invoice/scantitle",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/scantitle?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Quick_issuing/Interface_Instructions.html",
				FuncName:    "",
			},
		},
	},
	{
		Name:    `非税缴费`,
		Package: `nontax`,
		Apis: []Api{

			{
				Name:        "/nontax/getbillauthurl",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/getbillauthurl?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Nontax_Bill/API_list.html",
				FuncName:    "",
			},
			{
				Name:        "/nontax/createbillcard",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/createbillcard?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Nontax_Bill/API_list.html",
				FuncName:    "",
			},
			{
				Name:        "/nontax/insertbill",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/insertbill?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Nontax_Bill/API_list.html",
				FuncName:    "",
			},

			{
				Name:        "/nontax/queryfee",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/queryfee?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "",
			},
			{
				Name:        "/nontax/unifiedorder",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/unifiedorder?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "",
			},
			{
				Name:        "/nontax/getorder",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/getorder?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "",
			},
			{
				Name:        "/nontax/refund",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/refund?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "",
			},
			{
				Name:        "/nontax/downloadbill",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/downloadbill?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "",
			},
			{
				Name:        "/nontax/notifyinconsistentorder",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/notifyinconsistentorder?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "",
			},
			{
				Name:        "/nontax/mocknotification",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/mocknotification?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "",
			},
			{
				Name:        "/nontax/mockqueryfee",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/mockqueryfee?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "",
			},
			{
				Name:        "/nontax/micropay",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/micropay?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "",
			},
			{
				Name:        "/nontax/getorderlist",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/getorderlist?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "",
			},
			{
				Name:        "/intp/realname/getauthurl",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/intp/realname/getauthurl?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "",
			},
			{
				Name:        "/nontax/getrealname",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/getrealname?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "",
			},

			{
				Name:        "/nontax/vehicle/querystate",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/vehicle/querystate?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/Non-tax_driver_platform.html",
				FuncName:    "",
			},
			{
				Name:        "/nontax/vehicle/entrancenotify",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/vehicle/entrancenotify?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/Non-tax_driver_platform.html",
				FuncName:    "",
			},
			{
				Name:        "/nontax/vehicle/payapply",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/vehicle/payapply?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/Non-tax_driver_platform.html",
				FuncName:    "",
			},
		},
	},
}
