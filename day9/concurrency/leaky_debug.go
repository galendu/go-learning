package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func main27() {
	go func() {
		//在8080端口接收debug
		if err := http.ListenAndServe("127.0.0.1:8082", nil); err != nil {
			panic(err)
		}
	}()

	go func() {
		//每隔一秒钟打印一次协程数量
		ticker := time.NewTicker(1 * time.Second)
		for {
			<-ticker.C
			fmt.Printf("当前协程数: %d\n", runtime.NumGoroutine())
		}
	}()
	for {
		handle()
	}
}
