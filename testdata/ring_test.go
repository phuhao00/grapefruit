package testdata

import (
	"container/ring"
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	bufferSize := 5
	buffer := ring.New(bufferSize)

	// 模拟向缓冲区写入数据
	for i := 1; i <= bufferSize+2; i++ {
		buffer.Value = i
		buffer = buffer.Next()
	}

	// 遍历缓冲区并打印数据
	buffer.Do(func(x interface{}) {
		fmt.Println(x)
	})
}

func TestIterate(t *testing.T) {
	data := []string{"apple", "banana", "cherry", "date", "elderberry"}
	r := ring.New(len(data))

	// 将数据添加到ring中
	for _, d := range data {
		r.Value = d
		r = r.Next()
	}

	// 使用迭代器遍历ring并打印数据
	iterator := r
	iterator.Do(func(x interface{}) {
		fmt.Println(x)
		iterator = iterator.Next()
	})

}
