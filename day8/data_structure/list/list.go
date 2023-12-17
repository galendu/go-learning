package main

import (
	"container/list"
	"fmt"
)

func TraversList(lst *list.List) { //遍历列表
	head := lst.Front()
	for head.Next() != nil {
		fmt.Printf("%v ", head.Value)
		head = head.Next()
	}
	fmt.Println(head.Value)
}

func ReverseList(lst *list.List) {
	tail := lst.Back() //返回列表的最后一个元素
	for tail.Prev() != nil {
		fmt.Printf("%v ", tail.Value)
		tail = tail.Prev()
	}
	fmt.Println(tail.Value)
}

func testList() {
	lst := list.New()
	lst.PushBack(1)
	lst.PushBack(2)
	lst.PushBack(3)
	lst.PushFront(4)
	lst.PushFront(5)
	lst.PushFront(6)
	TraversList(lst)
	ReverseList(lst)

	_ = lst.Remove(lst.Back())  //移除元素的同时返回元素的值,注意元素不能是nil
	_ = lst.Remove(lst.Front()) //Remove操作复杂的为O(1)
	fmt.Printf("length %d\n", lst.Len())
	TraversList(lst)
}

func testDoubleList() {
	list := new(DoubleList)
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Traverse()
	node := list.Get(3)
	list.InsertAfter(9, node)
	list.Traverse()

}
func main() {
	// testList()
	testDoubleList()
}
