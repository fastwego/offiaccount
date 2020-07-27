package util

import (
	"crypto/aes"
	"crypto/cipher"
)

//解密消息函数
func AesDecryptMessage(cipherText []byte, aesKey []byte, iv []byte) (rawData []byte, err error) {
	const blockSize = 32

	if len(cipherText) < blockSize {
		panic("cipher too  stort")
	}
	plainText := make([]byte, len(cipherText))

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plainText, cipherText)

	unPadding := int(plainText[len(plainText)-1])
	if unPadding < 1 || unPadding > blockSize {
		panic("Padding is incorrect")
	}
	plainText = plainText[:len(plainText)-unPadding]

	rawData = plainText
	return
}
