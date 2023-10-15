package main

import (
	"fmt"
	"grapefruit/internal/adapter/rpc"
	"grapefruit/internal/adapter/rpc/example/processor"
	"grapefruit/internal/adapter/rpc/usecase"
)

func main() {
	if err := rpc.Register(new(processor.MockProcessor)); err != nil {
		fmt.Println("register failed")
		return
	}
	rpcServer := &usecase.Server{}
	rpcServer.Init("127.0.0.1:444")

	rpcServer.Run()
}
