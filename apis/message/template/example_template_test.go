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

package template_test

import (
	"fmt"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/message/template"
)

func ExampleSetIndustry() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	payload := []byte("{}")
	resp, err := template.SetIndustry(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetIndustry() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	resp, err := template.GetIndustry(ctx)

	fmt.Println(resp, err)
}

func ExampleAddTemplate() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	payload := []byte("{}")
	resp, err := template.AddTemplate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetAllPrivateTemplate() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	resp, err := template.GetAllPrivateTemplate(ctx)

	fmt.Println(resp, err)
}

func ExampleDelPrivateTemplate() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	payload := []byte("{}")
	resp, err := template.DelPrivateTemplate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSend() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	payload := []byte("{}")
	resp, err := template.Send(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSubscribe() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	payload := []byte("{}")
	resp, err := template.Subscribe(ctx, payload)

	fmt.Println(resp, err)
}
