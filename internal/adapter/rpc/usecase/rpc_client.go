package usecase

import (
	"grapefruit/internal/adapter/rpc"
	"grapefruit/internal/adapter/rpc/jsonrpc"
)

type Client struct {
	pool *Pool
	Addr string
}

func NewRpcClient(addr string) *Client {
	rpcClient := &Client{
		pool: &Pool{
			MaxIdle:         10,
			IdleTimeout:     0,
			MaxConnLifetime: 300,
			Dial:            func() (*rpc.Client, error) { return jsonrpc.Dial("tcp", addr) },
		},
		Addr: addr,
	}
	return rpcClient
}

func (c *Client) Call(method string, args interface{}, reply interface{}, replyChan chan any, closeCh chan struct{}) error {
	rpcClient, err := c.pool.Get()
	if err != nil {
		return err
	}
	rpcClient.pc.c.Reply = replyChan
	rpcClient.pc.c.CloseCh = closeCh
	err = rpcClient.Call(method, args, reply)
	if err == rpc.ErrShutdown {
		return err
	}
	rpcClient.Close()
	return err
}
