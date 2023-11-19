package testdata

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func worker2(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 在函数结束时通知 WaitGroup 任务完成

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // 模拟工作时间
	fmt.Printf("Worker %d done\n", id)
}

func TestWaitGroup(t *testing.T) {
	runtime.GOMAXPROCS(4)

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // 添加一个任务到 WaitGroup

		go worker2(i, &wg)
	}

	wg.Wait() // 等待所有任务完成

	fmt.Println("All workers have finished")
}
