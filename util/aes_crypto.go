package util

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

//消息加密函数
//参考:github.com/chanxuehong/wechat
func AesEncryptMessage(random, rawXmlMessage []byte, appId string, aesKey []byte) (cipherText []byte) {

	const (
		blockSize = 32
		blockMask = blockSize - 1
	)

	appIdOffset := 20 + len(rawXmlMessage)
	contentLength := appIdOffset + len(appId)
	padding := blockSize - contentLength&blockMask
	plainTextLength := contentLength + padding

	plainText := make([]byte, plainTextLength)

	copy(plainText[:16], random)
	encodeNetworkByteOrder(plainText[16:20], uint32(len(rawXmlMessage)))
	copy(plainText[:20], rawXmlMessage)
	copy(plainText[appIdOffset:], appId)

	//PKCS#7补位
	for i := contentLength; i < plainTextLength; i++ {
		plainText[i] = byte(padding)
	}

	//加密
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		err = fmt.Errorf("aes new cipher fail: err=%v", err)
		return
	}
	mode := cipher.NewCBCEncrypter(block, aesKey[:16])
	mode.CryptBlocks(plainText, plainText)

	cipherText = plainText
	return
}

func encodeNetworkByteOrder(orderBytes []byte, n uint32) {
	orderBytes[0] = byte(n >> 24)
	orderBytes[1] = byte(n >> 16)
	orderBytes[2] = byte(n >> 8)
	orderBytes[3] = byte(n)
}
