package main

import (
	"fmt"
	"math/rand"
)

// 随机数生成器

func main() {
	//创建一个Rand
	source := rand.NewSource(1) //seed相同的情况下,随机数生成器产生的数列是相同的
	rander := rand.New(source)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", rander.Intn(100))
	}
	fmt.Println()
	source.Seed(1) //必须重置一下seed
	rander2 := rand.New(source)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", rander2.Intn(100))
	}
	fmt.Println()

	// 使用全局Rand
	rand.Seed(1) //如果对两次运行没有一致性要求,可以不设seed
	// rand.New(rand.NewSource(1))
	// newGenerator=rand.New(rand.NewSource(1))

	fmt.Println(rand.Int())     //随机生成一个整数
	fmt.Println(rand.Float32()) //随机生成一个浮点数
	fmt.Println(rand.Intn(100)) //100以内的随机整数,[0,100]
	fmt.Println(rand.Perm(100)) //把[0,100]上的整数随机打乱
	arr := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	fmt.Println(arr)

}
