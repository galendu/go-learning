package main

import (
	"fmt"
	"sync"
	"time"
)

var globalResource = make(map[string]string)
var oc sync.Once

func LoadResource() {
	oc.Do(func() {
		fmt.Println("load global resource")
		globalResource["1"] = "A"
		globalResource["2"] = "B"
	})
}

type Singleton struct {
	Name string
}

var singleton *Singleton
var singletonOnce sync.Once

func GetSingletonInstance() *Singleton {
	singletonOnce.Do(func() {
		fmt.Println("init Singleton")
		singleton = &Singleton{Name: "Tom"}
	})
	return singleton
}

func main14() {
	go LoadResource()
	go LoadResource()
	inst1 := GetSingletonInstance()
	inst2 := GetSingletonInstance()
	fmt.Printf("inst1 address %p\n", inst1)
	fmt.Printf("inst1 address %p\n", inst2)
	time.Sleep(100 * time.Millisecond)
}
