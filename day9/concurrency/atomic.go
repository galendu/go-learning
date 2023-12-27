package main

import (
	"fmt"
	"sync"
	"time"
)

var n int32 //0

var lock sync.RWMutex

func main7() {
	now := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			// n++
			// atomic.AddInt32(&n, 1)
			lock.Lock()
			n++
			lock.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(n, time.Since(now))
}
