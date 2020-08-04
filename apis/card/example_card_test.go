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

package card_test

import (
	"fmt"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/card"
)

func ExampleCreate() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.Create(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSetPayCell() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.SetPayCell(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSetSelfConsumeCell() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.SetSelfConsumeCell(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCreateQRCode() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.CreateQRCode(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCreateLandingPage() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.CreateLandingPage(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMpnewsGetHtml() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.MpnewsGetHtml(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSetTestWhitelist() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.SetTestWhitelist(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetCode() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.GetCode(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleConsumeCode() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.ConsumeCode(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDecryptCode() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.DecryptCode(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUserCardList() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.GetUserCardList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGet() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.Get(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBatchGet() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.BatchGet(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdate() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.Update(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleModifyStock() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.ModifyStock(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdateCode() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.UpdateCode(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelete() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.Delete(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUnavailableCoed() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.UnavailableCoed(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetCardBizUinInfo() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.GetCardBizUinInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetCardInfo() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.GetCardInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetMemberCardInfo() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.GetMemberCardInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetMemberCardDetail() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := card.GetMemberCardDetail(ctx, payload)

	fmt.Println(resp, err)
}
