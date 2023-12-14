package main

import (
	"fmt"
	"os/exec"
)

func sys_call() {
	//查看系统命令所在的目录,确保命令已安装
	cmd_path, err := exec.LookPath("df")
	if err != nil {
		fmt.Println("could not found echo")
	}
	fmt.Printf("command echo in path %s\n", cmd_path)

	cmd := exec.Command("df", "-h") //相当于df -h
	// cmd.Output()                    //运行命令并获得其输出结果
	if output, err := cmd.Output(); err != nil {
		fmt.Println("got output failed", err)
	} else {
		fmt.Println(string(output))
	}

	cmd = exec.Command("rm", "./data/test.log")
	err = cmd.Run()
	if err != nil {
		fmt.Println("run failed", err)
	}
}

func main() {
	sys_call()
}
