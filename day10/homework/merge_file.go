package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"
)

var fileChan = make(chan string, 10000)

// var readFinish = make(chan struct{}, 3)
var writeFinish = make(chan struct{}, 3)
var wg sync.WaitGroup

func readFile(fileName string) {
	defer wg.Done()
	//打开文件
	fin, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fin.Close()
	//构建FileReader
	reader := bufio.NewReader(fin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if len(line) > 0 {
					line += "\n"
					fileChan <- line
					fmt.Printf("put %s to fileChan\n", line)
				}
				break
			} else {
				fmt.Println(err)
				break
			}
		} else {
			fileChan <- line
		}
	}
}

func writeFile(fileName string) {
	defer close(writeFinish)
	fout, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fout.Close()
	writer := bufio.NewWriter(fout)
	// LOOP:
	// 	for {
	// 		select {
	// 		case <-readFinish:

	// 			close(fileChan)
	// 			for line := range fileChan {
	// 				writer.WriteString(line)
	// 			}
	// 			break LOOP
	// 		case line := <-fileChan:
	// 			writer.WriteString(line)
	// 		}
	// 		// line := <-fileChan
	// 		// writer.WriteString(line)
	// 	}
	for {
		if line, ok := <-fileChan; ok {
			writer.WriteString(line)
		} else {
			break
		}
	}
	writer.Flush()
}

func main() {

	wg.Add(3)
	for i := 1; i <= 3; i++ {
		fileName := "dir/" + strconv.Itoa(i)
		go readFile(fileName)
	}
	go writeFile("dir/merge")

	wg.Wait()
	// close(readFinish)
	close(fileChan)
	<-writeFinish
	time.Sleep(time.Second)
}
