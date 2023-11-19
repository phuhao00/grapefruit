package testdata

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	var mu sync.Mutex
	counter := 0

	// 加锁保护共享资源
	mu.Lock()
	counter++
	mu.Unlock()

	fmt.Println("Counter:", counter)
}

func TestRwLock(t *testing.T) {
	var mu sync.RWMutex
	counter := 0

	// 写操作
	go func() {
		mu.Lock()
		counter++
		time.Sleep(1 * time.Second)
		mu.Unlock()
	}()

	// 读操作
	go func() {
		mu.RLock()
		fmt.Println("Counter:", counter)
		time.Sleep(500 * time.Millisecond)
		mu.RUnlock()
	}()
	// 读操作
	go func() {
		mu.RLock()
		fmt.Println("Counter:", counter)
		time.Sleep(500 * time.Millisecond)
		mu.RUnlock()
	}()

	time.Sleep(2 * time.Second)
}

func TestSyncOnceDo(t *testing.T) {
	var once sync.Once

	for i := 0; i < 5; i++ {
		once.Do(func() {
			fmt.Println("Only once")
		})
	}

}

func TestMap(t *testing.T) {
	//mDemo := make(map[string]int64, 7)
	smDemo := sync.Map{}
	smDemo.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
	smDemo.Store("1", 1)
	value, ok := smDemo.Load("1")
	if ok {
		fmt.Println(value.(int))
	}
	smDemo.Delete("1")
	_, o := smDemo.Load("1")
	if !o {
		fmt.Println("not exist 1")
	}
	smDemo.LoadOrStore("3", 3.0)
	load, b := smDemo.Load("3")
	if b {
		fmt.Println(load.(float64))
	}

}
