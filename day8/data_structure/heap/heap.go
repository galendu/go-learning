package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Value    string
	priority int //优先级,数字越大,优先级越高
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority //golang默认提供的时小根堆,而优先队列是大跟堆,所以这里要反着定义Less
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// 往slice里append,需要传slice指针
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

// 让slice指向新的子切片,需要传slice指针
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]   //数组最后一个元素
	*pq = old[0 : n-1] //去掉最后一个元素
	return item
}

func testPriorityQueue() {
	pq := make(PriorityQueue, 0, 10)
	pq.Push(&Item{"A", 3})
	pq.Push(&Item{"B", 2})
	pq.Push(&Item{"C", 4})
	heap.Init(&pq)
	heap.Push(&pq, &Item{"D", 6}) //通过heap添加元素
	for pq.Len() > 0 {
		fmt.Println(heap.Pop(&pq)) //通过heap删除堆顶元素
	}
}
func main() {
	// buildHeap()
	testPriorityQueue()
	// testTimeoutCache()

}
