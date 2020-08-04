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

package datacube_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/datacube"
)

func ExampleGetUserSummary() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := datacube.GetUserSummary(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUserCumulate() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := datacube.GetUserCumulate(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetArticleSummary() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := datacube.GetArticleSummary(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetArticleTotal() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := datacube.GetArticleTotal(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUserRead() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := datacube.GetUserRead(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUserReadHour() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := datacube.GetUserReadHour(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUserShare() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := datacube.GetUserShare(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUserShareHour() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := datacube.GetUserShareHour(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUpstreamMsg() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := datacube.GetUpstreamMsg(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUpstreamMsgHour() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := datacube.GetUpstreamMsgHour(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUpstreamMsgWeek() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := datacube.GetUpstreamMsgWeek(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUpstreamMsgMonth() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := datacube.GetUpstreamMsgMonth(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUpstreamMsgDist() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := datacube.GetUpstreamMsgDist(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUpstreamMsgDistWeek() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := datacube.GetUpstreamMsgDistWeek(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetUpstreamMsgDistMonth() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := datacube.GetUpstreamMsgDistMonth(ctx, payload)

	fmt.Println(resp, err)
}

func ExamplePublisherStat() {
	var ctx *offiaccount.OffiAccount

	params := url.Values{}
	resp, err := datacube.PublisherStat(ctx, params)

	fmt.Println(resp, err)
}

func ExampleGetInterfaceSummary() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := datacube.GetInterfaceSummary(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetInterfaceSummaryHour() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := datacube.GetInterfaceSummaryHour(ctx, payload)

	fmt.Println(resp, err)
}
