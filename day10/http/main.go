package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	/*
		具体查看一下http协议
	*/
	fmt.Printf("request method: %s\n", r.Method)
	fmt.Printf("request host: %s\n", r.Host)
	fmt.Printf("request url: %s\n", r.URL)
	fmt.Printf("request proto: %s\n", r.Proto)
	fmt.Println("request headers")

	for k, v := range r.Header {
		fmt.Printf("%s: %v\n", k, v)
	}
	fmt.Println()
	fmt.Println("request cookie")
	for _, cookie := range r.Cookies() {
		fmt.Printf("name=%s value=%s\n", cookie.Name, cookie.Value)
	}
	fmt.Println()
	fmt.Printf("request body: ")
	io.Copy(os.Stdout, r.Body)

}
func main() {
	http.HandleFunc("/", HelloHandler) //路由,请求路径时,要去执行HelloHandler
	http.ListenAndServe(":8088", nil)
}
