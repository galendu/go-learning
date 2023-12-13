package main

import (
	"errors"
	"fmt"
	"sync"
)

func defer_exe_time() (i int) {
	i = 9
	defer func() { // defer 后可以跟一个func
		fmt.Printf("i=%d\n", i) //打印5,而非9
	}()
	defer fmt.Printf("i=%d\n", i) //变量在注册defer时被拷贝或计算
	return 5
}

func soo(a, b int) {
	defer func() {
		//recover 必须在defer中才能生效
		if err := recover(); err != nil {
			fmt.Printf("soo函数中发生了panic:%s\n", err)
		}
	}()
	panic(errors.New("my error"))
}

// var i interface{}=3.4
// v,ok:=i.(float64) //类型断言

func square(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Println(v * v)
	case float64:
		fmt.Println(v * v)
	default:
		fmt.Printf("unsurport data type %T\n", v)
	}

}

type User struct {
	Name string
	Age  int
}

var (
	sUser *User
	uOnce sync.Once
)

func NewDefaultUser() *User {
	return &User{
		Name: "",
		Age:  -1,
	}
}
func GetUserInstance() *User {
	uOnce.Do(func() { //确保即使在并发的情况下,下面的3行代码在整个go进程里只会被执行一次
		if sUser == nil {
			sUser = NewDefaultUser()
		}
	})
	return sUser
}

func sumGenerics[T int | int32 | int64 | int8 | float32 | float64](a, b T) T {
	return a + b
}

func main() {
	// reflect.Value
	sumInt := sumGenerics(3, 5)
	sumfloat := sumGenerics(3.3, 5.9)
	sumint := sumGenerics(3.9999999999999, 5.99999999)
	fmt.Println(sumInt, sumfloat, sumint)
	fmt.Println("starting...")
	fmt.Println(defer_exe_time())
	soo(1, 2)
}
