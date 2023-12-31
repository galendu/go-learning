package main

import (
	"fmt"
	"time"
)

func main2() {
	arr := []int{1, 2, 3, 4}
	for _, ele := range arr {
		// go func() {
		// 	fmt.Printf("%d\n", ele)
		// }() //ele 4 4 4 4 //用的时协程外面的全局变量ele
		go func(value int) {
			fmt.Printf("%d\n", value)
		}(ele) //ele 4 2 1 3 把ele的副本传到协程内部

	}
	time.Sleep(1 * time.Second)
}
