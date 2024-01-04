package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var limitCh = make(chan struct{}, 100) //最多并发处理100个请求

func timeMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		next.ServeHTTP(rw, r)
		timeElapsed := time.Since(begin)
		log.Printf("requst %s use %d ms\n", r.URL.Path, timeElapsed.Milliseconds())
	})
}

func limitMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		limitCh <- struct{}{}
		log.Printf("current %d\n", len(limitCh))
		next.ServeHTTP(rw, r)
		<-limitCh
	})
}

type middleware func(http.Handler) http.Handler

type Router struct {
	middleWareChan []middleware
	mux            map[string]http.Handler
}

func NewRouter() *Router {
	return &Router{
		middleWareChan: make([]middleware, 0, 10),
		mux:            make(map[string]http.Handler, 10),
	}
}

// 添加中间件
func (self *Router) Use(m middleware) {
	self.middleWareChan = append(self.middleWareChan, m)
}

// 自定义路由
func (self *Router) Add(path string, handler http.Handler) {
	var mergedHandler = handler
	for i := 0; i < len(self.middleWareChan); i++ {
		mergedHandler = self.middleWareChan[i](mergedHandler)
		self.mux[path] = mergedHandler
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond * 100)
	w.Write([]byte("how are you?"))
}

func main() {

	router := NewRouter()
	router.Use(timeMiddleWare)
	router.Use(limitMiddleWare)
	router.Add("/", http.HandlerFunc(get))

	for path, handler := range router.mux {
		http.Handle(path, handler)
	}
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Println(err)
	}
}
