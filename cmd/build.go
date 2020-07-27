package main

import (
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/iancoleman/strcase"
)

type API struct {
	request string
	see     string
}

func main() {
	fmt.Println("Start")

	apis := [][]API{
		{
			// 素材管理
			{
				"POST/FORM https://API.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE",
				"https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/New_temporary_materials.html",
			},
			{
				"GET https://API.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id=MEDIA_ID",
				"https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_temporary_materials.html",
			},
			{
				"POST https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=ACCESS_TOKEN",
				"https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html",
			},
		},
		{
			{
				"API",
				"see",
			},
			{
				"API",
				"see",
			},
		},
	}

	fmt.Println(apis)

	build(apis[0])

}

func build(apis []API) {
	var funcs []string
	var consts []string

	for _, api := range apis {
		tpl := postFunc
		switch {
		case strings.Contains(api.request, "GET http"):
			tpl = getFunc
		case strings.Contains(api.request, "POST http"):
			tpl = postFunc
		case strings.Contains(api.request, "POST/FORM http"):
			tpl = postFormFunc
		}

		split := strings.Split(api.request, " ")
		parse, _ := url.Parse(split[1])

		_FUNC_NAME_ := strcase.ToCamel(path.Base(parse.Path))
		_REQUEST_ := api.request
		_SEE_ := api.see

		funcCode := fmt.Sprintf(tpl, _FUNC_NAME_, _SEE_, _REQUEST_, _FUNC_NAME_, _FUNC_NAME_)
		funcs = append(funcs, funcCode)

		constCode := fmt.Sprintf(`api%s = "%s"`, _FUNC_NAME_, parse.Path)
		consts = append(consts, constCode)
	}

	fmt.Printf(`
const (
	%s
)

%s
`, strings.Join(consts, ``), strings.Join(funcs, ``))
}

var postFunc = `
/*
%s

See: %s

%s
*/
func %s(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(api%s, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
`
var getFunc = `
/*
%s

See: %s

%s
*/
func %s() (resp []byte, err error) {
	return offiaccount.HTTPGet(api%s)
}
`
var postFormFunc = `
/*
%s

See: %s

%s
*/
func %s(mediaPath string) (resp []byte, err error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("media", path.Base(mediaPath))
		if err != nil {
			return
		}
		file, err := os.Open(mediaPath)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}
	}()
	return offiaccount.HTTPPost(api%s, r, m.FormDataContentType())
}
`
