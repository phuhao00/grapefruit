package testdata

import (
	"fmt"
	"testing"
	"time"
)

func TestArray(t *testing.T) {

	// 示例1：固定长度的数组表示坐标点
	var point [2]int // 2维坐标点的数组，x和y坐标
	point[0] = 10    // x坐标
	point[1] = 20    // y坐标
	fmt.Println("Point coordinates:", point)

	// 示例2：存储固定数量的相同类型元素
	var weekdays [7]string // 存储一周的工作日
	weekdays = [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println("Weekdays:", weekdays)

	// 示例3：数组作为函数参数的值传递
	numbers := [3]int{1, 2, 3}
	doubleArray(numbers)
	fmt.Println("Original array:", numbers)

	// 示例4：性能考虑 - 比较数组和切片的迭代速度
	const size = 1000000
	var array [size]int
	slice := make([]int, size)

	// 使用数组进行迭代
	start := time.Now()
	for i := 0; i < size; i++ {
		array[i] = i
	}
	elapsedArray := time.Since(start)
	fmt.Println("Time taken by array:", elapsedArray)

	// 使用切片进行迭代
	start = time.Now()
	for i := 0; i < size; i++ {
		slice[i] = i
	}
	elapsedSlice := time.Since(start)
	fmt.Println("Time taken by slice:", elapsedSlice)
}

// 示例3函数：传入数组的副本
func doubleArray(arr [3]int) {
	for i := range arr {
		arr[i] *= 2
	}
	fmt.Println("Doubled array:", arr)
}
