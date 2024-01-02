package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to jafardu server")
}
func BoyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to jafardu server by boy")
}

func GirlHandler(w http.ResponseWriter, r *http.Request) {
	io.Copy(os.Stdout, r.Body)
	fmt.Println(r.Body)
	for k, v := range r.Header {
		fmt.Printf("%s=%v\n", k, v)
	}
	fmt.Fprint(w, "Welcome to jafardu server by girl")

}
func main() {

	//定义路由
	http.HandleFunc("/", Handler)
	http.HandleFunc("/boy", BoyHandler)
	http.HandleFunc("/girl", GirlHandler)
	//把服务启动起来
	http.ListenAndServe(":8088", nil) //如果不发生error,会一直阻塞
}
