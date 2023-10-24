package usecase

import (
	"context"
	"grapefruit/internal/adapter/rpc/jsonrpc"
	"grapefruit/kit/log"
	"net"
	"runtime/debug"
	"sync"
	"time"
)

type Server struct {
	Addr string
	ln   *net.TCPListener
	wgLn sync.WaitGroup
	cxt  context.Context
}

func (srv *Server) Init(addr string) {
	srv.cxt = context.Background()
	srv.Addr = addr

	tcpAddr, err := net.ResolveTCPAddr("tcp4", srv.Addr)

	if err != nil {
		log.Error("[net] addr resolve error", tcpAddr, err)
		return
	}

	ln, err := net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		log.Error("%v", err)
		return
	}

	srv.ln = ln
	log.Info("RpcServer Listen %s", srv.ln.Addr().String())
}

func (srv *Server) Run() {
	// 捕获异常
	defer func() {
		if err := recover(); err != nil {
			log.Error("[net] panic", err, "\n", string(debug.Stack()))
		}
	}()

	srv.wgLn.Add(1)
	defer srv.wgLn.Done()

	var tempDelay time.Duration
	for {
		conn, err := srv.ln.AcceptTCP()

		if err != nil {
			if _, ok := err.(net.Error); ok {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				log.Info("accept error: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return
		}
		tempDelay = 0

		// Try to open keepalive for tcp.
		err = conn.SetKeepAlive(true)
		if err != nil {
			log.Error("errwrap:%v", err.Error())
		}
		err = conn.SetKeepAlivePeriod(1 * time.Minute)
		if err != nil {
			log.Error("errwrap:%v", err.Error())
		}
		// disable Nagle algorithm.
		err = conn.SetNoDelay(true)
		if err != nil {
			log.Error("errwrap:%v", err.Error())
		}
		err = conn.SetWriteBuffer(128 * 1024)
		if err != nil {
			log.Error("errwrap:%v", err.Error())
		}
		err = conn.SetReadBuffer(128 * 1024)
		if err != nil {
			log.Error("errwrap:%v", err.Error())
		}
		go jsonrpc.ServeConn(conn)

		log.Debug("accept a rpc conn")
	}
}

func (srv *Server) Close() {
	srv.ln.Close()
	srv.wgLn.Wait()
}
