package main

import (
	"context"
	"fmt"
	"time"
)

const (
	WorkUseTime = 500 * time.Millisecond
	Timeout     = 100 * time.Millisecond
)

// 模拟一个耗时较长的任务
func LongTimeWork() {
	time.Sleep(WorkUseTime)
	return
}

// 模拟一个接口处理函数
func Handle1() {
	deadline := make(chan struct{}, 1)
	workDone := make(chan struct{}, 1)
	go func() {
		LongTimeWork()
		workDone <- struct{}{}
	}()
	go func() {
		time.Sleep(Timeout)
		deadline <- struct{}{}
	}()
	select { //下面的case只执行最早到来的那一个
	case <-workDone:
		fmt.Println("LongTimeWork return")
	case <-deadline:
		fmt.Println("LongTimeWork timeout")
	}
}

// 模拟一个接口处理函数
func Handle2() {
	workDone := make(chan struct{}, 1)
	go func() {
		LongTimeWork()
		workDone <- struct{}{}
	}()
	select {
	case <-workDone:
		fmt.Println("LongTimeWork return")
	case <-time.After(Timeout):
		fmt.Println("LongTimeWork timeout")
	}
}

// 模拟一个接口处理函数
func Handle3() {
	//通过显示sleep再调用cancle()来实现对函数的超时控制
	ctx, cancel := context.WithCancel(context.Background())

	workDone := make(chan struct{}, 1)
	go func() {
		LongTimeWork()
		workDone <- struct{}{}
	}()

	go func() {
		time.Sleep(Timeout)
		cancel()
	}()
	select {
	case <-workDone:
		fmt.Println("LongTimeWork return")
	case <-ctx.Done(): //ctx.Done()是一个管道,调用了cancel()都会关闭这个管道,然后读操作就会立即返回
		fmt.Println("LongTimeWork timeout")
	}
}

func Handle4() {
	//借助于带超时的context来实现对函数的超时控制
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel() //纯粹处于良好习惯,函数退出前调用cancel()
	workDone := make(chan struct{}, 1)
	go func() {
		LongTimeWork()
		workDone <- struct{}{}
	}()
	select {
	case <-workDone:
		fmt.Println("LongTimeWork return")
	case <-ctx.Done():
		fmt.Println("LongTimeWork Timeout")
	}
}

func main25() {
	Handle1()
	Handle2()
	Handle3()
	Handle4()

}
