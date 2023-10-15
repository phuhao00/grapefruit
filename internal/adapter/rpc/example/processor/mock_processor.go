package processor

import (
	"fmt"
	"grapefruit/internal/adapter/rpc"
	"time"
)

type MockProcessor int

type MockParam struct {
	Tag string
}

func (p MockProcessor) Print(req, rsp *string) error {
	fmt.Println(req)
	tmp := "hi,world"
	rsp = &tmp
	return nil
}

var count = 0

func (p MockProcessor) Print2(req, rsp *MockParam) error {
	//todo 队列 拿数据
	fmt.Println(req)
	if count < 20 {
		time.Sleep(1)
		count++
		rsp.Tag = "abc"
		fmt.Println("count:", count)
		return nil
	}
	count = 0
	return rpc.ResolveFinish
}
