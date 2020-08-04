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

package guide_test

import (
	"fmt"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/guide/guide"
)

func ExampleGetGuideAcct() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := guide.GetGuideAcct(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAddGuideAcct() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := guide.AddGuideAcct(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdateGuideAcct() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := guide.UpdateGuideAcct(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelGuideAcct() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := guide.DelGuideAcct(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGuideAcctList() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := guide.GetGuideAcctList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGuideCreateQrCode() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := guide.GuideCreateQrCode(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGuideBuyerChatRecord() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := guide.GetGuideBuyerChatRecord(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSetGuideConfig() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := guide.SetGuideConfig(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGuideConfig() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := guide.GetGuideConfig(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSetGuideAcctConfig() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := guide.SetGuideAcctConfig(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGuideAcctConfig() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := guide.GetGuideAcctConfig(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePushShowWxaPathMenu() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := guide.PushShowWxaPathMenu(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleNewGuideGroup() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := guide.NewGuideGroup(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGuideGroupList() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := guide.GetGuideGroupList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGroupInfo() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := guide.GetGroupInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAddGuide2GuideGroup() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := guide.AddGuide2GuideGroup(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelGuide2GuideGroup() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := guide.DelGuide2GuideGroup(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGroupByGuide() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := guide.GetGroupByGuide(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelGuideGroup() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := guide.DelGuideGroup(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGuideBuyerRelationList() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := guide.GetGuideBuyerRelationList(ctx, payload)

	fmt.Println(resp, err)
}
