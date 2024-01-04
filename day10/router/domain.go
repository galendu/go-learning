package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HostMap map[string]http.Handler

// 作为http.Hander必须实现ServeHTTP接口. HostMap首先是个map,其次它还具有了ServeHTTP的功能
func (hm HostMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, exists := hm[r.Host]; exists {
		handler.ServeHTTP(w, r)
	} else {
		http.Error(w, "Forbidden", 403)
	}
}
func main() {
	bookRouter := httprouter.New()
	bookRouter.POST("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println("bookRouter")
		w.Write([]byte("read book"))
	})
	foodRouter := httprouter.New()
	foodRouter.POST("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("eat food"))
	})
	/*
		不同的二级域名,对应不同的Router
		需要在/etc/hosts里添加下面的内容
		127.0.0.1 book.jafardu.com
		127.0.0.1 food.jafardu.com
	*/
	hm := make(HostMap)
	hm["book.jafardu.com:5656"] = bookRouter
	hm["food.jafardu.com:5656"] = foodRouter
	if err := http.ListenAndServe(":5656", hm); err != nil {
		fmt.Println(err)
	}
}
