package testdata

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	DeferDemo()
}

func DeferDemo() (i int) {
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(2)

	}()
	defer func() {
		fmt.Println(3)
	}()
	fmt.Println("return")
	defer func() {
		fmt.Println("-----", i)
	}()
	fmt.Println("=====", i)
	return 1

}
