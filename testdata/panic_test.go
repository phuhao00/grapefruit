package testdata

import (
	"fmt"
	"sync"
	"testing"
)

func recoverFromPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered:", r)
	}
}

func doSomething() {

	fmt.Println("Doing something...")
	f(1, 0)
	//panic("something went wrong")
	fmt.Println("This will not be executed")
}

func TestPanic(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer recoverFromPanic()
		defer wg.Done()
		doSomething()
		fmt.Println("Program continues after panic")
	}()

	wg.Wait()
	//
	go func() {}()

}

func f(a, b int) {
	fmt.Println(a / b)
}
