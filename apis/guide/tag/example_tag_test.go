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

package tag_test

import (
	"fmt"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/guide/tag"
)

func ExampleNewGuideTagOption() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := tag.NewGuideTagOption(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelguidetagoption() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := tag.Delguidetagoption(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAddGuideTagOption() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := tag.AddGuideTagOption(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGuideTagOption() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := tag.GetGuideTagOption(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAddGuideBuyerTag() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := tag.AddGuideBuyerTag(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGuideBuyerTag() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := tag.GetGuideBuyerTag(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleQueryGuideBuyerByTag() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := tag.QueryGuideBuyerByTag(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelGuideBuyerTag() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := tag.DelGuideBuyerTag(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAddGuideBuyerDisplayTag() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := tag.AddGuideBuyerDisplayTag(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGuideBuyerDisplayTag() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := tag.GetGuideBuyerDisplayTag(ctx, payload)

	fmt.Println(resp, err)
}
