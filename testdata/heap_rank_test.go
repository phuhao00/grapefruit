package testdata

import (
	"container/heap"
	"fmt"
	"testing"
)

// Item 是堆中的元素，包含了值和优先级
type Item struct {
	value    string
	priority int
}

// PriorityQueue 实现了堆的接口
type PriorityQueue []*Item

// Len 返回堆的长度
func (pq PriorityQueue) Len() int { return len(pq) }

// Less 比较两个元素的优先级
func (pq PriorityQueue) Less(i, j int) bool {
	// 优先级高的在堆顶
	return pq[i].priority > pq[j].priority
}

// Swap 交换元素位置
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 将元素推入堆中
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

// Pop 从堆中取出元素
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// 更新堆中的元素
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.priority)
}

func TestHeapRank(t *testing.T) {

	// 创建一个优先队列并初始化
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// 添加一些元素到堆中
	pq.Push(&Item{"Player 1", 1})
	pq.Push(&Item{"Player 2", 2})
	pq.Push(&Item{"Player 3", 3})
	pq.Push(&Item{"Player 4", 2})

	// 更新堆中的元素
	item := pq[0]
	pq.update(item, item.value, 3)

	// 取出优先级最高的元素
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%s: %d\n", item.value, item.priority)
	}
}
