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

package nontax_test

import (
	"fmt"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/nontax"
)

func ExampleGetBillAuthUrl() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := nontax.GetBillAuthUrl(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCreateBillCard() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := nontax.CreateBillCard(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleInsertBill() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := nontax.InsertBill(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleQueryFee() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := nontax.QueryFee(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUnifiedOrder() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := nontax.UnifiedOrder(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetOrder() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := nontax.GetOrder(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleRefund() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := nontax.Refund(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDownloadBill() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := nontax.DownloadBill(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleNotifyInconsistentOrder() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := nontax.NotifyInconsistentOrder(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMockNotification() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := nontax.MockNotification(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMockQueryFee() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := nontax.MockQueryFee(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMicroPay() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := nontax.MicroPay(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetOrderList() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := nontax.GetOrderList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleRealNameGetAuthUrl() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := nontax.RealNameGetAuthUrl(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetRealName() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := nontax.GetRealName(ctx, payload)

	fmt.Println(resp, err)
}
