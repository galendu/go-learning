package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func get() {
	resp, err := http.Get("http://127.0.0.1:8088/girl")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() //一定要调用,否则会协程泄漏
	io.Copy(os.Stdout, resp.Body)
	for k, v := range resp.Header { //打印response header
		fmt.Printf("%s=%v\n", k, v)
	}
	fmt.Println(resp.Proto)
	fmt.Println(resp.Status)
}

func post() {

	reader := strings.NewReader("hello server")
	resp, err := http.Post("http://127.0.0.1:8088/girl", "text/plain", reader)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	for k, v := range resp.Header {
		fmt.Printf("%s=%v\n", k, v)
	}
	fmt.Println(resp.Proto)
	fmt.Println(resp.Status)
}
func main() {
	get()
	post()
}
