package main

import (
	"bufio"
	"fmt"
	"grapefruit/internal/adapter/tcp"
	"os"
	"strings"
)

var aiRequest chan interface{}

func main() {

	client := tcp.NewClient(":5001", 200, tcp.WithClientMsgHandler(tcp.ClientHandleMessage))
	client.Run()
	aiRequest = make(chan interface{}, 2)
	go Tick(client)
	go GetCommandInput()
	select {}
}

func Tick(c *tcp.Client) {
	for {
		select {
		//case data := <-aiRequest:
		//	if c.Session != nil {
		//		c.Session.AsyncSend(1, &tcp.Message{
		//			Command: 1,
		//			Data:    []byte("hiiiiiii"),
		//		})
		//	}
		}
	}
}

func GetCommandInput() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')

		// 去掉空格和换行符
		text = text[:len(text)-1]
		text = strings.TrimSpace(text)

		//// 执行输入的代码
		//cmd := exec.Command("/bin/sh", "-c", text)
		//cmd.Stdout = os.Stdout
		//cmd.Stderr = os.Stderr
		//cmd.Run()

		fmt.Println()
	}
}
