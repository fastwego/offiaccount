package material

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path"

	"github.com/fastwego/offiaccount"
)

const (
	apiMediaUpload      = "/cgi-bin/media/upload"
	apiMediaGet         = "/cgi-bin/media/get"
	apiMediaGetJssdk    = "/cgi-bin/media/get/jssdk"
	apiAddNews          = "/cgi-bin/material/add_news"
	apiMediaUploadImg   = "/cgi-bin/media/uploadimg"
	apiAddMaterial      = "/cgi-bin/material/add_material"
	apiGetMaterial      = "/cgi-bin/material/get_material"
	apiDelMaterial      = "/cgi-bin/material/del_material"
	apiUpdateNews       = "/cgi-bin/material/update_news"
	apiGetMaterialCount = "/cgi-bin/material/get_materialcount"
	apiBatchgetMaterial = "/cgi-bin/material/batchget_material"
)

/*
新增临时素材

公众号经常有需要用到一些临时性的多媒体素材的场景，例如在使用接口特别是发送消息时，对多媒体文件、多媒体消息的获取和调用等操作，是通过media_id来进行的。素材管理接口对所有认证的订阅号和服务号开放。通过本接口，公众号可以新增临时素材（即上传临时多媒体文件）

See: https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html

POST(@media) https://api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE
*/
func MediaUpload(media string) (resp []byte, err error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("media", path.Base(media))
		if err != nil {
			return
		}
		file, err := os.Open(media)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}

	}()
	return offiaccount.HTTPPost(apiMediaUpload, r, m.FormDataContentType())
}

/*
获取临时素材

公众号可以使用本接口获取临时素材（即下载临时的多媒体文件）

See: https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_temporary_materials.html

GET https://api.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id=MEDIA_ID
*/
func MediaGet() (resp []byte, err error) {
	return offiaccount.HTTPGet(apiMediaGet)
}

/*
高清语音素材获取接口

公众号可以使用本接口获取从JSSDK的uploadVoice接口上传的临时语音素材，格式为speex，16K采样率。该音频比上文的临时素材获取接口（格式为amr，8K采样率）更加清晰，适合用作语音识别等对音质要求较高的业务

See: https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_temporary_materials.html

GET https://api.weixin.qq.com/cgi-bin/media/get/jssdk?access_token=ACCESS_TOKEN&media_id=MEDIA_ID
*/
func MediaGetJssdk(params url.Values) (resp []byte, err error) {
	return offiaccount.HTTPGet(apiMediaGetJssdk + "?" + params.Encode())
}

/*
新增永久图文素材

对于常用的素材，开发者可通过本接口上传到微信服务器，永久使用。新增的永久素材也可以在公众平台官网素材管理模块中查询管理

See: https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html

POST https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=ACCESS_TOKEN
*/
func AddNews(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiAddNews, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
上传图文消息内的图片获取URL

本接口所上传的图片不占用公众号的素材库中图片数量的100000个的限制。图片仅支持jpg/png格式，大小必须在1MB以下

See: https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html

POST(@media) https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN
*/
func MediaUploadImg(media string) (resp []byte, err error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("media", path.Base(media))
		if err != nil {
			return
		}
		file, err := os.Open(media)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}

	}()
	return offiaccount.HTTPPost(apiMediaUploadImg, r, m.FormDataContentType())
}

/*
新增其他类型永久素材

通过POST表单来调用接口，表单id为media，包含需要上传的素材内容，有filename、filelength、content-type等信息。请注意：图片素材将进入公众平台官网素材管理模块中的默认分组

See: https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html

POST(@media|field=description) https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=ACCESS_TOKEN&type=TYPE
*/
func AddMaterial(media string, payload []byte) (resp []byte, err error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("media", path.Base(media))
		if err != nil {
			return
		}
		file, err := os.Open(media)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}

		// field
		err = m.WriteField("description", string(payload))
		if err != nil {
			return
		}

	}()
	return offiaccount.HTTPPost(apiAddMaterial, r, m.FormDataContentType())
}

/*
获取永久素材

在新增了永久素材后，开发者可以根据media_id通过本接口下载永久素材。公众号在公众平台官网素材管理模块中新建的永久素材，可通过"获取素材列表"获知素材的media_id。请注意：临时素材无法通过本接口获取

See: https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Getting_Permanent_Assets.html

POST https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=ACCESS_TOKEN
*/
func GetMaterial(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiGetMaterial, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
删除永久素材

在新增了永久素材后，开发者可以根据本接口来删除不再需要的永久素材，节省空间

See: https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Deleting_Permanent_Assets.html

POST https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=ACCESS_TOKEN
*/
func DelMaterial(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiDelMaterial, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
修改永久图文素材

开发者可以通过本接口对永久图文素材进行修改

See: https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Editing_Permanent_Rich_Media_Assets.html

POST https://api.weixin.qq.com/cgi-bin/material/update_news?access_token=ACCESS_TOKEN
*/
func UpdateNews(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiUpdateNews, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}

/*
获取素材总数

1.永久素材的总数，也会计算公众平台官网素材管理中的素材 2.图片和图文消息素材（包括单图文和多图文）的总数上限为5000，其他素材的总数上限为1000

See: https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_the_total_of_all_materials.html

GET https://api.weixin.qq.com/cgi-bin/material/get_materialcount?access_token=ACCESS_TOKEN
*/
func GetMaterialCount() (resp []byte, err error) {
	return offiaccount.HTTPGet(apiGetMaterialCount)
}

/*
获取素材列表

在新增了永久素材后，开发者可以分类型获取永久素材的列表

See: https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_materials_list.html

POST https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=ACCESS_TOKEN
*/
func BatchgetMaterial(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(apiBatchgetMaterial, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
