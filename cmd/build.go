package main

import (
	"flag"
	"fmt"
	"net/url"
	"path"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
)

func main() {

	var pkgFlag string
	flag.StringVar(&pkgFlag, "package", "default", "")
	flag.Parse()

	for _, group := range apiConfig {
		if group.Package == pkgFlag {
			build(group)
		}
	}
}

func build(group ApiGroup) {
	var funcs []string
	var consts []string

	for _, api := range group.Apis {
		tpl := postFuncTpl
		_FUNC_NAME_ := ""
		_UPLOAD_ := "media"
		_FIELD_NAME_ := ""
		_FIELDS_ := ""
		_PAYLOAD_ := ""
		switch {
		case strings.Contains(api.Request, "GET http"):
			tpl = getFuncTpl
		case strings.Contains(api.Request, "POST http"):
			tpl = postFuncTpl
		case strings.Contains(api.Request, "POST(@media"):
			tpl = postUploadFuncTpl
			_UPLOAD_ = "media"

			pattern := `POST\(@media\|field=(\S+)\) http`
			reg := regexp.MustCompile(pattern)
			matched := reg.FindAllStringSubmatch(api.Request, -1)
			if matched != nil {
				_FIELD_NAME_ = matched[0][1]
				_PAYLOAD_ = ", payload []byte"
			}
		}

		split := strings.Split(api.Request, " ")
		parseUrl, _ := url.Parse(split[1])

		_TITLE_ := api.Name
		_DESCRIPTION_ := api.Description
		_REQUEST_ := api.Request
		_SEE_ := api.See
		if api.FuncName == "" {
			_FUNC_NAME_ = strcase.ToCamel(path.Base(parseUrl.Path))
		} else {
			_FUNC_NAME_ = api.FuncName
		}

		tpl = strings.ReplaceAll(tpl, "_TITLE_", _TITLE_)
		tpl = strings.ReplaceAll(tpl, "_DESCRIPTION_", _DESCRIPTION_)
		tpl = strings.ReplaceAll(tpl, "_REQUEST_", _REQUEST_)
		tpl = strings.ReplaceAll(tpl, "_SEE_", _SEE_)
		tpl = strings.ReplaceAll(tpl, "_FUNC_NAME_", _FUNC_NAME_)
		tpl = strings.ReplaceAll(tpl, "_UPLOAD_", _UPLOAD_)
		if _FIELD_NAME_ != "" {
			_FIELDS_ = strings.ReplaceAll(fieldTpl, "_FIELD_NAME_", _FIELD_NAME_)
		}
		tpl = strings.ReplaceAll(tpl, "_FIELDS_", _FIELDS_)
		tpl = strings.ReplaceAll(tpl, "_PAYLOAD_", _PAYLOAD_)

		funcs = append(funcs, tpl)

		tpl = strings.ReplaceAll(constTpl, "_FUNC_NAME_", _FUNC_NAME_)
		tpl = strings.ReplaceAll(tpl, "_API_PATH_", parseUrl.Path)
		consts = append(consts, tpl)
	}

	fmt.Printf(fileTpl, group.Package, strings.Join(consts, ``), strings.Join(funcs, ``))
}

var constTpl = `
	api_FUNC_NAME_ = "_API_PATH_"`
var commentTpl = `
/*
_TITLE_

_DESCRIPTION_

See: _SEE_

_REQUEST_
*/`
var postFuncTpl = commentTpl + `
func _FUNC_NAME_(payload []byte) (resp []byte, err error) {
	return offiaccount.HTTPPost(api_FUNC_NAME_, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
`
var getFuncTpl = commentTpl + `
func _FUNC_NAME_() (resp []byte, err error) {
	return offiaccount.HTTPGet(api_FUNC_NAME_)
}
`
var postUploadFuncTpl = commentTpl + `
func _FUNC_NAME_(_UPLOAD_ string_PAYLOAD_) (resp []byte, err error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		part, err := m.CreateFormFile("_UPLOAD_", path.Base(_UPLOAD_))
		if err != nil {
			return
		}
		file, err := os.Open(_UPLOAD_)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}

		_FIELDS_
	}()
	return offiaccount.HTTPPost(api_FUNC_NAME_, r, m.FormDataContentType())
}
`

var fieldTpl = `
		// field
		err = m.WriteField("_FIELD_NAME_", string(payload))
		if err != nil {
			return
		}
`

var fileTpl = `package %s
const (
	%s
)
%s`
