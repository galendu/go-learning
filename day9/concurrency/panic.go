package main

import (
	"fmt"
	"time"
)

func F() {
	defer fmt.Println("11111")
	defer fmt.Println("22222")
	fmt.Println("GGGGGGGGGG")

	defer func() {
		recover() //从panic发生的地方中途结束本协程,但是没有结束整个进程
	}()
	defer fmt.Println("33333")
	defer fmt.Println("44444")
	panic("oooooooooops")
	defer fmt.Println("55555")
	fmt.Println("FFFFFFFFFFFFFFFFFFFFFFFF")
}
func main() {

	go F() //G 4 3 2 1 ops
	time.Sleep(time.Second)
	fmt.Println("this is main")
}
