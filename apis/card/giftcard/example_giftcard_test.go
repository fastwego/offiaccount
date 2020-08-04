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

package giftcard_test

import (
	"fmt"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/card/giftcard"
)

func ExamplePageAdd() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := giftcard.PageAdd(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePageGet() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := giftcard.PageGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePageUpdate() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := giftcard.PageUpdate(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePageBatchGet() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := giftcard.PageBatchGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMaintainSet() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := giftcard.MaintainSet(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePayWhitelistAdd() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := giftcard.PayWhitelistAdd(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePaySubmchBind() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := giftcard.PaySubmchBind(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleWxaSet() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := giftcard.WxaSet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleOrderGet() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := giftcard.OrderGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleOrderBatchGet() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := giftcard.OrderBatchGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGeneralCardUpdateUser() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := giftcard.GeneralCardUpdateUser(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleOrderRefund() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := giftcard.OrderRefund(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleInvoiceSetBizAttr() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := giftcard.InvoiceSetBizAttr(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleInvoiceGetAuthData() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := giftcard.InvoiceGetAuthData(ctx, payload)

	fmt.Println(resp, err)
}
