package main

import (
	"fmt"
	"time"
)

func main8() {

	go func() {
		lock.RLock()
		fmt.Println("A lock successfully")
	}()

	fmt.Println()

	go func() {
		lock.RLock()
		fmt.Println("B lock successfully")

	}()

	fmt.Println()

	go func() {
		lock.Lock()
		fmt.Println("C lock successfully")
	}()

	time.Sleep(1 * time.Second)
}
