package main

import (
	"fmt"
	"runtime"
	"time"
)

type Glimit struct {
	limit int
	ch    chan struct{}
}

func NewGlimit(limit int) *Glimit {
	return &Glimit{
		limit: limit,
		ch:    make(chan struct{}, limit),
	}
}

func (g *Glimit) Run(f func()) {
	g.ch <- struct{}{}
	go func() {
		f()
		<-g.ch
	}()
}
func main28() {
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for {
			<-ticker.C
			fmt.Printf("当前协程数: %d\n", runtime.NumGoroutine())
		}
	}()

	work := func() {
		time.Sleep(100 * time.Millisecond)
	}
	glimit := NewGlimit(10)
	for i := 0; i < 1000; i++ {
		glimit.Run(work)
	}
	time.Sleep(10 * time.Second)

}
