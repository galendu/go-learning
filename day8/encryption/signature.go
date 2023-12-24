package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func ReadFile(keyFile string) ([]byte, error) {
	if f, err := os.Open(keyFile); err != nil {
		return nil, err
	} else {
		content := make([]byte, 4096)
		if n, err := f.Read(content); err != nil {
			return nil, err
		} else {
			return content[:n], nil
		}
	}
}

// go标准库不支持私钥加密,但直接提供了签名函数
func DigitalSignature(trade string) []byte {
	sha1 := sha1.New()
	sha1.Write([]byte(trade))
	digest := sha1.Sum([]byte(""))

	privateKey, err := ReadFile("rsa/data/rsa_private_key.pem")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	block, _ := pem.Decode(privateKey)
	priv, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	//用私钥生成签名
	signature, _ := rsa.SignPKCS1v15(nil, priv, crypto.Hash(0), digest)
	return signature
}

// 验证数字签名
func VerifySignature(trade string, signature []byte) bool {
	sha1 := sha1.New()
	sha1.Write([]byte(trade))
	digest := sha1.Sum([]byte(""))

	publicKey, _ := ReadFile("rsa/data/rsa_public_key.pem")
	block, _ := pem.Decode(publicKey)
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	pub := pubInterface.(*rsa.PublicKey)

	//用公钥验证签名
	return rsa.VerifyPKCS1v15(pub, crypto.Hash(0), digest, signature) == nil
}
func main() {

	trade := "jafardu transfer 10 BTC to jack"
	signature := DigitalSignature(trade)
	fmt.Println("验证数字签名", VerifySignature(trade, signature))
}
