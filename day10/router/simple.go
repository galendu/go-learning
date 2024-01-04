package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func handle(method string, w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Printf("request method: %s\n", r.Method)
	fmt.Printf("request body: ")
	io.Copy(os.Stdout, r.Body)
	fmt.Println()
	w.Write([]byte("Hi boy, you request " + method))
}

func get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	handle("get", w, r, params)
}

func post(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	handle("post", w, r, params)
}

func head(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	handle("post", w, r, params)
}

func options(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	handle("post", w, r, params)
}

func put(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	handle("post", w, r, params)
}

func patch(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	handle("post", w, r, params)
}
func delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	handle("post", w, r, params)
}

func panic(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var arr []int
	_ = arr[1] //数组越界
}

func main1() {
	router := httprouter.New()
	router.GET("/", get)
	router.POST("/", post)
	router.HEAD("/", head)
	router.PUT("/", put)
	router.PATCH("/", patch)
	router.DELETE("/", delete)
	//router没有提供connect和trace
	router.POST("/user/:name/:type/*addr", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Printf("name:%s, type:%s, addr:%s\n", p.ByName("name"), p.ByName("type"), p.ByName("addr"))
	})

	router.ServeFiles("/file/*filepath", http.Dir("./static"))
	http.ListenAndServe(":8088", router) //Router实现了ServerHTTP接口,所以它是一种http.Handler
}
