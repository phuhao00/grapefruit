package websocket

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/atomic"
	"grapefruit/kit/log"
	"io"
	"net/http"
	"time"
)

type Websocket struct {
	socket   *websocket.Conn
	ctx      *gin.Context
	CloseCh  chan struct{}
	RspMsgCh chan interface{}
	UserId   int64
	IsClose  atomic.Bool
}

var upGrader = websocket.Upgrader{
	HandshakeTimeout: 5 * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewWebsocket(ctx *gin.Context, w http.ResponseWriter, r *http.Request, responseHeader http.Header, userId int64) (*Websocket, error) {
	ws, err := upGrader.Upgrade(w, r, responseHeader)
	if err != nil {
		return nil, err
	}
	return &Websocket{
		socket:   ws,
		ctx:      ctx,
		CloseCh:  make(chan struct{}, 1),
		RspMsgCh: make(chan interface{}),
	}, nil
}

func (s *Websocket) SendMessage(message []byte) error {
	err := s.socket.WriteMessage(websocket.BinaryMessage, message)
	if err != nil {

		s.Close()
		log.Error("socket on send error", err.Error())
	}
	return err
}

func (s *Websocket) OnMessage(msg chan interface{}) {

loop:
	for {
		select {

		case _, ok := <-s.CloseCh:
			if !ok {
				log.Info("[OnMessage] break !")
				break loop
			}
			log.Info("[OnMessage] CloseCh")
		case data := <-msg:
			log.Info("%v", data)
			bytes, err := json.Marshal(struct {
			}{})
			if err != nil {
				log.Error("[ json.Marshal(clientRsp)]:%v", err.Error())
				continue
			}
			log.Info("[OnMessage] msg: %v", string(bytes))
			err = s.SendMessage(bytes)
			if err != nil {
				log.Error("[ s.SendMessag:%v", err.Error())
			}
		}
	}
}

func (s *Websocket) Close() {
	defer func() {
		_, ok := <-s.CloseCh
		if ok {
			close(s.CloseCh)
		}
	}()
	err := s.socket.Close()
	if err != nil {
		log.Error("socket on closed error", err.Error())
		return
	}
	s.IsClose.Store(true)

}

func (s *Websocket) ReadClientMessage(req chan interface{}) {
	for {
		if s.IsClose.Load() {
			break
		}
		_, msg, err := s.socket.ReadMessage()
		if err != nil && err != io.EOF {
			log.Error(err.Error())
			break
		}
		log.Info("[ReadClientMessage]", string(msg))

		err = json.Unmarshal(msg, struct{}{})
		if err != nil {
			log.Error(err.Error())
			continue
		}
	}
}

func (s *Websocket) DemoTest(req chan interface{}) {
	tk := time.NewTicker(time.Second)
	for {
		select {
		case <-tk.C:
		}
	}
}
