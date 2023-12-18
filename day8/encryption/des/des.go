package main

import (
	"crypto/des"
	"day8/encryption/common"
	"encoding/hex"
	"fmt"
	"log"
)

// XOR 异或运算,要求plain和key的长度相同
func XOR(plain string, key []byte) string {
	bPlain := []byte(plain)
	bCipher := make([]byte, len(key))
	for i, k := range key {
		bCipher[i] = k ^ bPlain[i]
	}
	cipher := string(bCipher)
	return cipher
}

// DesEncrypt DES加密
// 密钥必须是64位,所以key必须是长度为8的byte数组
func DesEncrypt(text string, key []byte) (string, error) {
	src := []byte(text)
	block, err := des.NewCipher(key) //用des创建一个加密器cipher
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()           //分组的大小,blockSize=8
	src = common.ZeroPadding(src, blockSize) //填充

	out := make([]byte, len(src)) //密文和明文的长度一致
	dst := out
	for len(src) > 0 {
		//分组加密
		block.Encrypt(dst, src[:blockSize]) //对src进行加密,加密结果放到dst里,移到下一组
		src = src[blockSize:]
		dst = dst[blockSize:]
	}
	return hex.EncodeToString(out), nil
}

// DesDecrypt DES解密
// 密钥必须是64为,所以key必须是长度为8的byte数组
// func DesDecrypt()

func main() {
	plain := "ABCD"
	key := []byte{1, 2, 3, 4}
	cipher := XOR(plain, key)
	plain = XOR(cipher, key)
	fmt.Printf("plain: %s\n", plain)
	fmt.Println("-------------------------------")

	key = []byte("ir489u58") //key必须是长度为8的byte数组
	plain = "因为我们没有什么不同"
	cipher, err := DesEncrypt(plain, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("密文: %s\n", cipher)
}
