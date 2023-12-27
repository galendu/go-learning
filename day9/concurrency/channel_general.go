package main

import (
	"fmt"
	"time"
)

var buffer chan string

func init() {
	buffer = make(chan string, 10000)
}

func put() {
	for i := 0; i < 10; i++ {
		buffer <- "111111111"
	}
}
func take() {
	for i := 0; i < 20; i++ {
		v := <-buffer
		fmt.Println(v)
	}
}
func main6() {
	go put()
	go put()
	go put()
	go put()

	go take()
	go take()

	time.Sleep(time.Second)
}
