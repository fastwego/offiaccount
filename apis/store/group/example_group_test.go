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

package group_test

import (
	"fmt"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/store/group"
)

func ExampleAdd() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := group.Add(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDel() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := group.Del(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePropertyMod() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := group.PropertyMod(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleProductMod() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := group.ProductMod(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetAll() {
	var ctx *offiaccount.OffiAccount

	resp, err := group.GetAll(ctx)

	fmt.Println(resp, err)
}

func ExampleGetById() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := group.GetById(ctx, payload)

	fmt.Println(resp, err)
}
