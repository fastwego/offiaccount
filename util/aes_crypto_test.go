package util

import (
	"encoding/base64"
	"testing"
)

func TestAesEncryptMessage(t *testing.T) {
	appId := "wx3808f4e98b0e3a57"
	aesKey, err := base64.StdEncoding.DecodeString("itZDdxGBnFRaBPmFD12CcYalbatKW1fLtuIfQw8iQ57=")

	if err != nil {
		t.Error("base64 error")
		return
	}

	random := []byte("1234567890")
	plainText := []byte("<xml><ToUserName><![CDATA[gh_815dcdde3c52]]></ToUserName>\n<FromUserName><![CDATA[fastwego]]></FromUserName>\n<CreateTime>1595869404</CreateTime>\n<MsgType><![CDATA[text]]></MsgType>\n<Content><![CDATA[tests_message]]></Content>\n<MsgId>1234567</MsgId>\n</xml>")

	SuccessBase64CipherText := "70nIsuWJcN2F4vcDSXJrRbs1q9h+FbMrgoEKP3q30K4CWX9OiCrEDBZtXb0w7RneUVYq8OHo3Q492jU8Zw8mDw9dPTFjLkfxXiQinzIAtXLiOblGmrZtDWseKcg7HvHMqHzb0gnf4lG9J1eO7KY5HnanX7OB7c4TOg53Q2LYPqLOLlLV9wpBSoL9+kKwNPTMJNB1DCpai9+I+Fh4kUIwQok84Bf1PNNEStRqlxlh7tLNOeScd4UmGWi8U2rQNvcfANyuQixy1fIHz5WVRIbTGMB8uUF4j4mhbATqPOaxv+fQxgeE7zutvOmBE2mjIzmBwcoSZHxQAyyGsQQo52E1pqfuUJgr15zepEdC2+qopQdYbKkonOMzSxG36DDsGbmrmW53BRXKvtgHDeR7qdmW7kG407zPI/7Yrg+kKCqlNVY="
	cipherText := AesEncryptMessage(random, plainText, appId, aesKey)
	base64CipherText := base64.StdEncoding.EncodeToString(cipherText)

	if base64CipherText != SuccessBase64CipherText {
		t.Error("test TestAesEncryptMessage fail")
	}
	return
}
