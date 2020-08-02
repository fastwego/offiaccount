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
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	payload := []byte("{}")
	resp, err := customservice.KfaccountAdd(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKfaccountUpdate() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	payload := []byte("{}")
	resp, err := customservice.KfaccountUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKfaccountDel() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	payload := []byte("{}")
	resp, err := customservice.KfaccountDel(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUploadHeadImg() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	media := ""
	params := url.Values{}
	resp, err := customservice.UploadHeadImg(ctx, media, params)

	fmt.Println(resp, err)
}

func ExampleGetKfList() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	resp, err := customservice.GetKfList(ctx)

	fmt.Println(resp, err)
}

func ExampleSendMessage() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	payload := []byte("{}")
	resp, err := customservice.SendMessage(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleTyping() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	payload := []byte("{}")
	resp, err := customservice.Typing(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetOnlineKfList() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	resp, err := customservice.GetOnlineKfList(ctx)

	fmt.Println(resp, err)
}

func ExampleInviteWorker() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	payload := []byte("{}")
	resp, err := customservice.InviteWorker(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKfSessionCreate() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	payload := []byte("{}")
	resp, err := customservice.KfSessionCreate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleKfSessionGet() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	params := url.Values{}
	resp, err := customservice.KfSessionGet(ctx, params)

	fmt.Println(resp, err)
}

func ExampleKfSessionGetList() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	params := url.Values{}
	resp, err := customservice.KfSessionGetList(ctx, params)

	fmt.Println(resp, err)
}

func ExampleKfSessionGetWaitCase() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	resp, err := customservice.KfSessionGetWaitCase(ctx)

	fmt.Println(resp, err)
}

func ExampleGetMsgList() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	payload := []byte("{}")
	resp, err := customservice.GetMsgList(ctx, payload)

	fmt.Println(resp, err)
}
