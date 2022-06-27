package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AesEncrypt AES对称加密
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, err
}

// AesDecrypt AES解密
func AesDecrypt(cryped, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(cryped))
	blockMode.CryptBlocks(origData, cryped)
	origData = PKCS7UnPadding(origData)
	return origData, err
}

func main() {
	text := "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3" // 你要加密的数据
	AesKey := []byte("#HvL%$o0oNNoOZnk#o2qbqCeQB1iXeIR")                       // 对称秘钥长度必须是16的倍数
	encrypt, err := AesEncrypt([]byte(text), AesKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(encrypt))
	decodeString, err := base64.StdEncoding.DecodeString("9UgSuNck6qg0uPrO6XO+SS5/b70Gbf5mmT2ILlhEhRF0hJFSjgqhda5xL2VgKHEeT4EG9nUDd6uxzbf5nYA2sg4LUfm4kt0EJrJZ50xtwcY=")
	origin, err := AesDecrypt(decodeString, AesKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origin))

}
