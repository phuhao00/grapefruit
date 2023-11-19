package mock

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

// 非泛型版本的切片过滤函数
func filterNonGeneric(numbers []int, f func(int) bool) []int {
	var result []int
	for _, num := range numbers {
		if f(num) {
			result = append(result, num)
		}
	}
	return result
}

// 泛型版本的切片过滤函数
func filterGeneric[T any](items []T, f func(T) bool) []T {
	var result []T
	for _, item := range items {
		if f(item) {
			result = append(result, item)
		}
	}
	return result
}
func TestGen(t *testing.T) {
	// 生成一个较大的整数切片
	var numbers []int
	for i := 0; i < 1000000; i++ {
		numbers = append(numbers, i)
	}

	// 测试非泛型版本的性能
	start := time.Now()
	filterNonGeneric(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Non-generic filter took:", time.Since(start))

	// 测试泛型版本的性能
	start = time.Now()
	filterGeneric(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Generic filter took:", time.Since(start))

	sort.Slice(numbers, func(i, j int) bool {
		return rand.Float64() < rand.Float64()
	})
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	// 这里可能还需要比较两个过滤后的切片是否一致，这里只展示了性能测试
}
