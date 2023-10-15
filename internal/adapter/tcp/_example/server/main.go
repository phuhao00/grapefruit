package main

import (
	"grapefruit/internal/adapter/tcp"
)

func main() {
	server := tcp.NewServer(":5001", 100, 200)

	go server.Run()
	select {}
}
