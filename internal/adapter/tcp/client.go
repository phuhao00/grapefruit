package tcp

import (
	"context"
	"grapefruit/kit/log"
	"net"
	"sync/atomic"
	"time"
)

type Client struct {
	Address               string
	ctx                   context.Context
	bufferSize            int
	running               atomic.Value
	Session               *Session
	SessionMessageHandler ClientMessageHandler
}

func NewClient(address string, connBuffSize int, options ...Option) *Client {
	client := &Client{
		bufferSize: connBuffSize,
		Address:    address,
	}
	client.running.Store(false)
	for _, o := range options {
		o(client)
	}
	return client
}

type Option func(*Client)

func WithClientMsgHandler(handler ClientMessageHandler) Option {
	return func(client *Client) {
		client.SessionMessageHandler = handler
	}
}

func (c *Client) Dial() (*net.TCPConn, error) {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", c.Address)

	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp4", nil, tcpAddr)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *Client) Run() {
	conn, err := c.Dial()
	if err != nil {
		log.Error("[Run] %v", err)
		time.Sleep(100 * time.Millisecond)
		if conn != nil {
			err = conn.Close()
			if err != nil {
				log.Info("[conn.Close]")
			}
		}
		return
	}
	c.RunSession(conn)
}

// Daemon 守护程序
func (c *Client) Daemon() {
	tk := time.NewTicker(time.Second)
	for {
		select {
		case <-tk.C:
			if c.Session.IsShutdown() {
				c.Run()
			}
		}
	}
}

func (c *Client) RunSession(conn *net.TCPConn) {
	sess, err := NewSession(conn, c.bufferSize)
	if err != nil {
		log.Error("[Run] %v", err)
	}
	sess.MessageHandler = c.SessionMessageHandler
	c.Session = sess
	log.Info("[Run] client dial successfully ")
	sess.Connect()
}

func (c *Client) Close() {
	c.running.CompareAndSwap(false, true)
}
