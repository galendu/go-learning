package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func read_file() {
	// fin, err := os.Open("data/verse.txt")

	// cont := make([]byte, 10)
	// fin.Read(cont)             //读出len(count)个字节,返回成功读取的字节数
	// fin.ReadAt(cont, int64(n)) //从指定的位置开始读len(cont)个字节
	// fin.Seek(int64(n), 0)      // 重新定位. whence: 0从文件开头计算偏移量,1从当前位置计算偏移量,2到文件末尾的偏移量
	if fin, err := os.Open("data/digit.txt"); err != nil {
		fmt.Printf("open data/digit.txt failed: %v\n", err)
	} else {
		defer fin.Close() //关闭文件句柄

		//读二进制文件
		cont := make([]byte, 1024)
		if n, err := fin.Read(cont); err != nil { //读出len(cont)个字节,返回成功读取的字节数
			fmt.Printf("read file failed: %v\n", err)
		} else {
			fmt.Println(string(cont[:n]))
			if m, err := fin.ReadAt(cont, int64(n)); err != nil { //从指定的位置开始读len(cont)个字节
				fmt.Printf("read file failed: %v\n", err)
			} else {
				fmt.Println(string(cont[:m]))
			}
			fin.Seek(int64(n), 0) //whence: 0从文件开头计算偏移量,1从当前位置计算偏移量,2到文件末尾的偏移量
			if n, err = fin.Read(cont); err != nil {
				fmt.Printf("read file failed: %v\n", err)
			} else {
				fmt.Println(string(cont[:n]))
			}
		}

		//读文本文件建议用bufio.Reader
		fin.Seek(0, 0) //定位到文件开头
		reader := bufio.NewReader(fin)
		for {
			if line, err := reader.ReadString('\n'); err != nil {
				if err == io.EOF {
					if len(line) > 0 {
						fmt.Println(line)
					}
					break
				} else {
					fmt.Printf("read file failed: %v\n", err)
				}
			} else {
				line = strings.TrimRight(line, "\n")
				fmt.Println(line)
			}
		}
	}
}

func writer_file() {
	if fout, err := os.OpenFile("data/verse.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666); err != nil {
		fmt.Printf("open file failed: %v\n", err)
	} else {
		defer fout.Close()

		//写文本文件建议使用
		writer := bufio.NewWriter(fout)
		writer.WriteString("明月多情应笑我")
		writer.WriteString("\n")
		writer.WriteString("笑我如今")
		writer.Flush() //buffer中的数据量积累到一定程度后才会真正写入磁盘. 调用flush强行把缓存中的所有内容写入磁盘
	}
}

func create_file() {
	os.Remove("data/verse.txt") //先删除,不去理会Remove可能返回的error
	if file, err := os.Create("data/verse.txt"); err != nil {
		fmt.Printf("create file failed: %v\n", err)
	} else {
		file.Chmod(0666)                 //设置文件权限
		fmt.Printf("fd=%d\n", file.Fd()) //获取文件描述符file descriptor,这是一个整数
		info, _ := file.Stat()
		info.IsDir()
		info.ModTime()
		info.Mode()
		info.Name()
		info.Size()
	}

	os.Mkdir("data/sys", os.ModePerm)          //穿件目录并设置权限
	os.MkdirAll("data/sys/a/b/c", os.ModePerm) //增强版Mkdir,沿途的目录不存在时会一并创建

	os.Rename("data/sys/a", "data/sys/p")        //给文件或目录重命名
	os.Rename("data/sys/p/b/c", "data/sys/pc/c") //Rename还可以实现move的功能

	os.Remove("data/sys")    //删除文件或目录,目录不为空时才能删除成功
	os.RemoveAll("data/sys") //递归删除
}
func main() {
	// func os.Open(name string) (*os.File,error)
	// fout, err := os.OpenFile("data/verse.txt",os.O_CREATE|os.O_TRUNC|os.O_WRONLY,0666)
	// read_file()
	// writer_file()
	create_file()

}
