package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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
	var testFuncs []string

	for _, api := range group.Apis {
		tpl := postFuncTpl
		_FUNC_NAME_ := ""
		_GET_PARAMS_ := ""
		_GET_SUFFIX_PARAMS_ := ""
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
		if len(api.GetParams) > 0 {
			_GET_PARAMS_ = `params url.Values`
			if strings.Contains(api.Request, "POST(@media") {
				_GET_PARAMS_ = `, ` + _GET_PARAMS_
			}
			_GET_SUFFIX_PARAMS_ = `+ "?" + params.Encode()`
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
		tpl = strings.ReplaceAll(tpl, "_GET_PARAMS_", _GET_PARAMS_)
		tpl = strings.ReplaceAll(tpl, "_GET_SUFFIX_PARAMS_", _GET_SUFFIX_PARAMS_)
		if _FIELD_NAME_ != "" {
			_FIELDS_ = strings.ReplaceAll(fieldTpl, "_FIELD_NAME_", _FIELD_NAME_)
		}
		tpl = strings.ReplaceAll(tpl, "_FIELDS_", _FIELDS_)
		tpl = strings.ReplaceAll(tpl, "_PAYLOAD_", _PAYLOAD_)

		funcs = append(funcs, tpl)

		tpl = strings.ReplaceAll(constTpl, "_FUNC_NAME_", _FUNC_NAME_)
		tpl = strings.ReplaceAll(tpl, "_API_PATH_", parseUrl.Path)

		consts = append(consts, tpl)

		// TestFunc
		_TEST_ARGS_STRUCT_ := ""
		switch {
		case strings.Contains(api.Request, "GET http"):
			_TEST_ARGS_STRUCT_ = _GET_PARAMS_
		case strings.Contains(api.Request, "POST http"):
			_TEST_ARGS_STRUCT_ = `payload []byte` + _GET_PARAMS_
		case strings.Contains(api.Request, "POST(@media"):
			_TEST_ARGS_STRUCT_ = _UPLOAD_ + ` string` + _PAYLOAD_ + _GET_PARAMS_
		}
		_TEST_ARGS_STRUCT_ = strings.ReplaceAll(_TEST_ARGS_STRUCT_, ",", "\n")

		_TEST_FUNC_SIGNATURE_ := ""
		if strings.TrimSpace(_TEST_ARGS_STRUCT_) != "" {
			signatures := strings.Split(_TEST_ARGS_STRUCT_, "\n")
			paramNames := []string{}
			for _, signature := range signatures {
				signature = strings.TrimSpace(signature)
				tmp := strings.Split(signature, " ")
				paramNames = append(paramNames, "tt.args."+tmp[0])
			}
			_TEST_FUNC_SIGNATURE_ = strings.Join(paramNames, ",")
		}

		tpl = strings.ReplaceAll(testFuncTpl, "_FUNC_NAME_", _FUNC_NAME_)
		tpl = strings.ReplaceAll(tpl, "_TEST_ARGS_STRUCT_", _TEST_ARGS_STRUCT_)
		tpl = strings.ReplaceAll(tpl, "_TEST_FUNC_SIGNATURE_", _TEST_FUNC_SIGNATURE_)
		testFuncs = append(testFuncs, tpl)
	}

	fmt.Printf(fileTpl, path.Base(group.Package), strings.Join(consts, ``), strings.Join(funcs, ``))

	// output Test
	testFileContent := fmt.Sprintf(testFileTpl, path.Base(group.Package), strings.Join(testFuncs, ``))
	//fmt.Println(testFileContent)
	ioutil.WriteFile("../"+group.Package+"/"+path.Base(group.Package)+"_test.go", []byte(testFileContent), 0644)
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
func _FUNC_NAME_(payload []byte_GET_PARAMS_) (resp []byte, err error) {
	return offiaccount.HTTPPost(api_FUNC_NAME__GET_SUFFIX_PARAMS_, bytes.NewBuffer(payload), offiaccount.ContentTypeApplicationJson)
}
`
var getFuncTpl = commentTpl + `
func _FUNC_NAME_(_GET_PARAMS_) (resp []byte, err error) {
	return offiaccount.HTTPGet(api_FUNC_NAME__GET_SUFFIX_PARAMS_)
}
`
var postUploadFuncTpl = commentTpl + `
func _FUNC_NAME_(_UPLOAD_ string_PAYLOAD__GET_PARAMS_) (resp []byte, err error) {
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
	return offiaccount.HTTPPost(api_FUNC_NAME__GET_SUFFIX_PARAMS_, r, m.FormDataContentType())
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

var testFileTpl = `package %s

func TestMain(m *testing.M) {
	test.Setup()
	os.Exit(m.Run())
}

%s
`

var testFuncTpl = `
func Test_FUNC_NAME_(t *testing.T) {
	mockResp := map[string][]byte{
		"case1": []byte("{\"errcode\":0,\"errmsg\":\"ok\"}"),
	}
	var resp []byte
	test.MockSvrHandler.HandleFunc(api_FUNC_NAME_, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	})

	type args struct {
		_TEST_ARGS_STRUCT_
	}
	tests := []struct {
		name     string
		args     args
		wantResp []byte
		wantErr  bool
	}{
		{name: "case1", args: args{}, wantResp: mockResp["case1"], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp = mockResp[tt.name]
			gotResp, err := _FUNC_NAME_(_TEST_FUNC_SIGNATURE_)
			//fmt.Println(string(gotResp), err)
			if (err != nil) != tt.wantErr {
				t.Errorf("_FUNC_NAME_() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("_FUNC_NAME_() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}`
