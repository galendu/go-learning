package main

import (
	"container/list"
	"container/ring"
	"fmt"
)

// 链表中的一个元素
type Node struct {
	Info int
	Next *Node
	Prev *Node
}

// 链表
type List struct {
	Head *Node
	Len  int
	Tail *Node
}

// Head   Tail
// 1  ----->  8  	--> 	4
func (list *List) Add(ele int) {
	node := &Node{Info: ele, Next: nil, Prev: nil}
	if list.Len == 0 {
		list.Head = node
		list.Tail = node
	} else {
		list.Tail.Next = node
		node.Prev = list.Tail
		list.Tail = node
	}
	list.Len += 1
}

// 遍历链表,打印其中的每一个元素
func (list List) Travs() {
	if list.Len == 0 {
		return
	}
	head := list.Head
	fmt.Printf("%d ", head.Info)
	for i := 1; i < list.Len; i++ {
		head = head.Next
		fmt.Printf("%d ", head.Info)
	}
	fmt.Println()
}

// 1-->8 -->7 -->5 head
func (list *List) Visit(ele int) {
	if list.Len == 0 {
		return
	}
	head := list.Head
	for {
		if head == nil {
			break
		}
		if head.Info != ele {
			head = head.Next
		} else {
			break
		}
	}
	if head == nil {
		return
	} else {
		post := head.Next //5
		pre := head.Prev  //8
		pre.Next = post   //8-->5
		post.Prev = pre   //8<--5

		head.Next = list.Head //7-->1
		list.Head.Prev = head //7 <--1
		list.Head = head      //7成为了链表的head
	}
}

// 4 ->6 -> 2
func TravsList(lst *list.List) {
	head := lst.Front()
	for head.Next() != nil {
		fmt.Printf("%v ", head.Value) // 4 6
		head = head.Next()
	}
	fmt.Printf("%v ", head.Value) //2
	fmt.Println()

	//pop
	top := lst.Back()
	lst.Remove(top)

	//Push
	lst.PushBack(7)
}

func RingDemo() {
	ring := ring.New(100)
	for i := 0; i < 1000; i++ {
		ring.Value = i
		ring = ring.Next()
	}

	sum := 0
	ring.Do(func(i interface{}) { //ring.Do遍历整个环,对每个元素采取的处理措施由回调函数func(i interface{}){} 来定义
		fmt.Printf("%v ", i)
		num := i.(int)
		sum += num
	})
	fmt.Println()
	fmt.Println(sum / 100)
}

func main1() {
	// list := List{}
	// list.Add(1)
	// list.Add(8)
	// list.Add(4)
	// list.Add(3)
	// list.Add(5)
	// list.Travs()

	// list.Visit(4)
	// list.Travs()

	lst := list.New() //创建一个空的双向链表
	lst.PushBack(4)
	lst.PushBack(6)
	lst.PushBack(2)  //往链表尾部添加2
	lst.PushFront(9) //往链表头部添加9
	TravsList(lst)

	RingDemo()
}
