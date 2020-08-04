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

package material_test

import (
	"fmt"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/guide/material"
)

func ExampleSetGuideCardMaterial() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := material.SetGuideCardMaterial(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGuideCardMaterial() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := material.GetGuideCardMaterial(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelGuideCardMaterial() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := material.DelGuideCardMaterial(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSetGuideImageMaterial() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := material.SetGuideImageMaterial(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGuideImageMaterial() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := material.GetGuideImageMaterial(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelGuideImageMaterial() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := material.DelGuideImageMaterial(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSetGuideWordMaterial() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := material.SetGuideWordMaterial(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetGuideWordMaterial() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := material.GetGuideWordMaterial(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelGuideWordMaterial() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := material.DelGuideWordMaterial(ctx, payload)

	fmt.Println(resp, err)
}
