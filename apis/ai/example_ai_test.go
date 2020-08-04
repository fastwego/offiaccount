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

package ai_test

import (
	"fmt"
	"net/url"

	"github.com/fastwego/offiaccount"
	"github.com/fastwego/offiaccount/apis/ai"
)

func ExampleSemantic() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := ai.Semantic(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAddVoiceToRecoForText() {
	var ctx *offiaccount.OffiAccount

	media := ""
	params := url.Values{}
	resp, err := ai.AddVoiceToRecoForText(ctx, media, params)

	fmt.Println(resp, err)
}

func ExampleQueryRecoResultForText() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	params := url.Values{}
	resp, err := ai.QueryRecoResultForText(ctx, payload, params)

	fmt.Println(resp, err)
}

func ExampleTranslateContent() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	params := url.Values{}
	resp, err := ai.TranslateContent(ctx, payload, params)

	fmt.Println(resp, err)
}

func ExampleOCRIDCard() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := ai.OCRIDCard(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleOCRBankcard() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := ai.OCRBankcard(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleOCRDrivingLicense() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := ai.OCRDrivingLicense(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleOCRBizLicense() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := ai.OCRBizLicense(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleOCRCommon() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := ai.OCRCommon(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleQRCode() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := ai.QRCode(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleSuperResolution() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := ai.SuperResolution(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleAICrop() {
	var ctx *offiaccount.OffiAccount

	payload := []byte("{}")
	resp, err := ai.AICrop(ctx, payload)

	fmt.Println(resp, err)
}
