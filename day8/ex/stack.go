package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	slc []int
	len int
	cap int
}

func NewStack(n int) *Stack {
	slc := make([]int, n)
	return &Stack{
		slc: slc,
		len: 0,
		cap: n,
	}
}

// slc={2 5 3 5 7 0} len=3
func (stack *Stack) Push(ele int) error {
	if stack.len >= stack.cap {
		return errors.New("stack already full")
	}
	stack.slc[stack.len] = ele
	stack.len += 1
	return nil
}

func (stack *Stack) Pop() (int, error) {
	if stack.len == 0 {
		return 0, errors.New("stack already empty")
	}
	top := stack.slc[stack.len-1]
	stack.len -= 1
	return top, nil
}

func main() {

	stack := NewStack(5)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	// fmt.Println(stack.Push(6))
	n, err := stack.Pop()
	fmt.Println(n, err)
	n, err = stack.Pop()
	fmt.Println(n, err)
	n, err = stack.Pop()
	fmt.Println(n, err)
	n, err = stack.Pop()
	fmt.Println(n, err)
	n, err = stack.Pop()
	fmt.Println(n, err)

	n, err = stack.Pop()
	fmt.Println(n, err)

	fmt.Print(stack.slc)
}
