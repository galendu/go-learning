package main

import (
	"fmt"
	"log"
	"os"
)

func logger() {
	log.Printf("%d+%d=%d\n", 3, 4, 3+4)
	log.Println("Hello Golang")
	// log.Fatalln("Bye, the world") //日志输出后会执行os.Exit(1)

	//以append方式打开日志文件
	fout, err := os.OpenFile("data/test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("open log failed: %v\n", err)
	}

	defer fout.Close()
	logWriter := log.New(fout, "[MY_BIZ]", log.Ldate|log.Lmicroseconds|log.Lmicroseconds)
	logWriter.Printf("%d+%d=%d\n", 3, 4, 3+4)
	logWriter.Println("Hello Golang")
	logWriter.Fatalln("Bye, the world")
}

func main() {
	logger()
}
