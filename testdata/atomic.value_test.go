package testdata

import (
	"fmt"
	"go.uber.org/atomic"
	"sync"
	"testing"
)

func TestRecover(t *testing.T) {
	var a atomic.Int32
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println(err)
			}
			a.Store(1)
		}()
		load := a.Load()
		if load == 0 {
			panic("errr")
		}
		fmt.Println(1111111111111)
		wg.Done()
	}()
	wg.Wait()
}

func AA() {

	//atomic.Value{}

}
