package main

import (
	"fmt"
	"sync"
	"time"
)

type Student struct {
	age  int
	name string
}

var (
	lst     []int = make([]int, 5)
	arr     [5]int
	student Student
	mp      sync.Map
	// mp      map[int]int = make(map[int]int)
)

func rwShareMem1() {
	// for i := 1; i < len(lst); i += 2 {
	// 	lst[i] = 555
	// }

	for i := 0; i < 100; i++ {
		// mp[i] = i * 2
		mp.Store(i, i*2)
	}

}

func rwShareMem2() {
	// for i := 0; i < len(lst); i += 2 {
	// 	lst[i] = 888
	// }

	// for i := 0; i < 100; i++ {
	// 	mp[i] = i * 2
	// }
	for i := 0; i < 100; i++ {
		// mp[i] = i * 2
		mp.Store(i, i*2)
	}
}

func main9() {
	go rwShareMem1()
	go rwShareMem2()

	time.Sleep(time.Second)
	fmt.Println(mp.Load(0))
}
