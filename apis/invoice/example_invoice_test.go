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

package invoice_test

import (
	"fmt"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/invoice"
)

func ExampleGetAuthUrl() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := invoice.GetAuthUrl(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetAuthData() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := invoice.GetAuthData(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleRejectInsert() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := invoice.RejectInsert(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMakeOutInvoice() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := invoice.MakeOutInvoice(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleClearOutInvoice() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := invoice.ClearOutInvoice(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleQueryInvoceInfo() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := invoice.QueryInvoceInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSetUrl() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := invoice.SetUrl(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePlatformCreateCard() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := invoice.PlatformCreateCard(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePlatformSetpdf() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := invoice.PlatformSetpdf(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePlatformGetpdf() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := invoice.PlatformGetpdf(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleInsert() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := invoice.Insert(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePlatformUpdateStatus() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := invoice.PlatformUpdateStatus(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleReimburseGetInvoiceInfo() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := invoice.ReimburseGetInvoiceInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleReimburseGetInvoiceBatch() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := invoice.ReimburseGetInvoiceBatch(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleReimburseUpdateInvoiceStatus() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := invoice.ReimburseUpdateInvoiceStatus(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleReimburseUpdateStatusBatch() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := invoice.ReimburseUpdateStatusBatch(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUserTitleUrl() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := invoice.GetUserTitleUrl(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetSelectTitleUrl() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := invoice.GetSelectTitleUrl(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleScanTitle() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := invoice.ScanTitle(ctx, payload)

	fmt.Println(resp, err)
}
