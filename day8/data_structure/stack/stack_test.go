package main

import (
	"container/list"
	"sync"
	"testing"
)

// go自带的List已经包含了栈的功能,这里实现一个线程安全的栈
type Stack struct {
	list *list.List
	lock *sync.RWMutex
}

func NewStack() *Stack {
	list := list.New()
	lock := &sync.RWMutex{}
	return &Stack{list, lock}
}

// 方法的接受者为对象的指针,可以在方法内部修改对象的成员变量;否则修改的是成员变量的副本
func (stack *Stack) Push(value interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	ele := stack.list.Back()
	if ele == nil {
		return nil
	} else {
		return stack.list.Remove(ele) //移除ele的同时返回ele的值,注意ele不能是nil
	}
}

func (stack *Stack) Peak() interface{} {
	stack.lock.RLock()
	defer stack.lock.RUnlock()
	ele := stack.list.Back()
	if ele == nil {
		return nil
	} else {
		return ele.Value
	}
}

func (stack *Stack) Len() int {
	stack.lock.RLock()
	defer stack.lock.RUnlock()
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	stack.lock.RLock()
	defer stack.lock.RUnlock()
	return stack.list.Len() == 0
}

type (
	node struct {
		value interface{}
		prev  *node
	}
	MyStack struct {
		top    *node
		length int
		lock   *sync.RWMutex
	}
)

func NewMyStack() *MyStack {
	return &MyStack{nil, 0, &sync.RWMutex{}}
}

func (stack *MyStack) Push(value interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	n := &node{value, stack.top}
	stack.top = n
	stack.length++
}

func (stack *MyStack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.length == 0 {
		return nil
	}
	n := stack.top
	stack.length--
	return n.value
}

func (stack *MyStack) Peak() interface{} {
	stack.lock.RLock()
	defer stack.lock.RUnlock()
	if stack.length == 0 {
		return nil
	}
	return stack.top.value
}

func (stack *MyStack) Len() int {
	return stack.length
}

func (stack *MyStack) Empty() bool {
	return stack.Len() == 0
}

func BenchmarkStackPush(b *testing.B) {
	stack := NewStack()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkStackPop(b *testing.B) {
	stack := NewStack()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}

func BenchmarkMyStackPush(b *testing.B) {
	stack := NewMyStack()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkMyStackPop(b *testing.B) {
	stack := NewMyStack()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}
