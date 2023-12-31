package main

import (
	"fmt"
	"time"
)

var asyncChann = make(chan int, 1) //缓冲长度为1, put可以比take多一次

func producer() {
	for i := 0; i < 10; i++ {
		asyncChann <- i //往channel里面send一个元素
		fmt.Printf("send %d\n", i)
	}
}

func consumer() {
	for i := 0; i < 3; i++ {
		v := <-asyncChann //从channel里take一个元素
		fmt.Printf("take %d\n", v)
	}
}

func main17() {
	go consumer()
	// producer() //缓冲满时,在main协程时继续send会fatal error
	go producer() //缓冲满时,在子协程时继续send会阻塞
	time.Sleep(time.Second)
}
