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

package comment_test

import (
	"fmt"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/comment"
)

func ExampleOpen() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := comment.Open(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleClose() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := comment.Close(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleList() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := comment.List(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleMarkElect() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := comment.MarkElect(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleUnMarkElect() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := comment.UnMarkElect(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDelete() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := comment.Delete(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleReplyAdd() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := comment.ReplyAdd(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleReplyDelete() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := comment.ReplyDelete(ctx, payload)

	fmt.Println(resp, err)
}
