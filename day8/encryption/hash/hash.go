package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Sha1(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum(nil))
}

func Md5(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	return hex.EncodeToString(md5.Sum(nil))
}

func main4() {
	data := "因为我们没有什么不同"
	sha256 := sha256.New()

	fmt.Printf("Md5: %s\n", hex.EncodeToString(sha256.Sum(nil)))
	fmt.Printf("SHA-1: %s\n", Sha1(data))
	fmt.Printf("Md5: %s\n", Md5(data))
}
