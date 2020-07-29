package main

type Param struct {
	Name string
	Type string
}

type Api struct {
	Name        string
	Description string
	Request     string
	See         string
	FuncName    string
	GetParams   []Param
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
		Name:    `消息管理`,
		Package: `message`,
		Apis: []Api{
			{
				Name:        "获取公众号的自动回复规则",
				Description: "开发者可以通过该接口，获取公众号当前使用的自动回复规则，包括关注后自动回复、消息自动回复（60分钟内触发一次）、关键词自动回复",
				Request:     "GET https://api.weixin.qq.com/cgi-bin/get_current_autoreply_info?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Getting_Rules_for_Auto_Replies.html",
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
				Request:     "POST(@media) http://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=ACCESS_TOKEN&kf_account=KFACCOUNT",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html",
				FuncName:    "UploadHeadImg",
				GetParams:   []Param{{Name: `kf_account`, Type: `string`}},
			},
			{
				Name:        "获取所有客服账号",
				Description: "开发者通过本接口，获取公众号中所设置的客服基本信息，包括客服工号、客服昵称、客服登录账号",
				Request:     "GET https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html",
				FuncName:    "GetKfList",
			},
			{
				Name:        "发消息",
				Description: "发送文本/图片/语言/语音/视频/音乐/图文/菜单/卡券... 消息",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html",
				FuncName:    "",
			},
			{
				Name:        "客服输入状态",
				Description: "开发者可通过调用“客服输入状态”接口，返回客服当前输入状态给用户",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/custom/typing?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Service_Center_messages.html",
				FuncName:    "",
			},
			// 新版客服功能
			{
				Name:        "获取客服基本信息",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html",
				FuncName:    "GetOnlineKfList",
			},
			{
				Name:        "邀请绑定客服帐号",
				Description: "新添加的客服帐号是不能直接使用的，只有客服人员用微信号绑定了客服账号后，方可登录Web客服进行操作。此接口发起一个绑定邀请到客服人员微信号，客服人员需要在微信客户端上用该微信号确认后帐号才可用。尚未绑定微信号的帐号可以进行绑定邀请操作，邀请未失效时不能对该帐号进行再次绑定微信号邀请",
				Request:     "POST https://api.weixin.qq.com/customservice/kfaccount/inviteworker?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Customer_Service_Management.html",
				FuncName:    "InviteWorker",
			},
			{
				Name:        "创建会话",
				Description: "此接口在客服和用户之间创建一个会话，如果该客服和用户会话已存在，则直接返回0。指定的客服帐号必须已经绑定微信号且在线",
				Request:     "POST https://api.weixin.qq.com/customservice/kfsession/create?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html",
				FuncName:    "KfSessionCreate",
			},
			{
				Name:        "获取客户会话状态",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/customservice/kfsession/getsession?access_token=ACCESS_TOKEN&openid=OPENID",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html",
				FuncName:    "KfSessionGet",
				GetParams:   []Param{{Name: `openid`, Type: `string`}},
			},
			{
				Name:        "获取客服会话列表",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/customservice/kfsession/getsessionlist?access_token=ACCESS_TOKEN&kf_account=KFACCOUNT",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html",
				FuncName:    "KfSessionGetList",
				GetParams:   []Param{{Name: `kf_account`, Type: `string`}},
			},
			{
				Name:        "获取未接入会话列表",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/customservice/kfsession/getwaitcase?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Session_control.html",
				FuncName:    "KfSessionGetWaitCase",
			},
			{
				Name:        "获取聊天记录",
				Description: "此接口返回的聊天记录中，对于图片、语音、视频，分别展示成文本格式的[image]、[voice]、[video]",
				Request:     "POST https://api.weixin.qq.com/customservice/msgrecord/getmsglist",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Customer_Service/Obtain_chat_transcript.html",
				FuncName:    "GetMsgList",
			},
		},
	},
	{
		Name:    `群发消息`,
		Package: `message/mass`,
		Apis: []Api{

			{
				Name:        "上传图文消息素材",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html",
				FuncName:    "MediaUploadNews",
			},
			{
				Name:        "根据标签进行群发",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html",
				FuncName:    "SendAll",
			},
			{
				Name:        "上传视频素材",
				Description: "群发 视频 的 media_id 需通过此接口特别地得到",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/media/uploadvideo?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html",
				FuncName:    "MediaUploadVideo",
			},
			{
				Name:        "根据OpenID列表群发",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html",
				FuncName:    "",
			},
			{
				Name:        "删除群发",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/mass/delete?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html",
				FuncName:    "",
			},
			{
				Name:        "预览",
				Description: "开发者可通过该接口发送消息给指定用户，在手机端查看消息的样式和排版。为了满足第三方平台开发者的需求，在保留对openID预览能力的同时，增加了对指定微信号发送预览的能力，但该能力每日调用次数有限制（100次），请勿滥用",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/mass/preview?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html",
				FuncName:    "",
			},
			{
				Name:        "查询群发消息发送状态",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/mass/get?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html",
				FuncName:    "",
			},
			{
				Name:        "获取群发速度",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/mass/speed/get?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html",
				FuncName:    "SpeedGet",
			},
			{
				Name:        "设置群发速度",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/mass/speed/set?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html",
				FuncName:    "SpeedSet",
			},
		},
	},

	{
		Name:    `模板消息`,
		Package: `message/template`,
		Apis: []Api{

			{
				Name:        "设置所属行业",
				Description: "设置行业可在微信公众平台后台完成，每月可修改行业1次，帐号仅可使用所属行业中相关的模板，为方便第三方开发者，提供通过接口调用的方式来修改账号所属行业",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html",
				FuncName:    "SetIndustry",
			},
			{
				Name:        "获取设置的行业信息",
				Description: "获取帐号设置的行业信息。可登录微信公众平台，在公众号后台中查看行业信息。为方便第三方开发者，提供通过接口调用的方式来获取帐号所设置的行业信息",
				Request:     "GET https://api.weixin.qq.com/cgi-bin/template/get_industry?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html",
				FuncName:    "GetIndustry",
			},
			{
				Name:        "获得模板ID",
				Description: "从行业模板库选择模板到帐号后台，获得模板ID的过程可在微信公众平台后台完成。为方便第三方开发者，提供通过接口调用的方式来获取模板ID",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/template/api_add_template?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html",
				FuncName:    "AddTemplate",
			},
			{
				Name:        "获取模板列表",
				Description: "获取已添加至帐号下所有模板列表，可在微信公众平台后台中查看模板列表信息。为方便第三方开发者，提供通过接口调用的方式来获取帐号下所有模板信息",
				Request:     "GET https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "删除模板",
				Description: "删除模板可在微信公众平台后台完成，为方便第三方开发者，提供通过接口调用的方式来删除某帐号下的模板",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "发送模板消息",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html",
				FuncName:    "",
			},

			{
				Name:        "推送订阅模板消息给到授权微信用户",
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
				FuncName:    `MediaUpload`,
			},
			{
				Name:        "获取临时素材",
				Description: "公众号可以使用本接口获取临时素材（即下载临时的多媒体文件）",
				Request:     "GET https://api.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id=MEDIA_ID",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_temporary_materials.html",
				FuncName:    `MediaGet`,
			},
			{
				Name:        "高清语音素材获取接口",
				Description: "公众号可以使用本接口获取从JSSDK的uploadVoice接口上传的临时语音素材，格式为speex，16K采样率。该音频比上文的临时素材获取接口（格式为amr，8K采样率）更加清晰，适合用作语音识别等对音质要求较高的业务",
				Request:     "GET https://api.weixin.qq.com/cgi-bin/media/get/jssdk?access_token=ACCESS_TOKEN&media_id=MEDIA_ID",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_temporary_materials.html",
				FuncName:    `MediaGetJssdk`,
				GetParams:   []Param{{Name: `media_id`, Type: `string`}},
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
				FuncName:    "MediaUploadImg",
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
				FuncName:    `GetMaterialCount`,
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
		Name:    "图文消息留言管理",
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
				FuncName:    `MarkElect`,
			},
			{
				Name:        ` 将评论取消精选`,
				Description: ``,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/comment/unmarkelect?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html`,
				FuncName:    `UnMarkElect`,
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
		Package: "user/tags",
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
				Request:     `POST https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html`,
				FuncName:    "GetUsersByTag",
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
				FuncName:    "BatchUnTagging",
			},
			{
				Name:        `获取用户身上的标签列表`,
				Description: `标签功能目前支持公众号为用户打上最多20个标签`,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/tags/getidlist?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/User_Tag_Management.html`,
				FuncName:    `GetTagIdList`,
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
				FuncName:    `UpdateRemark`,
			},
			{
				Name:        "获取用户基本信息",
				Description: `在关注者与公众号产生消息交互后，公众号可获得关注者的OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的。对于不同公众号，同一用户的openid不同）。公众号可通过本接口来根据OpenID获取用户基本信息，包括昵称、头像、性别、所在城市、语言和关注时间`,
				Request:     `GET https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId`,
				FuncName:    `GetUserInfo`,
				GetParams:   []Param{{Name: `openid`, Type: `string`}, {Name: `lang`, Type: `string`}},
			},
			{
				Name:        `批量获取用户基本信息`,
				Description: `开发者可通过该接口来批量获取用户基本信息。最多支持一次拉取100条`,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId`,
				FuncName:    `BatchGetUserInfo`,
			},
			{
				Name:        `获取用户列表`,
				Description: `公众号可通过本接口来获取帐号的关注者列表，关注者列表由一串OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成。一次拉取调用最多拉取10000个关注者的OpenID，可以通过多次拉取的方式来满足需求`,
				Request:     `GET https://api.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&next_openid=NEXT_OPENID`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/Getting_a_User_List.html`,
				GetParams:   []Param{{Name: `next_openid`, Type: `string`}},
			},
			{
				Name:        `获取公众号的黑名单列表`,
				Description: `公众号可通过该接口来获取帐号的黑名单列表，黑名单列表由一串 OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成`,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html`,
				FuncName:    `GetBlackList`,
			},
			{
				Name:        `拉黑用户`,
				Description: `公众号可通过该接口来拉黑一批用户，黑名单列表由一串 OpenID （加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成`,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html`,
				FuncName:    `BatchBlackList`,
			},
			{
				Name:        `取消拉黑用户`,
				Description: `公众号可通过该接口来取消拉黑一批用户，黑名单列表由一串OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成`,
				Request:     `POST https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist?access_token=ACCESS_TOKEN`,
				See:         `https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html`,
				FuncName:    `BatchUnBlackList`,
			},
		},
	},
	{
		Name:    `账号管理`,
		Package: `account`,
		Apis: []Api{

			{
				Name:        "创建二维码ticket",
				Description: "每次创建二维码ticket需要提供一个开发者自行设定的参数",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Account_Management/Generating_a_Parametric_QR_Code.html",
				FuncName:    "",
			},

			{
				Name:        "长链接转短链接",
				Description: "将一条长链接转成短链接。主要使用场景： 开发者用于生成二维码的原链接（商品、支付二维码等）太长导致扫码速度和成功率下降，将原长链接通过此接口转成短链接再生成二维码将大大提升扫码速度和成功率",
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
				Name:        "获取用户增减数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getusersummary?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/User_Analysis_Data_Interface.html",
				FuncName:    "GetUserSummary",
			},
			{
				Name:        "获取累计用户数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getusercumulate?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/User_Analysis_Data_Interface.html",
				FuncName:    "GetUserCumulate",
			},

			{
				Name:        "获取图文群发每日数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getarticlesummary?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html",
				FuncName:    "GetArticleSummary",
			},
			{
				Name:        "获取图文群发总数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getarticletotal?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html",
				FuncName:    "GetArticleTotal",
			},
			{
				Name:        "获取图文统计数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getuserread?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html",
				FuncName:    "GetUserRead",
			},
			{
				Name:        "获取图文统计分时数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getuserreadhour?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html",
				FuncName:    "GetUserReadHour",
			},
			{
				Name:        "获取图文分享转发数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getusershare?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html",
				FuncName:    "GetUserShare",
			},
			{
				Name:        "获取图文分享转发分时数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getusersharehour?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html",
				FuncName:    "GetUserShareHour",
			},

			{
				Name:        "获取消息发送概况数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getupstreammsg?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html",
				FuncName:    "GetUpstreamMsg",
			},
			{
				Name:        "获取消息分送分时数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getupstreammsghour?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html",
				FuncName:    "GetUpstreamMsgHour",
			},
			{
				Name:        "获取消息发送周数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getupstreammsgweek?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html",
				FuncName:    "GetUpstreamMsgWeek",
			},
			{
				Name:        "获取消息发送月数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getupstreammsgmonth?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html",
				FuncName:    "GetUpstreamMsgMonth",
			},
			{
				Name:        "获取消息发送分布数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getupstreammsgdist?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html",
				FuncName:    "GetUpstreamMsgDist",
			},
			{
				Name:        "获取消息发送分布周数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getupstreammsgdistweek?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html",
				FuncName:    "GetUpstreamMsgDistWeek",
			},
			{
				Name:        "获取消息发送分布月数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getupstreammsgdistmonth?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html",
				FuncName:    "GetUpstreamMsgDistMonth",
			},

			{
				Name:        "获取公众号分广告位数据/公众号返佣商品数据/公众号结算收入数据及结算主体信息",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/publisher/stat?action=publisher_adpos_general&access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Ad_Analysis.html",
				FuncName:    "PublisherStat",
				GetParams: []Param{
					{Name: `action`, Type: `string`},
				},
			},
			{
				Name:        "获取接口分析数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getinterfacesummary?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Analytics_API.html",
				FuncName:    "GetInterfaceSummary",
			},
			{
				Name:        "获取接口分析分时数据",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/datacube/getinterfacesummaryhour?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Analytics/Analytics_API.html",
				FuncName:    "GetInterfaceSummaryHour",
			},
		},
	},
	{
		Name:    `微信卡券`,
		Package: `card`,
		Apis: []Api{

			{
				Name:        "创建卡券",
				Description: "创建卡券接口是微信卡券的基础接口，用于创建一类新的卡券，获取card_id，创建成功并通过审核后，商家可以通过文档提供的其他接口将卡券下发给用户，每次成功领取，库存数量相应扣除",
				Request:     "POST https://api.weixin.qq.com/card/create?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Create_a_Coupon_Voucher_or_Card.html",
				FuncName:    "",
			},
			{
				Name:        "设置买单",
				Description: "创建卡券之后，开发者可以通过设置微信买单接口设置该card_id支持微信买单功能。值得开发者注意的是，设置买单的card_id必须已经配置了门店，否则会报错",
				Request:     "POST https://api.weixin.qq.com/card/paycell/set?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Create_a_Coupon_Voucher_or_Card.html",
				FuncName:    "SetPayCell",
			},
			{
				Name:        "设置自助核销",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/selfconsumecell/set?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Create_a_Coupon_Voucher_or_Card.html",
				FuncName:    "SetSelfConsumeCell",
			},

			{
				Name:        "创建二维码",
				Description: "开发者可调用该接口生成一张卡券二维码供用户扫码后添加卡券到卡包",
				Request:     "POST https://api.weixin.qq.com/card/qrcode/create?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Distributing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "CreateQRCode",
			},
			{
				Name:        "创建货架",
				Description: "开发者需调用该接口创建货架链接，用于卡券投放。创建货架时需填写投放路径的场景字段",
				Request:     "POST https://api.weixin.qq.com/card/landingpage/create?access_token=$TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Distributing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "CreateLandingPage",
			},
			{
				Name:        "图文消息群发卡券",
				Description: "支持开发者调用该接口获取卡券嵌入图文消息的标准格式代码，将返回代码填入 新增临时素材中content字段，即可获取嵌入卡券的图文消息素材",
				Request:     "POST https://api.weixin.qq.com/card/mpnews/gethtml?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Distributing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "MpnewsGetHtml",
			},
			{
				Name:        "设置测试白名单",
				Description: "由于卡券有审核要求，为方便公众号调试，可以设置一些测试帐号，这些帐号可领取未通过审核的卡券，体验整个流程",
				Request:     "POST https://api.weixin.qq.com/card/testwhitelist/set?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Distributing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "SetTestWhitelist",
			},

			{
				Name:        "线下核销-查询 Code",
				Description: "我们强烈建议开发者在调用核销code接口之前调用查询code接口，并在核销之前对非法状态的code(如转赠中、已删除、已核销等)做出处理",
				Request:     "POST https://api.weixin.qq.com/card/code/get?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Redeeming_a_coupon_voucher_or_card.html",
				FuncName:    "GetCode",
			},
			{
				Name:        "核销 Code",
				Description: "消耗code接口是核销卡券的唯一接口,开发者可以调用当前接口将用户的优惠券进行核销，该过程不可逆",
				Request:     "POST https://api.weixin.qq.com/card/code/consume?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Redeeming_a_coupon_voucher_or_card.html",
				FuncName:    "ConsumeCode",
			},
			{
				Name:        "线上核销 - Code解码",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/code/decrypt?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Redeeming_a_coupon_voucher_or_card.html",
				FuncName:    "DecryptCode",
			},

			{
				Name:        "获取用户已领取卡券",
				Description: "用于获取用户卡包里的，属于该appid下所有可用卡券，包括正常状态和异常状态",
				Request:     "POST https://api.weixin.qq.com/card/user/getcardlist?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "GetUserCardList",
			},
			{
				Name:        "查看卡券详情",
				Description: "开发者可以调用该接口查询某个card_id的创建信息、审核状态以及库存数量",
				Request:     "POST https://api.weixin.qq.com/card/get?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "",
			},
			{
				Name:        "批量查询卡券列表",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/batchget?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "BatchGet",
			},
			{
				Name:        "更改卡券信息",
				Description: "支持更新所有卡券类型的部分通用字段及特殊卡券（会员卡、飞机票、电影票、会议门票）中特定字段的信息",
				Request:     "POST https://api.weixin.qq.com/card/update?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "",
			},
			{
				Name:        "修改库存",
				Description: "调用修改库存接口增减某张卡券的库存",
				Request:     "POST https://api.weixin.qq.com/card/modifystock?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "ModifyStock",
			},
			{
				Name:        "更改Code",
				Description: "为确保转赠后的安全性，微信允许自定义Code的商户对已下发的code进行更改。 注：为避免用户疑惑，建议仅在发生转赠行为后（发生转赠后，微信会通过事件推送的方式告知商户被转赠的卡券Code）对用户的Code进行更改",
				Request:     "POST https://api.weixin.qq.com/card/code/update?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "UpdateCode",
			},
			{
				Name:        "删除卡券",
				Description: "删除卡券接口允许商户删除任意一类卡券。删除卡券后，该卡券对应已生成的领取用二维码、添加到卡包JS API均会失效。 注意：如用户在商家删除卡券前已领取一张或多张该卡券依旧有效。即删除卡券不能删除已被用户领取，保存在微信客户端中的卡券",
				Request:     "POST https://api.weixin.qq.com/card/delete?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "",
			},
			{
				Name:        "设置卡券失效",
				Description: "为满足改票、退款等异常情况，可调用卡券失效接口将用户的卡券设置为失效状态",
				Request:     "POST https://api.weixin.qq.com/card/code/unavailable?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "UnavailableCoed",
			},
			{
				Name:        "拉取卡券概况数据",
				Description: "支持调用该接口拉取本商户的总体数据情况，包括时间区间内的各指标总量",
				Request:     "POST https://api.weixin.qq.com/datacube/getcardbizuininfo?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "GetCardBizUinInfo",
			},
			{
				Name:        "获取免费券数据",
				Description: "支持开发者调用该接口拉取免费券（优惠券、团购券、折扣券、礼品券）在固定时间区间内的相关数据",
				Request:     "POST https://api.weixin.qq.com/datacube/getcardcardinfo?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "GetCardCardInfo",
			},
			{
				Name:        "拉取会员卡概况数据",
				Description: "支持开发者调用该接口拉取公众平台创建的会员卡相关数据",
				Request:     "POST https://api.weixin.qq.com/datacube/getcardmembercardinfo?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "GetCardMemberCardInfo",
			},
			{
				Name:        "拉取单张会员卡数据",
				Description: "支持开发者调用该接口拉取API创建的会员卡数据情况",
				Request:     "POST https://api.weixin.qq.com/datacube/getcardmembercarddetail?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html",
				FuncName:    "GetCardMemberCardDetail",
			},
		},
	},
	{
		Name:    `微信礼品卡`,
		Package: `card/giftcard`,
		Apis: []Api{

			{
				Name:        "创建-礼品卡货架",
				Description: "开发者可以通过该接口创建一个礼品卡货架并且用于公众号、门店的礼品卡售卖",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/page/add?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "GiftCardPageAdd",
			},
			{
				Name:        "查询-礼品卡货架信息",
				Description: "开发者可以查询某个礼品卡货架信息",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/page/get?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "GiftCardPageGet",
			},
			{
				Name:        "修改-礼品卡货架信息",
				Description: "开发者可以通过该接口更新礼品卡货架信息",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/page/update?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "GiftCardPageUpdate",
			},
			{
				Name:        "查询-礼品卡货架列表",
				Description: "开发者可以通过该接口查询当前商户下所有的礼品卡货架id",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/page/batchget?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "GiftCardPageBatchGet",
			},
			{
				Name:        "下架-礼品卡货架",
				Description: "开发者可以通过该接口查询当前商户下所有的礼品卡货架id",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/maintain/set?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "",
			},
			{
				Name:        "申请微信支付礼品卡权限",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/pay/whitelist/add?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "GiftCardPayWhitelistAdd",
			},
			{
				Name:        " 绑定商户号到礼品卡小程序",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/pay/submch/bind?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "GiftCardPaySubmchBind",
			},
			{
				Name:        "上传小程序代码",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/wxa/set",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "GiftCardWxaSet",
			},
			{
				Name:        "查询-单个礼品卡订单信息",
				Description: "开发者可以通过该接口查询某个订单号对应的订单详情",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/order/get?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "GiftCardOrderGet",
			},
			{
				Name:        "批量查询礼品卡订单信息",
				Description: "开发者可以通过该接口查询该商户某个时间段内创建的所有订单详情",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/order/batchget?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "GiftCardOrderBatchGet",
			},
			{
				Name:        "更新用户礼品卡信息",
				Description: "当礼品卡被使用后，开发者可以通过该接口变更某个礼品卡的余额信息",
				Request:     "POST https://api.weixin.qq.com/card/generalcard/updateuser?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "GeneralCardUpdateUser",
			},
			{
				Name:        "退款",
				Description: "开发者可以通过该接口对某一笔订单操作退款，注意该接口比较隐私，请开发者提高操作该功能的权限等级",
				Request:     "POST https://api.weixin.qq.com/card/giftcard/order/refund?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "GiftCardOrderRefund",
			},
			{
				Name:        "设置支付后开票信息",
				Description: "商户可以通过该接口设置某个商户号发生收款后在支付消息上出现开票授权按钮",
				Request:     "POST https://api.weixin.qq.com/card/invoice/setbizattr?action=set_pay_mch&access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "InvoiceSetBizAttr",
			},
			{
				Name:        "查询开票信息",
				Description: "用户完成授权后，商户可以调用该接口查询某一个订单",
				Request:     "POST https://api.weixin.qq.com/card/invoice/getauthdata",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/gift_card.html",
				FuncName:    "InvoiceGetAuthData",
			},
		},
	},
	{
		Name:    `会员卡专区`,
		Package: `card/membercard`,
		Apis: []Api{

			{
				Name:        "获取开卡插件参数",
				Description: "开发者可以通过该接口获取到调用开卡插件所需的参数",
				Request:     "POST https://api.weixin.qq.com/card/membercard/activate/geturl?access_token=",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Coupons-Mini_Program_Start_Up.html",
				FuncName:    "MembercardActivateGetUrl",
			},
			{
				Name:        "获取用户开卡时提交的信息（跳转型开卡组件）",
				Description: "开发者可以通过该接口获取到用户开卡时填写的字段值",
				Request:     "POST https://api.weixin.qq.com/card/membercard/activatetempinfo/get?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Coupons-Mini_Program_Start_Up.html",
				FuncName:    "MembercardActivateTempInfoGet",
			},
			{
				Name:        "激活用户领取的会员卡（跳转型开卡组件）",
				Description: "开发者可以通过该接口获取到用户开卡时填写的字段值",
				Request:     "POST https://api.weixin.qq.com/card/membercard/activate?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Coupons-Mini_Program_Start_Up.html",
				FuncName:    "MembercardActivate",
			},

			{
				Name:        "设置开卡字段",
				Description: "开发者在创建时填入wx_activate字段后，需要调用该接口设置用户激活时需要填写的选项，否则一键开卡设置不生效",
				Request:     "POST https://api.weixin.qq.com/card/membercard/activateuserform/set?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Create_a_membership_card.html",
				FuncName:    "MembercardActivateUserFormSet",
			},
			{
				Name:        "拉取会员信息",
				Description: "支持开发者根据CardID和Code查询会员信息",
				Request:     "POST https://api.weixin.qq.com/card/membercard/userinfo/get?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Create_a_membership_card.html",
				FuncName:    "MembercardUserinfoGet",
			},
			{
				Name:        "更新会员信息",
				Description: "当会员持卡消费后，支持开发者调用该接口更新会员信息。会员卡交易后的每次信息变更需通过该接口通知微信，便于后续消息通知及其他扩展功能",
				Request:     "POST https://api.weixin.qq.com/card/membercard/updateuser?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Create_a_membership_card.html",
				FuncName:    "MembercardUpdateUser",
			},

			{
				Name:        "设置支付后投放卡券",
				Description: "支持商户设置支付后投放卡券规则，可以区分时间段和金额区间发会员卡",
				Request:     "POST https://api.weixin.qq.com/card/paygiftcard/add?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Manage_Member_Card.html",
				FuncName:    "PayGiftcardAdd",
			},
			{
				Name:        "删除支付后投放卡券规则",
				Description: "支持商户删除之前设置的规则id",
				Request:     "POST https://api.weixin.qq.com/card/paygiftcard/delete?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Manage_Member_Card.html",
				FuncName:    "PayGiftcardDelete",
			},
			{
				Name:        "查询支付后投放卡券规则详情",
				Description: "可以查询某个支付即会员规则内容",
				Request:     "POST https://api.weixin.qq.com/card/paygiftcard/getbyid?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Manage_Member_Card.html",
				FuncName:    "PayGiftcardGetById",
			},
			{
				Name:        "批量查询支付后投放卡券规则",
				Description: "可以批量查询某个商户支付即会员规则内容",
				Request:     "POST https://api.weixin.qq.com/card/paygiftcard/batchget?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Membership_Cards/Manage_Member_Card.html",
				FuncName:    "PayGiftcardBatchget",
			},
		},
	},
	{
		Name:    `特殊票券`,
		Package: `card/ticket`,
		Apis: []Api{

			{
				Name:        "更新会议门票",
				Description: "持调用“更新会议门票”接口update 入场时间、区域、座位等信息",
				Request:     "POST https://api.weixin.qq.com/card/meetingticket/updateuser?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Special_ticket.html",
				FuncName:    "MeetingTicketUpdateUser",
			},
			{
				Name:        "更新电影票",
				Description: "领取电影票后通过调用“更新电影票”接口update电影信息及用户选座信息",
				Request:     "POST https://api.weixin.qq.com/card/movieticket/updateuser?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Special_ticket.html",
				FuncName:    "MovieTicketUpdateUser",
			},
			{
				Name:        "更新飞机票信息",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/boardingpass/checkin?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Special_ticket.html",
				FuncName:    "BoardingPassCheckin",
			},
		},
	},
	{
		Name:    `票券-第三方开发者模式`,
		Package: `card/submerchant`,
		Apis: []Api{

			{
				Name:        "创建子商户",
				Description: "支持母商户调用该接口传入子商户的相关资料，并获取子商户ID，用于子商户的卡券功能管理。 子商户的资质包括：商户名称、商户logo（图片）、卡券类目、授权函（扫描件或彩照）、授权函有效期截止时间",
				Request:     "POST https://api.weixin.qq.com/card/submerchant/submit?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html",
				FuncName:    "SubmerchantSubmit",
			},
			{
				Name:        "卡券开放类目查询",
				Description: "通过调用该接口查询卡券开放的类目ID，类目会随业务发展变更，请每次用接口去查询获取实时卡券类目",
				Request:     "POST https://api.weixin.qq.com/card/getapplyprotocol?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html",
				FuncName:    "GetApplyProtocol",
			},
			{
				Name:        "更新子商户",
				Description: "支持调用该接口更新子商户信息",
				Request:     "POST https://api.weixin.qq.com/card/submerchant/update?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html",
				FuncName:    "SubmerchantUpdate",
			},
			{
				Name:        "拉取单个子商户信息",
				Description: "通过指定的子商户appid，拉取该子商户的基础信息。 注意，用母商户去调用接口，但接口内传入的是子商户的appid",
				Request:     "POST https://api.weixin.qq.com/card/submerchant/get?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html",
				FuncName:    "SubmerchantGet",
			},
			{
				Name:        "批量拉取子商户信息",
				Description: "母商户可以通过该接口批量拉取子商户的相关信息，一次调用最多拉取100个子商户的信息，可以通过多次拉去满足不同的查询需求",
				Request:     "POST https://api.weixin.qq.com/card/submerchant/batchget?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Third-party_developer_mode.html",
				FuncName:    "SubmerchantBatchget",
			},
		},
	},
	{
		Name:    `微信门店`,
		Package: `poi`,
		Apis: []Api{

			{
				Name:        "创建门店",
				Description: "创建门店接口是为商户提供创建自己门店数据的接口，门店数据字段越完整，商户页面展示越丰富，越能够吸引更多用户，并提高曝光度",
				Request:     "POST http://api.weixin.qq.com/cgi-bin/poi/addpoi?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "查询门店信息",
				Description: "创建门店后获取poi_id 后，商户可以利用poi_id，查询具体某条门店的信息。 若在查询时，update_status 字段为1，表明在5 个工作日内曾用update 接口修改过门店扩展字段，该扩展字段为最新的修改字段，尚未经过审核采纳，因此不是最终结果。最终结果会在5 个工作日内，最终确认是否采纳，并前端生效（但该扩展字段的采纳过程不影响门店的可用性，即available_state仍为审核通过状态）",
				Request:     "POST http://api.weixin.qq.com/cgi-bin/poi/getpoi?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "查询门店列表",
				Description: "商户可以通过该接口，批量查询自己名下的门店list，并获取已审核通过的poiid、商户自身sid 用于对应、商户名、分店名、地址字段",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/poi/getpoilist?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html",
				FuncName:    "GetPoiList",
			},
			{
				Name:        "修改门店服务信息",
				Description: "商户可以通过该接口，修改门店的服务信息，包括：sid、图片列表、营业时间、推荐、特色服务、简介、人均价格、电话8个字段（名称、坐标、地址等不可修改）修改后需要人工审核",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/poi/updatepoi?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "删除门店",
				Description: "商户可以通过该接口，删除已经成功创建的门店。请商户慎重调用该接口",
				Request:     "POST https://api.weixin.qq.com/cgi-bin/poi/delpoi?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "门店类目表",
				Description: "类目名称接口是为商户提供自己门店类型信息的接口。门店类目定位的越规范，能够精准的吸引更多用户，提高曝光率",
				Request:     "POST http://api.weixin.qq.com/cgi-bin/poi/getwxcategory?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html",
				FuncName:    "GetWXCategory",
			},
		},
	},
	{
		Name:    `微信门店小程序`,
		Package: `poi/miniprogram`,
		Apis: []Api{
			{
				Name:        "拉取门店小程序类目",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/wxa/get_merchant_category?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "创建门店小程序",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/apply_merchant?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "查询门店小程序审核结果",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/wxa/get_merchant_audit_info?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "修改门店小程序信息",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/modify_merchant?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "从腾讯地图拉取省市区信息",
				Description: "",
				Request:     "GET https://api.weixin.qq.com/wxa/get_district?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "在腾讯地图中搜索门店",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/search_map_poi?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "在腾讯地图中创建门店",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/create_map_poi?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "添加门店",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/add_store?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "更新门店信息",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/update_store?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "获取门店小程序配置的卡券",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/storewxa/get?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "CardStorewxaGet",
			},
			{
				Name:        "获取单个门店信息",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/get_store_info?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "获取门店信息列表",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/wxa/get_store_list?access_token=TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Shop_Miniprogram_Interface.html",
				FuncName:    "",
			},
			{
				Name:        "删除门店",
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
				Name:        "语义理解",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/semantic/semproxy/search?access_token=YOUR_ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Natural_Language_Processing.html",
				FuncName:    "Semantic",
			},

			{
				Name:        "提交语音",
				Description: "",
				Request:     "POST http://api.weixin.qq.com/cgi-bin/media/voice/addvoicetorecofortext?access_token=ACCESS_TOKEN&format=&voice_id=xxxxxx&lang=zh_CN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/AI_Open_API.html",
				FuncName:    "AddVoiceToRecoForText",
			},
			{
				Name:        "获取语音识别结果",
				Description: "",
				Request:     "POST http://api.weixin.qq.com/cgi-bin/media/voice/queryrecoresultfortext?access_token=ACCESS_TOKEN&voice_id=xxxxxx&lang=zh_CN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/AI_Open_API.html",
				FuncName:    "QueryRecoResultForText",
			},
			{
				Name:        "微信翻译",
				Description: "",
				Request:     "POST http://api.weixin.qq.com/cgi-bin/media/voice/translatecontent?access_token=ACCESS_TOKEN&lfrom=xxx&lto=xxx",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/AI_Open_API.html",
				FuncName:    "TranslateContent",
			},

			{
				Name:        "身份证OCR识别",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cv/ocr/idcard?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html",
				FuncName:    "IDCard",
			},
			{
				Name:        "银行卡OCR识别",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cv/ocr/bankcard",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html",
				FuncName:    "",
			},
			{
				Name:        "行驶证/驾驶证 OCR识别",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cv/ocr/drivinglicense?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html",
				FuncName:    "DrivingLicense",
			},
			{
				Name:        "营业执照OCR识别",
				Description: "",
				Request:     "POST http://api.weixin.qq.com/cv/ocr/bizlicense?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html",
				FuncName:    "BizLicense",
			},
			{
				Name:        "通用印刷体OCR识别",
				Description: "",
				Request:     "POST http://api.weixin.qq.com/cv/ocr/comm?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html",
				FuncName:    "Comm",
			},

			{
				Name:        "二维码/条码识别",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cv/img/qrcode?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Img_Proc.html",
				FuncName:    "QRCode",
			},
			{
				Name:        "图片高清化",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/cv/img/superresolution?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Img_Proc.html",
				FuncName:    "SuperResolution",
			},
			{
				Name:        "图片智能裁剪",
				Description: "",
				Request:     "POST http://api.weixin.qq.com/cv/img/aicrop?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Img_Proc.html",
				FuncName:    "AICrop",
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
		Package: `guide/tag`,
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
		Package: `onethingonecode`,
		Apis: []Api{

			{
				Name:        "申请二维码",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/intp/marketcode/applycode?access_token=ACCESSTOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html",
				FuncName:    "ApplyCode",
			},
			{
				Name:        "查询二维码申请单",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/intp/marketcode/applycodequery?access_token=ACCESSTOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html",
				FuncName:    "ApplyCodeQuery",
			},
			{
				Name:        "下载二维码包",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/intp/marketcode/applycodedownload?access_token=ACCESSTOKE",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html",
				FuncName:    "ApplyCodeDownload",
			},
			{
				Name:        "激活二维码",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/intp/marketcode/codeactive?access_token=ACCESSTOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html",
				FuncName:    "CodeActive",
			},
			{
				Name:        "查询二维码激活状态",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/intp/marketcode/codeactivequery?access_token=ACCESSTOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html",
				FuncName:    "CodeActiveQuery",
			},
			{
				Name:        "code_ticket换code",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/intp/marketcode/tickettocode?access_token=ACCESSTOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Unique_Item_Code/Unique_Item_Code_API_Documentation.html",
				FuncName:    "TicketToCode",
			},
		},
	},
	{
		Name:    `微信发票`,
		Package: `invoice`,
		Apis: []Api{

			{
				Name:        "获取授权页链接",
				Description: "本接口供商户调用。商户通过本接口传入订单号、开票平台标识等参数，获取授权页的链接。在微信中向用户展示授权页，当用户点击了授权页上的“领取发票”/“申请开票”按钮后，即完成了订单号与该用户的授权关系绑定，后续开票平台可凭此订单号发起将发票卡券插入用户卡包的请求，微信也将据此授权关系校验是否放行插卡请求",
				Request:     "POST https://api.weixin.qq.com/card/invoice/getauthurl?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html",
				FuncName:    "GetAuthUrl",
			},
			{
				Name:        "查询授权完成状态",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/getauthdata?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html",
				FuncName:    "GetAuthData",
			},
			{
				Name:        "拒绝开票",
				Description: "户完成授权后，商户若发现用户提交信息错误、或者发生了退款时，可以调用该接口拒绝开票并告知用户。拒绝开票后，该订单无法向用户再次开票。已经拒绝开票的订单，无法再次使用，如果要重新开票，需使用新的order_id，获取授权链接，让用户再次授权",
				Request:     "POST https://api.weixin.qq.com/card/invoice/rejectinsert?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html",
				FuncName:    "RejectInsert",
			},
			{
				Name:        "统一开票接口-开具蓝票",
				Description: "对于使用微信电子发票开票接入能力的商户，在公众号后台选择任何一家开票平台的套餐，都可以使用本接口实现电子发票的开具",
				Request:     "POST https://api.weixin.qq.com/card/invoice/makeoutinvoice?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html",
				FuncName:    "",
			},
			{
				Name:        "统一开票接口-发票冲红",
				Description: "对于使用微信电子发票开票接入能力的商户，在公众号后台选择任何一家开票平台的套餐，都可以使用本接口实现电子发票的冲红",
				Request:     "POST https://api.weixin.qq.com/card/invoice/clearoutinvoice?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html",
				FuncName:    "ClearOutInvoice",
			},
			{
				Name:        "统一开票接口-查询已开发票",
				Description: "对于使用微信电子发票开票接入能力的商户，在公众号后台选择任何一家开票平台的套餐，都可以使用本接口实现已开具电子发票的查询",
				Request:     "POST https://api.weixin.qq.com/card/invoice/queryinvoceinfo?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Vendor_API_List.html",
				FuncName:    "QueryInvoceInfo",
			},

			{
				Name:        "获取自身的开票平台识别码",
				Description: "开票平台可以通过此接口获得本开票平台的预开票url，进而获取s_pappid。开票平台将该s_pappid并透传给商户，商户可以通过该s_pappid参数在微信电子发票方案中标识出为自身提供开票服务的开票平台",
				Request:     "POST https://api.weixin.qq.com/card/invoice/seturl?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html",
				FuncName:    "SetUrl",
			},
			{
				Name:        "创建发票卡券模板",
				Description: "通过本接口可以为创建一个商户的发票卡券模板，为该商户配置发票卡券模板上的自定义栏位。创建发票卡券模板生成的card_id将在创建发票卡券时被引用，故创建发票卡券模板是创建发票卡券的基础",
				Request:     "POST https://api.weixin.qq.com/card/invoice/platform/createcard?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html",
				FuncName:    "PlatformCreateCard",
			},
			{
				Name:        "上传PDF",
				Description: "商户或开票平台可以通过该接口上传PDF。PDF上传成功后将获得发票文件的标识，后续可以通过插卡接口将PDF关联到用户的发票卡券上，一并插入到收票用户的卡包中",
				Request:     "POST https://api.weixin.qq.com/card/invoice/platform/setpdf?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html",
				FuncName:    "PlatformSetpdf",
			},
			{
				Name:        "查询已上传的PDF文件",
				Description: "用于供发票PDF的上传方查询已经上传的发票或消费凭证PDF",
				Request:     "POST https://api.weixin.qq.com/card/invoice/platform/getpdf?action=get_url&access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html",
				FuncName:    "PlatformGetpdf",
			},
			{
				Name:        "将电子发票卡券插入用户卡包",
				Description: "本接口由开票平台或自建平台商户调用。对用户已经授权过的开票请求，开票平台可以使用本接口将发票制成发票卡券放入用户的微信卡包中",
				Request:     "POST https://api.weixin.qq.com/card/invoice/insert?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html",
				FuncName:    "",
			},
			{
				Name:        "更新发票卡券状态",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/platform/updatestatus?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Invoicing_Platform_API_List.html",
				FuncName:    "PlatformUpdateStatus",
			},

			{
				Name:        "查询报销发票信息",
				Description: "通过该接口查询电子发票的结构化信息，并获取发票PDF文件",
				Request:     "POST https://api.weixin.qq.com/card/invoice/reimburse/getinvoiceinfo?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Reimburser_API_List.html",
				FuncName:    "ReimburseGetInvoiceInfo",
			},
			{
				Name:        "批量查询报销发票信息",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/reimburse/getinvoicebatch?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Reimburser_API_List.html",
				FuncName:    "ReimburseGetInvoiceBatch",
			},
			{
				Name:        "报销方更新发票状态",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/reimburse/updateinvoicestatus?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Reimburser_API_List.html",
				FuncName:    "ReimburseUpdateInvoiceStatus",
			},
			{
				Name:        "报销方批量更新发票状态",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/card/invoice/reimburse/updatestatusbatch?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/E_Invoice/Reimburser_API_List.html",
				FuncName:    "ReimburseUpdateStatusBatch",
			},

			{
				Name:        "将发票抬头信息录入到用户微信中",
				Description: "调用接口，获取添加存储发票抬头信息的链接，将链接发给微信用户，用户确认后将保存该信息",
				Request:     "POST https://api.weixin.qq.com/card/invoice/biz/getusertitleurl?access_token={access_token",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Quick_issuing/Interface_Instructions.html",
				FuncName:    "GetUserTitleUrl",
			},
			{
				Name:        "获取用户抬头（方式一）:获取商户专属二维码，立在收银台",
				Description: "商户调用接口，获取链接，将链接转成二维码，用户扫码，可以选择抬头发给商户",
				Request:     "POST https://api.weixin.qq.com/card/invoice/biz/getselecttitleurl?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Quick_issuing/Interface_Instructions.html",
				FuncName:    "GetSelectTitleUrl",
			},
			{
				Name:        "获取用户抬头（方式二）:商户扫描用户的发票抬头二维码",
				Description: "商户扫用户“我的—个人信息—我的发票抬头”里面的抬头二维码后，通过调用本接口，可以获取用户抬头信息",
				Request:     "POST https://api.weixin.qq.com/card/invoice/scantitle?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Quick_issuing/Interface_Instructions.html",
				FuncName:    "ScanTitle",
			},
		},
	},
	{
		Name:    `非税票据/缴费`,
		Package: `nontax`,
		Apis: []Api{

			{
				Name:        "获取授权页链接",
				Description: "通过此接口，获取授权页链接，让用户跳转到授权页",
				Request:     "POST https://api.weixin.qq.com/nontax/getbillauthurl?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Nontax_Bill/API_list.html",
				FuncName:    "GetBillAuthUrl",
			},
			{
				Name:        "创建财政电子票据接口",
				Description: "财政局可以通过这个接口帮助执收单位创建一张财政电子票据模板。同一个财政局可以对应多个执收单位，同一个执收单位，使用同一个card_id，不同的执收单位，使用不同的card_id",
				Request:     "POST https://api.weixin.qq.com/nontax/createbillcard?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Nontax_Bill/API_list.html",
				FuncName:    "CreateBillCard",
			},
			{
				Name:        "将财政电子票据添加到用户微信卡包",
				Description: "执收单位完成用户插卡授权后，向财政局请求给某一个订单号进行领取财政电子票据，财政局须调用该接口对用户进行开票",
				Request:     "POST https://api.weixin.qq.com/nontax/insertbill?access_token={access_token}",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/WeChat_Invoice/Nontax_Bill/API_list.html",
				FuncName:    "InsertBill",
			},

			{
				Name:        "查询应收信息（提供给委办局）",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/queryfee?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "QueryFee",
			},
			{
				Name:        "支付下单（提供给委办局）",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/unifiedorder?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "UnifiedOrder",
			},
			{
				Name:        "查询订单（提供给委办局、银行、财政）",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/getorder?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "",
			},
			{
				Name:        "申请退款（提供给银行）",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/refund?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "",
			},
			{
				Name:        "下载对帐单（提供给银行）",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/downloadbill?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "DownloadBill",
			},
			{
				Name:        "通知不一致订单（提供给财政）",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/notifyinconsistentorder?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "NotifyInconsistentOrder",
			},
			{
				Name:        "测试支付结果通知",
				Description: "接口提供给接入方在联调前自己调试支付通知接口的加解密和其它基本逻辑。 若version为1，会通知两次，分别测试加解密是否正确和是否有验签。 调用此接口会触发微信后台向 url 发送支付结果通知的测试数据",
				Request:     "POST https://api.weixin.qq.com/nontax/mocknotification?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "MockNotification",
			},
			{
				Name:        "测试查询应收信息",
				Description: "此接口提供给接入方在联调前自己调试查询应收信息接口的加解密和其它基本逻辑。 若version为1，会通知两次，分别测试加解密是否正确和是否有验签。 调用此接口会触发微信后台向 url 发送查询应收信息的测试数据",
				Request:     "POST https://api.weixin.qq.com/nontax/mockqueryfee?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "MockQueryFee",
			},
			{
				Name:        "提交刷卡支付",
				Description: "提交支付请求后微信会同步返回支付结果。 接口返回系统失败时，等待5秒重新调用看返回码。 当结果返回用户支付中需要输入密码时，可每间隔一段时间(建议10秒)重新调用该接口，直到有明确成功、失败，或者超时（建议30秒）",
				Request:     "POST https://api.weixin.qq.com/nontax/micropay?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "",
			},
			{
				Name:        "查询订单列表",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/getorderlist?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "GetOrderList",
			},
			{
				Name:        "获取用户实名信息-获取授权链接",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/intp/realname/getauthurl?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "RealNameGetAuthUrl",
			},
			{
				Name:        "获取实名信息",
				Description: "",
				Request:     "POST https://api.weixin.qq.com/nontax/getrealname?access_token=ACCESS_TOKEN",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/API_document.html",
				FuncName:    "GetRealName",
			},
		},
	},
	{
		Name:    `非税车主平台`,
		Package: `nontax/vehicle`,
		Apis: []Api{
			{
				Name:        "用户状态查询",
				Description: "在停车场、高速、加油等场景下，商户需获取用户车主服务状态/需要关联车主服务。本接口，会查询用户是否开通、授权、有欠费或黑名单用户情况，并将对应的用户状态进行返回",
				Request:     "POST https://api.weixin.qq.com/nontax/vehicle/querystate?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/Non-tax_driver_platform.html",
				FuncName:    "",
			},
			{
				Name:        "用户入场通知",
				Description: "在停车场场景下，如用户已加入车主平台，则进行入场通知；如用户已经欠费，会发送用户欠费入场通知。本接口，会查询用户是否有欠费或黑名单用户情况，并将对应的用户状态进行返回",
				Request:     "POST https://api.weixin.qq.com/nontax/vehicle/entrancenotify?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/Non-tax_driver_platform.html",
				FuncName:    "EntranceNotify",
			},
			{
				Name:        "申请扣款",
				Description: "委托代扣可应用于定期扣款或需事后扣款以期提高效率的场景。例如高速，停车场等通过用户授权给商户，进行委托扣款的场景",
				Request:     "POST https://api.weixin.qq.com/nontax/vehicle/payapply?access_token=$AccessToken",
				See:         "https://developers.weixin.qq.com/doc/offiaccount/Non_tax_pay/Non-tax_driver_platform.html",
				FuncName:    "PayApply",
			},
		},
	},
}
