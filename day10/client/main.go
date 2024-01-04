package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func get() {
	resp, err := http.Get("http://127.0.0.1:8088")
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
	// resp, err := http.Post("http://127.0.0.1:8088/girl", "text/plain", reader)
	// resp, err := http.Post("http://127.0.0.1:8088/user/zcy/vip/gs/pingliang", "text/plain", reader)
	resp, err := http.Post("http://book.jafardu.com:5656", "text/plain", reader)

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
func complexHttpRequest() {
	reader := strings.NewReader("hello server girl")
	if req, err := http.NewRequest("POST", "http://127.0.0.1:8088/girl", reader); err != nil {

		panic(err)
	} else {
		//自定义请求头
		req.Header.Add("User-Agent", "中国")
		req.Header.Add("MyHeaderKey", "MyHeaderValue")
		//自定义cookie
		req.AddCookie(&http.Cookie{
			Name:  "auth",
			Value: "passwd",

			Path:    "/girl",
			Domain:  "localhost",
			Expires: time.Now().Add(time.Duration(time.Hour)),
		})
		client := &http.Client{
			Timeout: 100 * time.Millisecond,
		}
		if resp, err := client.Do(req); err != nil { //提交http请求
			fmt.Println(err)
		} else {
			defer resp.Body.Close()
			io.Copy(os.Stdout, resp.Body)
			for k, v := range resp.Header {
				fmt.Printf("%s=%v\n", k, v)
			}
			fmt.Println(resp.Proto)
			fmt.Println(resp.Status)
		}

	}
}
func main() {
	// wg := sync.WaitGroup{}
	// wg.Add(200)
	// for i := 0; i < 200; i++ {
	// 	go func() {
	// 		defer wg.Done()
	// 		post()
	// 	}()
	// }
	// wg.Wait()
	post()
	// complexHttpRequest()
}
