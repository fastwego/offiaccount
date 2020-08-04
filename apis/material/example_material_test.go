// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package material_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/material"
)

func ExampleMediaUpload() {
	var ctx *offiaccount.OffiAccount

	media := ""
	resp, err := material.MediaUpload(ctx, media)

	fmt.Println(resp, err)
}

func ExampleMediaGet() {
	var ctx *offiaccount.OffiAccount

	resp, err := material.MediaGet(ctx)

	fmt.Println(resp, err)
}

func ExampleMediaGetJssdk() {
	var ctx *offiaccount.OffiAccount

	params := url.Values{}
	resp, err := material.MediaGetJssdk(ctx, params)

	fmt.Println(resp, err)
}

func ExampleAddNews() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := material.AddNews(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMediaUploadImg() {
	var ctx *offiaccount.OffiAccount

	media := ""
	resp, err := material.MediaUploadImg(ctx, media)

	fmt.Println(resp, err)
}

func ExampleAddMaterial() {
	var ctx *offiaccount.OffiAccount

	media := ""
	payload := []byte("{}")
	resp, err := material.AddMaterial(ctx, media, payload)

	fmt.Println(resp, err)
}

func ExampleGetMaterial() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := material.GetMaterial(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelMaterial() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := material.DelMaterial(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdateNews() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := material.UpdateNews(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetMaterialCount() {
	var ctx *offiaccount.OffiAccount

	resp, err := material.GetMaterialCount(ctx)

	fmt.Println(resp, err)
}

func ExampleBatchgetMaterial() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := material.BatchgetMaterial(ctx, payload)

	fmt.Println(resp, err)
}
