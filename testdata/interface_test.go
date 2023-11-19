package testdata

import (
	"fmt"
	"testing"
)

type A interface {
	Aa()
}

type Any interface {
	Push(int2 int)
	Pop() int
	A
}

type Stack[T Any] struct {
	elements []T
}

func (s Stack[T]) Push(int2 int) {
	//TODO implement me
	panic("implement me")
}

func (s Stack[T]) Pop() int {
	//TODO implement me
	panic("implement me")
}

//func (s Stack[T]) Aa() {
//
//}

func TestStack(t *testing.T) {
	stack := Stack[Any]{}
	stack.Push(10)
	stack.Push(20)
	fmt.Println(stack.Pop()) // 输出: 20
}
