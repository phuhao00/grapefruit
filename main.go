package main

import (
	"context"
	"grapefruit/config"
	"grapefruit/internal/app/server"
)

func main() {
	_, err := config.LoadConfig(context.Background())
	if err != nil {
		panic(err)
	}
	newServer := server.NewServer()
	newServer.Run()
}
