package testdata

import (
	"fmt"
	"testing"
)

type AB interface {
	AA()
	AC
}

type AC interface {
	BB()
}

type CC struct {
	num int
}

func (C CC) AA() {
	fmt.Println("bb")
	C.num = 1
}

func (C CC) BB() {
	fmt.Println("cc")
}

func TestCC(t *testing.T) {
	c := CC{}
	c.AA()
	c.BB()
}
