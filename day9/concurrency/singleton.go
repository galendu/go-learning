package main

import (
	"fmt"
	"sync"
	"time"
)

var oc sync.Once
var a int = 5

func main3() {
	go func() {
		oc.Do(func() {
			a++

		})
	}()
	go func() {
		oc.Do(func() {
			a++

		})
	}()
	time.Sleep(time.Second)
	fmt.Println(a) //6
}
