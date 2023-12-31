package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main13() {

	const N = 10
	wg := sync.WaitGroup{}

	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(a, b int) {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			_ = a + b
		}(i, i+1)
	}
	fmt.Printf("当前协程数：%d\n", runtime.NumGoroutine())
	wg.Wait()
	fmt.Printf("当前协程数：%d\n", runtime.NumGoroutine())
}
