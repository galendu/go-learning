package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

var (
	wg     sync.WaitGroup
	ctx    context.Context
	cancle context.CancelFunc
)

func init() {
	wg = sync.WaitGroup{}
	wg.Add(3)
	ctx, cancle = context.WithCancel(context.Background())
}

func listenSignal() {
	defer wg.Done()
	c := make(chan os.Signal)
	//监听指定信号 SIGINT和SIGTERM. 按下control+c向进程发送SIGINT信号
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-ctx.Done():
			return
		case sig := <-c:
			fmt.Printf("got signal %d\n", sig)
			cancle()
			return
		}
	}
}

func listenHttp(port int) {
	defer wg.Done()
	server := &http.Server{Addr: ":" + strconv.Itoa(port), Handler: nil}
	go func() {
		for {
			select {
			case <-ctx.Done():
				server.Close()
				return
			}
		}
	}()
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("stop listen on port %d\n", port)
}
func main() {

	go listenSignal()
	go listenHttp(8083)
	go listenHttp(8084)
	wg.Wait()
}
