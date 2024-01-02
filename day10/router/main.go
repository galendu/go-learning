package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func handle(method string, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request method: %s\n", r.Method)
	fmt.Printf("request body:")
	io.Copy(os.Stdout, r.Body)
	fmt.Println()
	// fmt.Fprint(w,"Hello Girl")
	w.Write([]byte("Hello Girl,your request method is " + method))
}
func get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	handle("GET", w, r)
}
func post(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	handle("POST", w, r)
}

func p1(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var a = 0
	_ = 4 / a
	var arr []int
	_ = arr[1]
}
func main() {
	router := httprouter.New()
	router.GET("/girl", get)
	router.POST("/girl", post)

	//restful
	router.POST("/user/:name/:type/*addr", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Printf("name=%s type=%s addr=%s\n", p.ByName("name"), p.ByName("type"), p.ByName("addr"))
	})

	//返回静态文件
	router.ServeFiles("/file/*filepath", http.Dir("./static"))

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprintf(w, "server panic %v", i)
	}
	router.GET("/panic", p1)

	http.ListenAndServe(":8088", router)

}
