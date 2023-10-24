package websocket

import (
	"crypto/tls"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"testing"
)

func TestGetAIResponse(t *testing.T) {
	u := url.URL{Scheme: "wss", Host: "127.0.0.1", Path: "/chat_gpt"}
	log.Printf("clientconnecting to %s", u.String())
	websocket.DefaultDialer.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true, // 跳过证书验证（仅供示例使用，实际应用中应谨慎处理）
	}
	header := http.Header{}
	header.Add("Authorization", "token")
	c, _, err := websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		log.Fatal("dial server:", err)
	}
	defer c.Close()
	done := make(chan struct{})
	bytes, err := json.Marshal(struct {
	}{})
	err = c.WriteMessage(websocket.BinaryMessage, bytes)
	if err != nil {
		return
	}
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("client read errwrap:", err)
				return
			}
			log.Printf("client recv msg: %s", message)
		}
	}()
	select {}
}
