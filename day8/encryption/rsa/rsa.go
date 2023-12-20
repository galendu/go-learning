package main

import "os"

/**
生成1024位的RSA私钥：
openssl genrsa -out data/rsa_private_key.pem 1024
根据私钥生成公钥：
openssl rsa -in data/rsa_private_key.pem -pubout -out data/rsa_public_key.pem

pem是一种标准格式，它通常包含页眉和页脚
*/

var (
	publicKey  []byte
	privateKey []byte
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

func ReadRSAKey(publicKeyFile, privateKeyFile string) (err error) {
	if publicKey, err = ReadFile(publicKeyFile); err != nil {
		return err
	}
	if privateKey, err = ReadFile(privateKeyFile); err != nil {
		return err
	}
	return
}
