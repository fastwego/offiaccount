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

package menu_test

import (
	"fmt"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/menu"
)

func ExampleCreate() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	payload := []byte("{}")
	resp, err := menu.Create(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetCurrentSelfmenuInfo() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	resp, err := menu.GetCurrentSelfmenuInfo(ctx)

	fmt.Println(resp, err)
}

func ExampleDelete() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	resp, err := menu.Delete(ctx)

	fmt.Println(resp, err)
}

func ExampleAddConditional() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	payload := []byte("{}")
	resp, err := menu.AddConditional(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelConditional() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	payload := []byte("{}")
	resp, err := menu.DelConditional(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleTryMatch() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	payload := []byte("{}")
	resp, err := menu.TryMatch(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGet() {
	ctx := offiaccount.New(offiaccount.OffiAccountConfig{
		Appid:          "APPID",
		Secret:         "SECRET",
		Token:          "TOKEN",
		EncodingAESKey: "EncodingAESKey",
	})

	resp, err := menu.Get(ctx)

	fmt.Println(resp, err)
}
