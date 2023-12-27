package main

import (
	"fmt"
	"os"
	"time"
)

// 倒计时
func countDown(countCh chan int, n int, finishCh chan struct{}) {
	if n <= 0 { //从0开始倒数
		return
	}
	ticker := time.NewTicker(1 * time.Second) //创建一个周期性的定时器,每隔1秒钟执行一次
	for {
		countCh <- n //把n放到管道
		<-ticker.C   //等1秒钟
		n--          //n-1
		if n <= 0 {  //n减到0时退出
			ticker.Stop()          //停止定时器
			finishCh <- struct{}{} //成功结束
			break
		}
	}
}

// 中止
func abort(ch chan struct{}) {
	buffer := make([]byte, 1)
	os.Stdin.Read(buffer) //阻塞式IO,如果标准输入里没数据,该行不会执行
	ch <- struct{}{}
}

func main() {

	countCh := make(chan int)
	finishCh := make(chan struct{})
	go countDown(countCh, 10, finishCh) //开一个子协程,去往countCh和finishCh里放数据
	abortCh := make(chan struct{})
	go abort(abortCh)

Loop:
	for {
		select {
		case n := <-countCh:
			fmt.Println(n)
		case <-finishCh:
			fmt.Println("finish")
			break Loop
		case <-abortCh:
			fmt.Println("abort")
			break Loop
		}
	}
}
