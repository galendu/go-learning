package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var limitChan = make(chan struct{}, 100)

func limitMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limitChan <- struct{}{}
		next.ServeHTTP(w, r)
		<-limitChan
	})
}
func timeMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		requestPath := r.URL.Path
		next.ServeHTTP(w, r)
		elapse := time.Since(begin).Milliseconds()
		log.Printf("%s use time %s", requestPath, elapse)
	})
}

func boy(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello boy"))
}

func girl(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello girl"))
}

type middleware func(next http.Handler) http.Handler
type Router struct {
	middleWareChain []middleware
	mux             map[string]http.Handler
}

func (router *Router) Use(m middleware) { //添加中间件
	router.middleWareChain = append(router.middleWareChain, m)
}

func (router *Router) Add(path string, handler http.Handler) {
	// router.mux[path] = handler
	var mergeHandler = handler
	for i := len(router.middleWareChain) - 1; i >= 0; i-- {
		mw := router.middleWareChain[i]
		mergeHandler = mw(mergeHandler)
	}
	router.mux[path] = mergeHandler
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.Path

	if handler, exists := router.mux[requestPath]; exists {
		handler.ServeHTTP(w, r)
	} else {
		http.NotFoundHandler().ServeHTTP(w, r) // 404
	}
}
func NewRouter() *Router {
	return &Router{
		middleWareChain: make([]middleware, 0),
		mux:             make(map[string]http.Handler),
	}
}
func main() {
	// http.Handle("/boy", limitMiddleWare(timeMiddleWare(http.HandlerFunc(boy))))
	// http.Handle("/girl", timeMiddleWare(http.HandlerFunc(girl)))
	router := NewRouter()
	router.Use(limitMiddleWare)
	router.Use(timeMiddleWare)
	router.Add("/boy", http.HandlerFunc(boy))
	router.Add("/girl", http.HandlerFunc(girl))
	if err := http.ListenAndServe(":5656", router); err != nil {
		fmt.Println(err)
	}
}
