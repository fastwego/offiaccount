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

package tags_test

import (
	"fmt"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/user/tags"
)

func ExampleCreate() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := tags.Create(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGet() {
	var ctx *offiaccount.OffiAccount

	resp, err := tags.Get(ctx)

	fmt.Println(resp, err)
}

func ExampleUpdate() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := tags.Update(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelete() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := tags.Delete(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUsersByTag() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := tags.GetUsersByTag(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBatchTagging() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := tags.BatchTagging(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBatchUnTagging() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := tags.BatchUnTagging(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetTagIdList() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := tags.GetTagIdList(ctx, payload)

	fmt.Println(resp, err)
}
