package main

import "fmt"

func main3() {
	fmt.Println("please input two word")
	var word1 string
	var word2 string
	fmt.Scan(&word1, &word2) //读入多个单词,空格分隔. 如果输入了更多单词会背缓存起来,丢给下一次scan

	fmt.Println("please input an int")
	var i int
	fmt.Scanf("%d", &i) //类似于scan,转为特定格式的数据
}
