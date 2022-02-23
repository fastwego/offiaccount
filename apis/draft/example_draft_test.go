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

package draft_test

import (
	"fmt"
	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/draft"
)

func ExampleAddDraft() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := draft.AddDraft(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetDraft() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := draft.GetDraft(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDeleteDraft() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := draft.DeleteDraft(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUpdateDraft() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := draft.UpdateDraft(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCountDrafts() {
	var ctx *offiaccount.OffiAccount

	resp, err := draft.CountDrafts(ctx)

	fmt.Println(resp, err)
}

func ExampleGetDraftList() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := draft.GetDraftList(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleTemporaryMPSwitch() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := draft.TemporaryMPSwitch(ctx, payload)

	fmt.Println(resp, err)
}
