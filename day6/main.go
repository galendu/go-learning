package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/golang/glog"
)

// 读文件
func read_file() {
	if fin, err := os.Open("go.mod"); err != nil {
		fmt.Printf("Error opening Go.mod: %v\n", err)
	} else {
		defer fin.Close()

		// 读二进制文件
		cont := make([]byte, 10)
		if n, err := fin.Read(cont); err != nil { // 读出len(cont)个字节,返回成功读取的字节数
			fmt.Printf("read_file failed: %v\n", err)
		} else {
			fmt.Println(string(cont[:n]))
			if m, err := fin.ReadAt(cont, int64(n)); err != nil { //从指定的位置开始读len(cont)个字节
				fmt.Printf("read_file failed: %v\n", err)
			} else {
				fmt.Println(string(cont[:m]))
			}
			fin.Seek(int64(n), 0) // whence: 0从文件开头计算偏移量,1从当前位置计算偏移量,2到文件末尾的偏移量
			if n, err = fin.Read(cont); err != nil {
				fmt.Printf("read_file failed: %v\n", err)
			} else {
				fmt.Println(string(cont[:n]))
			}
		}

		// 读文本文件
		fin.Seek(0, 0) //定位到文件开头
		reader := bufio.NewReader(fin)
		for {
			if line, err := reader.ReadString('\n'); err != nil { //指定分隔符
				if err == io.EOF {
					if len(line) > 0 { //如果最后一行没有换行符,则此时最后一行就存在line里
						fmt.Println(line)
					}
					break //已经读到文件末尾
				} else {
					fmt.Printf("read file failed: %v\n", err)
				}
			} else {
				line = strings.TrimRight(line, "\n") //line里面是包含换行符的,需要去掉
				fmt.Println(line)
			}
		}
	}

}

// 写文件
func write_file() {
	//OpenFile()比Open()有更多的参数选项. os.O_WRONLY以只写的方式打开文件,os.O_TRUNC把文件之前的内容先清空掉,os.O_CREATE如果文件不存在则先创建,0666新建文件的权限设置
	if fout, err := os.OpenFile("data/verse.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666); err != nil {
		fmt.Printf("open file failed: %v\n", err)
	} else {
		defer fout.Close() //关闭文件句柄

		//写文件建议使用
		writer := bufio.NewWriter(fout)
		writer.WriteString("爱你一生一世")
		writer.WriteString("\n") //需要手动写入换行符
		writer.WriteString("爱你一生一世")
		writer.Flush() //buffer中的数据量积累到一定程度后才会真正写入磁盘.调用Flush强行把缓冲中的所有内容写入磁盘
	}
}

// 创建文件
func create_file() {
	os.Remove("data/verse.txt")
	if file, err := os.Create("data/verse.txt"); err != nil {
		fmt.Printf("create file failed: %v\n", err)
	} else {
		file.Chmod(0666)                 // set permissions
		fmt.Printf("fd=%d\n", file.Fd()) // get file descriptor, this is a integer
		info, _ := file.Stat()
		fmt.Printf("id dir %t\n", info.IsDir())
		fmt.Printf("modify time %s\n", info.ModTime())
		fmt.Printf("file name %s\n", info.Name())
		fmt.Printf("size %d\n", info.Size())
	}

	os.Mkdir("data/sys", os.ModePerm)          // create directory and set permissions
	os.MkdirAll("data/sys/a/b/c", os.ModePerm) //create multiple directories

	os.Rename("data/sys/a", "data/sys/p")       // rename file or directory name
	os.Rename("data/sys/p/b/c", "data/sys/p/c") // rename file and move function

	os.Remove("data/sys")    //delete file or directory,if it exists
	os.RemoveAll("data/sys") // recursive delete
}

// 遍历一个目录
func walk(path string) error {
	if subFiles, err := ioutil.ReadDir(path); err != nil {
		return err
	} else {
		for _, file := range subFiles {
			fmt.Println(file.Name())
			if file.IsDir() {
				if err := walk(path + "/" + file.Name()); err != nil {
					return err
				}
			}
		}
	}
	return nil

}

// 打日志
func logger() {
	log.Printf("%d=%d=%d\n", 3, 4, 3+4)
	log.Println("Hello Golang!")
	// log.Fatalln("Bye, the world") //日志输出后会执行os.Exit(1)

	// 以append方式打开日志文件
	fout, err := os.OpenFile("data/test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("open log file failed: %v\n", err)
	}

	defer fout.Close()
	logWriter := log.New(fout, "[MY_BIZ]", log.Ldate|log.Lmicroseconds) //通过flag参数定义日志的格式,时间精确到微秒1E-6s
	logWriter.Printf("%d+%d=%d\n", 3, 4, 3+4)
	logWriter.Println("Hello Golang")
	// logWriter.Fatalln("Bye, the world")
}

func logger1() {
	log.Printf("%d\n", 5)

	fout, err := os.OpenFile("my.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return
	}
	defer fout.Close()
	logWriter := log.New(fout, "CHINA ", log.Ldate|log.Lmicroseconds)
	logWriter.Printf("%s\n", "abc")
	glog.Info("test")
	glog.Error("error creating")
	glog.Warning("warning")
	glog.Fatal("Fatal error")
}

// 执行系统命令
func sys_call() {
	//查看系统命令所在的目录,确保命令已安装
	cmd_path, err := exec.LookPath("df")
	if err != nil {
		fmt.Println("cloud not found command echo")
	}
	fmt.Printf("command echo in path %s\n", cmd_path) // /bin/df

	cmd := exec.Command("df", "-h") //相当于df -h
	// cmd.Output() //运行命令并获取其输出结果
	if output, err := cmd.Output(); err != nil {
		fmt.Println("got output failed:", err)

	} else {
		fmt.Println(string(output))
	}

	cmd = exec.Command("rm", "./data/test.log")
	err = cmd.Run()
	if err != nil {
		fmt.Println("run failed:", err)
	}
}

func main() {
	// walk("../")
	// read_file()
	// logger()
	// sys_call()
	logger1()
}
