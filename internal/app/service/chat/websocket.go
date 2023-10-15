package chat

import (
	"github.com/gin-gonic/gin"
	"grapefruit/kit/log"
)

type Websocket struct {
}

func (w *Websocket) RunAsClient() {

}

func (w *Websocket) RunAsServer() {
	engine := gin.Default()
	err := engine.Run(":8082")
	if err != nil {
		log.Error("[RunAsServer]err:%v", err.Error())
	}
}

func (w *Websocket) Start() {

}

func (w *Websocket) Stop() {

}
