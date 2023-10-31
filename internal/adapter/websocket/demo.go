package websocket

import (
	"github.com/gorilla/websocket"
	"grapefruit/kit/log"
	"time"
)

var demoMsgCh chan []byte

func (s *Websocket) Demo() {
	go s.demoReadMessage()
	s.demoSendClientMsg()
}

func (s *Websocket) demoReadMessage() {
	tk := time.NewTicker(time.Second)
	demoMsgCh = make(chan []byte)
	for {
		select {
		case _, ok := <-s.CloseCh:
			if !ok {
				break
			}
		case <-tk.C:
			demoMsgCh <- []byte("hello chatgpt gpt")
		}
	}
}

func (s *Websocket) demoSendClientMsg() {
	for {
		select {
		case msgC := <-demoMsgCh:
			err := s.socket.WriteMessage(websocket.BinaryMessage, msgC)
			if err != nil {
				s.Close()
				log.Error(err.Error())
				s.CloseCh <- struct{}{}
			}
		}
	}
}
