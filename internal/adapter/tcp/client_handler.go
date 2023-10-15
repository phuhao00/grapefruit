package tcp

import (
	"context"
	"fmt"
	"time"
)

type ClientMessageHandler func(ctx context.Context, packet *Message, sess *Session)

func ClientHandleMessage(ctx context.Context, packet *Message, sess *Session) {
	if packet.Command == 1 {
		fmt.Println("hello world1")
		time.Sleep(time.Second)
		sess.AsyncSend(2, &Message{
			Command: 2,
			Data:    []byte("client"),
		})
	}
	if packet.Command == 2 {
		fmt.Println("hello world2")

	}
}
