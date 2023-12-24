package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"
)

// 输出一个int对应的二进制表示
func binaryFormat(n int) string {
	sb := strings.Builder{}
	c := int(math.Pow(2, 31)) //最高位上是1,其他位全是0. 这里的int是64位
	for i := 0; i < 32; i++ {
		if n&c != 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
		c >>= 1 //"1"往右移一位
	}
	return sb.String()

}

func main6() {
	fmt.Printf("os arch %s, int size %d\n", runtime.GOARCH, strconv.IntSize) //int是4字节还是8字节,取决于操作系统是32位还是64位
	fmt.Printf("260  %b\n", 260)
	fmt.Println("-260  " + binaryFormat(-260)) //在对应正数二进制表示的基础上,按取反,再末尾加1
	fmt.Println("260^4 " + binaryFormat(260^4))
	fmt.Println("260|4 " + binaryFormat(260|4))
	fmt.Println("260^7 " + binaryFormat(260^7))       //^作为二元运算符时表示异或
	fmt.Println("^-260 " + binaryFormat(^-260))       //^作为一元运算符时表示按位取反,符号位保持不变
	fmt.Println("-260>>10 " + binaryFormat(-260>>10)) //正数高位补0,负数高位补1
	fmt.Println("-260<<23 " + binaryFormat(-260<<23)) //负数左移,可能变成正数
}
