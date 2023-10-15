package main

import (
	"fmt"
	"grapefruit/internal/adapter/rpc/example/processor"
	"grapefruit/internal/adapter/rpc/usecase"
)

func main() {
	chRply := make(chan any, 100)
	closeCh := make(chan struct{})
	c := usecase.NewRpcClient("127.0.0.1:444")
	req := &processor.MockParam{Tag: "123"}
	rsp := &processor.MockParam{Tag: ""}
	err := c.Call("MockProcessor.Print2", req, rsp, chRply, closeCh)
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case data := <-chRply:
			fmt.Println(data)
		case <-closeCh:
			fmt.Println("client close !!!")
			break
		}
	}
}
