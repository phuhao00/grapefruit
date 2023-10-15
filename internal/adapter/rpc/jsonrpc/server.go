// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonrpc

import (
	"encoding/json"
	"errors"
	"grapefruit/internal/adapter/rpc"
	"io"
	"sync"
)

var errMissingParams = errors.New("jsonrpc: request body missing params")

type ServerCodec struct {
	dec     *json.Decoder // for reading JSON values
	enc     *json.Encoder // for writing JSON values
	c       io.Closer
	Req     serverRequest
	mutex   sync.Mutex // protects Seq, Pendin
	Seq     uint64
	Pending map[uint64]*json.RawMessage
}

// NewServerCodec returns a new rpc.ServerCodec using JSON-RPC on conn.
func NewServerCodec(conn io.ReadWriteCloser) rpc.ServerCodec {
	return &ServerCodec{
		dec:     json.NewDecoder(conn),
		enc:     json.NewEncoder(conn),
		c:       conn,
		Pending: make(map[uint64]*json.RawMessage),
	}
}

type serverRequest struct {
	Method string           `json:"method"`
	Params *json.RawMessage `json:"params"`
	Id     *json.RawMessage `json:"id"`
}

func (r *serverRequest) reset() {
	r.Method = ""
	r.Params = nil
	r.Id = nil
}

type serverResponse struct {
	Id     *json.RawMessage `json:"id"`
	Result any              `json:"result"`
	Error  any              `json:"error"`
}

func (c *ServerCodec) ReadRequestHeader(r *rpc.Request) error {
	c.Req.reset()
	if err := c.dec.Decode(&c.Req); err != nil {
		return err
	}
	r.ServiceMethod = c.Req.Method
	c.mutex.Lock()
	c.Seq++
	c.Pending[c.Seq] = c.Req.Id
	c.Req.Id = nil
	r.Seq = c.Seq
	c.mutex.Unlock()

	return nil
}

func (c *ServerCodec) ReadRequestBody(x any) error {
	if x == nil {
		return nil
	}
	if c.Req.Params == nil {
		return errMissingParams
	}
	var params [1]any
	params[0] = x
	return json.Unmarshal(*c.Req.Params, &params)
}

var null = json.RawMessage([]byte("null"))

func (c *ServerCodec) WriteResponse(r *rpc.Response, x any) error {
	c.mutex.Lock()
	b, ok := c.Pending[r.Seq]
	if !ok {
		c.mutex.Unlock()
		return errors.New("invalid sequence number in response")
	}
	c.mutex.Unlock()

	if b == nil {
		// Invalid request so no id. Use JSON null.
		b = &null
	}
	resp := serverResponse{Id: b}
	if r.Error == "" {
		resp.Result = x
	} else {
		resp.Error = r.Error
	}
	return c.enc.Encode(resp)
}

func (c *ServerCodec) Close() error {
	return c.c.Close()
}

func ServeConn(conn io.ReadWriteCloser) {
	rpc.ServeCodec(NewServerCodec(conn))
}

func (c *ServerCodec) Lock() {
	c.mutex.Lock()

}

func (c *ServerCodec) UnLock() {
	c.mutex.Unlock()
}

func (c *ServerCodec) DeleteSeq(seq uint64) {
	c.mutex.Lock()
	delete(c.Pending, seq)
	c.mutex.Unlock()
}
