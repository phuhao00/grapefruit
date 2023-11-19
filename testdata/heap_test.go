package testdata

import (
	"container/heap"
	"fmt"
	"testing"
)

// 自定义堆类型
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestHeap(t *testing.T) {

	// 创建一个空堆
	var h IntHeap

	// 添加元素到堆中
	heap.Push(&h, 5)
	heap.Push(&h, 2)
	heap.Push(&h, 10)
	heap.Push(&h, 8)

	// 弹出堆顶元素
	fmt.Println(heap.Pop(&h)) // Output: 2

	// 遍历堆中剩余的元素
	for h.Len() > 0 {
		fmt.Println(heap.Pop(&h))
	}
}
