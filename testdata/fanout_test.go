package testdata

import (
	"fmt"
	"sync"
	"testing"
)

func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range in {
		fmt.Println("Received:", val)
	}
}

func fanOut(in <-chan int, n int) []chan int {
	chs := make([]chan int, n)
	for i := 0; i < n; i++ {
		chs[i] = make(chan int)
		go func(index int) {
			for val := range in {
				chs[index] <- val
			}
			close(chs[index])
		}(i)
	}
	return chs
}

func TestFanOut(t *testing.T) {

	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go producer(ch, &wg)

	consumers := fanOut(ch, 3)
	wg.Add(len(consumers))

	for _, c := range consumers {
		go consumer(c, &wg)
	}

	wg.Wait()
}
