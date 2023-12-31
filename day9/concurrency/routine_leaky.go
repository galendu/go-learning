package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
	// "golang.org/x/net/context"
)

// 模拟一个耗时较长的任务
func work() {
	time.Sleep(time.Duration(500) * time.Millisecond)
	return
}

// 模拟一个接口处理函数
func handle() {
	//借助于带超时的context来实现对函数的超时控制
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100) //改成1000测试
	defer cancel()
	// begin :=time.Now()
	workDone := make(chan struct{})
	go func() {
		work()
		workDone <- struct{}{}
	}()
	select {
	case <-workDone:
		fmt.Println("LongTimeWork return")
	case <-ctx.Done():
		fmt.Println("LongTimeWork timeout")
	}
}

func main26() {
	for i := 0; i < 10; i++ {
		handle()
	}
	time.Sleep(2 * time.Second)
	fmt.Printf("当前协程数: %d", runtime.NumGoroutine())
}
