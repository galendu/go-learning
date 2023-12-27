package main

import "fmt"

func main5() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	close(c)
	for ele := range c {
		fmt.Println(ele)
	}
	v := <-c //close channel 之后,读操作总是会立即返回,如果channel里已没有元素,则返回'0'值
	fmt.Println(v)
}
