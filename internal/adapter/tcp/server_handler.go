package tcp

import (
	"context"
	"fmt"
)

func ServerHandleMessage(context context.Context, packet *Message, session *Session) {
	if packet.Command == 1 {
		fmt.Println("hi 1", string(packet.Data))
		session.AsyncSend(
			1,
			&Message{
				Command: 2,
				Data:    []byte("server"),
			},
		)
	}
}
