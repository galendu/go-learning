package main

import (
	"crypto/aes"
	"crypto/cipher"
	"day8/encryption/common"
	"encoding/hex"
	"fmt"
	"log"
)

func AesEncrypt(text string, key []byte) (string, error) {
	blockSize := aes.BlockSize //AES的分组大小为16
	src := []byte(text)
	src = common.ZeroPadding(src, blockSize) //填充
	encrypted := make([]byte, len(src))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	encrypter := cipher.NewCBCEncrypter(block, key) //CBC分组模式加密
	encrypter.CryptBlocks(encrypted, src)
	return hex.EncodeToString(encrypted), nil
}

func AesDecrypt(text string, key []byte) (string, error) {
	src, err := hex.DecodeString(text) //转为[]byte
	if err != nil {
		return "", err
	}
	decrypted := make([]byte, len(src))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	edecrypter := cipher.NewCBCDecrypter(block, key) //CBC分组模式解密
	edecrypter.CryptBlocks(decrypted, src)
	out := common.ZeroUnPadding(decrypted)
	return string(out), nil
}

func main1() {
	key := []byte("ir489u96ir489u54") //key必须是长度为16的byte数组
	plain := "因为我们没有什么不同"
	cipher, err := AesEncrypt(plain, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("密文: %s\n", cipher)

	plain, err = AesDecrypt(cipher, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("明文: %s\n", plain)
}
