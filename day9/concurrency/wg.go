package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func Add() {
	defer wg.Done()
	time.Sleep(3000 * time.Millisecond)
	fmt.Print("over\n")
}

func sub() {
	time.Sleep(2 * time.Second)
	fmt.Println("sub over")
}
func main1() {
	wg.Add(2) //2

	// go Add() //开启了一个协程

	go func() {
		defer wg.Done()
		go sub()
		// time.Sleep(3000 * time.Millisecond)
		fmt.Print("over2\n")
	}()

	go func() {
		defer wg.Done() //减一
		go sub()
		// time.Sleep(3000 * time.Millisecond)
		fmt.Print("over2\n")

	}()

	wg.Wait() //等 减到0
	//子协程运行结束之后(即退出之后),并不妨碍孙协程的继续执行
	time.Sleep(3 * time.Second) //等孙协程

}
