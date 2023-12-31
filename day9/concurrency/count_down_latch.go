package main

import (
	"fmt"
	"time"
)

func upstream(ch chan struct{}) {
	time.Sleep(15 * time.Millisecond)
	fmt.Println("一个上游协程执行结束")
	ch <- struct{}{}
}

func downstream(ch chan struct{}) {
	<-ch
	fmt.Println("下游协程开始执行")
}

func main21() {
	upstreamNum := 4   //上游协程的数量
	downstreamNum := 5 //下游协程的数量

	upstreamCh := make(chan struct{}, upstreamNum)
	downstreamCh := make(chan struct{}, downstreamNum)

	//启动上游协程和下游协程,实际下游协程会先阻塞
	for i := 0; i < upstreamNum; i++ {
		go upstream(upstreamCh)
	}
	for i := 0; i < downstreamNum; i++ {
		go downstream(downstreamCh)
	}

	//同步点
	for i := 0; i < upstreamNum; i++ {
		<-upstreamCh
	}

	//通过管道让下游协程开始执行
	for i := 0; i < downstreamNum; i++ {
		downstreamCh <- struct{}{}
	}

	time.Sleep(10 * time.Millisecond) //等下游协程执行结束

}
