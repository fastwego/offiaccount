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

package customservice_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/customservice"
)

func ExampleKfaccountAdd() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := customservice.KfaccountAdd(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKfaccountUpdate() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := customservice.KfaccountUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKfaccountDel() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := customservice.KfaccountDel(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUploadHeadImg() {
	var ctx *offiaccount.OffiAccount

	media := ""
	params := url.Values{}
	resp, err := customservice.UploadHeadImg(ctx, media, params)

	fmt.Println(resp, err)
}

func ExampleGetKfList() {
	var ctx *offiaccount.OffiAccount

	resp, err := customservice.GetKfList(ctx)

	fmt.Println(resp, err)
}

func ExampleSendMessage() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := customservice.SendMessage(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleTyping() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := customservice.Typing(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetOnlineKfList() {
	var ctx *offiaccount.OffiAccount

	resp, err := customservice.GetOnlineKfList(ctx)

	fmt.Println(resp, err)
}

func ExampleInviteWorker() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := customservice.InviteWorker(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKfSessionCreate() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := customservice.KfSessionCreate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKfSessionGet() {
	var ctx *offiaccount.OffiAccount

	params := url.Values{}
	resp, err := customservice.KfSessionGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleKfSessionGetList() {
	var ctx *offiaccount.OffiAccount

	params := url.Values{}
	resp, err := customservice.KfSessionGetList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleKfSessionGetWaitCase() {
	var ctx *offiaccount.OffiAccount

	resp, err := customservice.KfSessionGetWaitCase(ctx)

	fmt.Println(resp, err)
}

func ExampleGetMsgList() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := customservice.GetMsgList(ctx, payload)

	fmt.Println(resp, err)
}
