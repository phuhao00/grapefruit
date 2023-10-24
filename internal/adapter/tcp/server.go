package tcp

import (
	"context"
	"grapefruit/kit/log"
	"net"
	"os"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

type Server struct {
	pid  int64
	Addr string

	MaxConnNum   int
	ln           *net.TCPListener
	connSet      map[net.Conn]interface{}
	counter      int64
	idCounter    int64
	mutexConn    sync.Mutex
	wgLn         sync.WaitGroup
	wgConn       sync.WaitGroup
	connBuffSize int
	ctx          context.Context
}

func NewServer(addr string, maxConnNum int, buffSize int) *Server {
	s := &Server{
		Addr:         addr,
		MaxConnNum:   maxConnNum,
		connBuffSize: buffSize,
	}
	s.Init()
	return s
}

func (s *Server) Init() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", s.Addr)

	if err != nil {
		log.Error("[Init] addr resolve error", tcpAddr, err)
	}

	ln, err := net.ListenTCP("tcp4", tcpAddr)

	if err != nil {
		log.Fatal("%v", err)
	}

	if s.MaxConnNum <= 0 {
		s.MaxConnNum = 100
		log.Info("[Init] invalid MaxConnNum, reset to %v", s.MaxConnNum)
	}

	s.ln = ln
	s.connSet = make(map[net.Conn]interface{})
	s.counter = 1
	s.idCounter = 1
	s.pid = int64(os.Getpid())
	log.Info("[Init] Server Listen %s", s.ln.Addr().String())
}

func (s *Server) Run() {
	defer func() {
		if err := recover(); err != nil {
			log.Error("[Run] panic", err, "\n", string(debug.Stack()))
		}
	}()

	s.wgLn.Add(1)
	defer s.wgLn.Done()

	var tempDelay time.Duration
	for {
		conn, err := s.ln.AcceptTCP()

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
				log.Info("[Run]accept error: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return
		}
		tempDelay = 0

		if atomic.LoadInt64(&s.counter) >= int64(s.MaxConnNum) {
			err = conn.Close()
			if err != nil {
				log.Error("[Run] errwrap:%v", err.Error())
			}
			log.Info("[Run] too many connections %v", atomic.LoadInt64(&s.counter))
			continue
		}
		tcpSession, err := NewSession(conn, s.connBuffSize)
		if err != nil {
			log.Error("[Run] errwrap:%v", err)
			return
		}
		tcpSession.MessageHandler = ServerHandleMessage
		s.addConn(conn, tcpSession)
		s.wgConn.Add(1)
		go func() {
			tcpSession.Connect()
			s.removeConn(conn, tcpSession)
			s.wgConn.Done()
		}()
	}
}

func (s *Server) Close() {
	err := s.ln.Close()
	if err != nil {
		log.Error(err.Error())
	}
	s.wgLn.Wait()

	s.mutexConn.Lock()
	for conn := range s.connSet {
		err = conn.Close()
		if err != nil {
			log.Error(err.Error())
		}
	}
	s.connSet = nil
	s.mutexConn.Unlock()
	s.wgConn.Wait()
}

func (s *Server) addConn(conn net.Conn, tcpSession *Session) {
	s.mutexConn.Lock()
	atomic.AddInt64(&s.counter, 1)
	s.connSet[conn] = conn
	nowTime := time.Now().Unix()
	idCounter := atomic.AddInt64(&s.idCounter, 1)
	connId := (nowTime << 32) | (s.pid << 24) | idCounter
	tcpSession.ConnID = connId
	s.mutexConn.Unlock()
	tcpSession.OnConnect()
}

func (s *Server) removeConn(conn net.Conn, tcpConn *Session) {
	log.Error("removeConn---")
	s.mutexConn.Lock()
	atomic.AddInt64(&s.counter, -1)
	delete(s.connSet, conn)
	s.mutexConn.Unlock()
}
