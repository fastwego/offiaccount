package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
)

const ServerUrl = `https://developers.weixin.qq.com`

var apiUniqMap = map[string]bool{}

func main() {

	file, err := ioutil.ReadFile("./data/doc_links.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	pattern := `href="(/doc/offiaccount/.+\.html)"`
	reg := regexp.MustCompile(pattern)
	matched := reg.FindAllStringSubmatch(string(file), -1)

	for _, match := range matched {
		//fmt.Println(match[1])

		link := ServerUrl + match[1]
		resp, err := http.Get(link)
		if err != nil {
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			continue
		}

		apiPattern := `http[s]*://api.weixin.qq.com/\S+`
		reg := regexp.MustCompile(apiPattern)
		apiMatched := reg.FindAllStringSubmatch(string(body), -1)
		if apiMatched == nil {
			continue
		}
		for _, apis := range apiMatched {
			//fmt.Println(apis[0])

			parse, _ := url.Parse(apis[0])
			if _, ok := apiUniqMap[parse.Path]; !ok {
				apiUniqMap[parse.Path] = true
			} else {
				continue
			}
			_NAME_ := parse.Path
			_DESCRIPTION_ := ""

			indexByte := strings.IndexByte(apis[0], '<')
			if indexByte > 0 {
				apis[0] = apis[0][0:indexByte]
			}
			_REQUEST_ := "POST " + apis[0]
			_SEE_ := link
			_FUNCNAME_ := ""

			if strings.Contains(apis[0], "/cgi-bin/guide/") {
				r := regexp.MustCompile(`shopping-guide\.(.*)\.html`)
				found := r.FindStringSubmatch(link)
				if found[1] != "" {
					_FUNCNAME_ = strcase.ToCamel(found[1])
				}

				r = regexp.MustCompile(`<h1 id="shopping-guide-.+</blockquote> <p>(\S+)</p>`)
				found = r.FindStringSubmatch(string(body))
				if len(found) > 0 && found[1] != "" {
					_NAME_ = found[1]
				}

			}

			tpl := strings.ReplaceAll(itemTpl, "_NAME_", _NAME_)
			tpl = strings.ReplaceAll(tpl, "_DESCRIPTION_", _DESCRIPTION_)
			tpl = strings.ReplaceAll(tpl, "_REQUEST_", _REQUEST_)
			tpl = strings.ReplaceAll(tpl, "_SEE_", _SEE_)
			tpl = strings.ReplaceAll(tpl, "_FUNCNAME_", _FUNCNAME_)

			fmt.Println(tpl)
		}

		fmt.Println(`


`)

		//break
	}

}

var itemTpl = `{
	Name:        "_NAME_",
	Description: "_DESCRIPTION_",
	Request:     "_REQUEST_",
	See:         "_SEE_",
	FuncName:    "_FUNCNAME_",
},`
