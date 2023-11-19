package testdata

import (
	"fmt"
	"testing"
	"time"
)

func TestChanel(t *testing.T) {
	//ch := make(chan int)
	//ch <- 1
	//data, ok := <-ch
	//if ok {
	//	fmt.Println(data)
	//}

	done := make(chan bool)
	fmt.Println("1")
	go func() {
		fmt.Println("3")
		// 执行一些操作
		fmt.Println("hello")
		done <- true // 操作完成后向通道发送信号
	}()
	<-done
	//todo
	fmt.Println("2")
}

func TestEditWholeVar(t *testing.T) {
	var counter int
	ch := make(chan bool)
	go func() {
		fmt.Println("---", counter)
		counter++
		// 并发修改counter
		ch <- true
	}()
	go func() {
		fmt.Println("====", counter)
		counter++
		// 并发修改counter
		ch <- true
	}()
	<-ch // 等待两个Goroutine完成
	fmt.Println(counter)
	<-ch
	fmt.Println(counter)

}

func readOnly(ch <-chan int) {
	data := <-ch // 从只读通道读取数据
	fmt.Println(data)
}

func writeOnly(ch chan<- int) {
	ch <- 10 // 向只写通道写入数据
}
func TestReadWirte(t *testing.T) {

	ch := make(chan int)

	go readOnly(ch) // 传递只读通道给函数

	go func() {
		writeOnly(ch) // 传递只写通道给函数
	}()

	time.Sleep(time.Second)
}

func worker1(input chan int, output chan int) {
	for num := range input {
		result := num * 2
		output <- result
	}
	//defer func() {
	//	close(output)
	//
	//}()
}

func TestChannel1(t *testing.T) {
	input := make(chan int)
	output := make(chan int)

	go worker1(input, output)
	go func() {
		for result := range output {
			fmt.Println(result)
		}
	}()
	// 发送数据到输入通道
	for i := 1; i <= 10; i++ {
		input <- i
	}
	//defer func() {
	//	close(input)
	//
	//}()

	// 从输出通道接收结果

}
