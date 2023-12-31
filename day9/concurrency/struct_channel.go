package main

import "fmt"

var sc = make(chan struct{}) //channel仅作为协程间同步的工具,不需要传递具体的数据,管道类型可以用struct{}
func subG() {
	fmt.Println("subG finish")
	sc <- struct{}{}
	// close(sc)
}
func main19() {
	go subG()
	<-sc //等待子协程结束. 关闭管道或者往管道里send一个数据,该行都可以解除阻塞
}
