package testdata

import (
	"fmt"
	"sync"
	"testing"
)

func TestGoUse(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan bool)
	go func() {
		fmt.Println(666)
		wg.Done()
		ch <- false

	}()
	go Abbb(&wg, ch)
	wg.Wait()
}

func Abbb(group *sync.WaitGroup, ch chan bool) {
	fmt.Println(7777)
	group.Done()
	select {
	case data := <-ch:
		fmt.Println(data)
	}

}
