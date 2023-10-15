package tcp

import (
	"context"
	"encoding/binary"
	"fmt"
	"grapefruit/kit/log"
	"net"
	"reflect"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

const timeoutTime = 30 // 连接通过验证的超时时间

type Session struct {
	Conn           net.Conn
	ConnID         int64
	verify         int32
	closed         int32
	stopped        chan bool
	signal         chan interface{}
	lastSignal     chan interface{}
	wgRW           sync.WaitGroup
	msgParser      *BufferPacker
	msgBuffSize    int
	MessageHandler func(context context.Context, packet *Message, sess *Session)
	ctx            context.Context
}

func NewSession(conn *net.TCPConn, msgBuffSize int) (*Session, error) {
	tcpConn := &Session{
		closed:      -1,
		verify:      0,
		stopped:     make(chan bool, 1),
		signal:      make(chan interface{}, 100),
		lastSignal:  make(chan interface{}, 1),
		Conn:        conn,
		msgBuffSize: msgBuffSize,
		msgParser:   newInActionPacker(),
		ctx:         context.Background(),
	}
	// Try to open keepalive for tcp.
	err := conn.SetKeepAlive(true)
	if err != nil {
		return nil, err
	}
	err = conn.SetKeepAlivePeriod(1 * time.Minute)
	if err != nil {
		return nil, err
	}
	// disable Nagle algorithm.
	err = conn.SetNoDelay(true)
	if err != nil {
		return nil, err
	}
	err = conn.SetWriteBuffer(msgBuffSize)
	if err != nil {
		return nil, err
	}
	err = conn.SetReadBuffer(msgBuffSize)
	if err != nil {
		return nil, err
	}
	tcpConn.Verify()
	return tcpConn, nil
}

func (c *Session) Connect() {
	if atomic.CompareAndSwapInt32(&c.closed, -1, 0) {
		c.wgRW.Add(1)
		go c.HandleRead()
		c.wgRW.Add(1)
		go c.HandleWrite()
	}
	timeout := time.NewTimer(time.Second * timeoutTime)
L:
	for {
		select {
		case <-timeout.C:
			if !c.Verified() {
				log.Error("[Connect] 验证超时 ip addr %s", c.RemoteAddr())
				break L
			}
		case <-c.stopped:
			break L
		}
	}
	timeout.Stop()
	c.wgRW.Wait()
	c.OnClose()
}

func (c *Session) HandleRead() {
	defer func() {
		if err := recover(); err != nil {
			log.Error("[HandleRead] panic ", err, "\n", string(debug.Stack()))
		}
	}()
	defer func() {
		log.Info("[HandleRead] c.wgRW.Done")
		c.wgRW.Done()
	}()
	for {
		data, err := c.msgParser.Read(c)
		if err != nil {
			//log.Error("[HandleRead]read message error: %v", err)
			time.Sleep(200 * time.Millisecond)
			continue
		}
		//log.Info("[HandleRead] read err:", err.Error())
		//break
		message, err := c.msgParser.Unpack(data)
		if err != nil {
			log.Error("[HandleRead]:%v", err.Error())
		}
		c.OnMessage(message)
	}
}

func (c *Session) HandleWrite() {
	defer func() {
		if err := recover(); err != nil {
			log.Error("[HandleWrite] panic", err, "\n", string(debug.Stack()))
		}
	}()
	defer c.wgRW.Done()
	for {
		select {
		case signal := <-c.signal: // 普通消息
			data, ok := signal.([]byte)
			if !ok {
				log.Error("write message %v error: msg is not bytes", reflect.TypeOf(signal))
				return
			}
			err := c.msgParser.Write(c, data...)
			if err != nil {
				log.Error("write message %v error: %v", reflect.TypeOf(signal), err)
				return
			}
			log.Info("[HandleWrite] signal had send msg :%v", string(data))
		case signal := <-c.lastSignal: // 最后一个通知消息
			data, ok := signal.([]byte)
			if !ok {
				log.Error("write message %v error: msg is not bytes", reflect.TypeOf(signal))
				return
			}
			err := c.msgParser.Write(c, data...)
			if err != nil {
				log.Error("write message %v error: %v", reflect.TypeOf(signal), err)
				return
			}
			log.Info("[HandleWrite] lastSignal  had send msg :%v", string(data))
			time.Sleep(2 * time.Second)
			return
		case <-c.stopped: // 连接关闭通知
			return
		}
	}
}

func (c *Session) AsyncSend(msgID uint64, msg *Message) bool {
	if c.IsShutdown() {
		log.Error("[AsyncSend] session is close !")
		return false
	}

	data, err := c.msgParser.Pack(msgID, msg)
	if err != nil {
		log.Error("[AsyncSend] Pack msgID:%v and msg to bytes error:%v", msgID, err)
		return false
	}

	if uint32(len(data)) > c.msgParser.maxMsgLen {
		log.Error("[AsyncSend] 发送的消息包体过长 msgID:%v", msgID)
		return false
	}

	err = c.Signal(data)
	if err != nil {
		log.Error("%v", err)
		return false
	}
	log.Info("[AsyncSend] 已经发送数据:%v", string(data))

	return true
}

func (c *Session) AsyncSendRowMsg(data []byte) bool {
	if c.IsShutdown() {
		return false
	}

	if uint32(len(data)) > c.msgParser.maxMsgLen {
		log.Error("[AsyncSendRowMsg] 发送的消息包体过长 AsyncSendRowMsg")
		return false
	}

	err := c.Signal(data)
	if err != nil {
		log.Error("%v", err)
		return false
	}

	return true
}

// AsyncSendLastPacket 缓存在发送队列里等待发送goroutine取出 (发送最后一个消息 发送会关闭tcp连接 终止tcp goroutine)
func (c *Session) AsyncSendLastPacket(msgID uint64, msg *Message) bool {
	data, err := c.msgParser.Pack(msgID, msg)
	if err != nil {
		log.Error("[AsyncSendLastPacket] Pack msgID:%v and msg to bytes error:%v", msgID, err)
		return false
	}

	if uint32(len(data)) > c.msgParser.maxMsgLen {
		log.Error("[AsyncSendLastPacket] 发送的消息包体过长 msgID:%v", msgID)
		return false
	}

	err = c.LastSignal(data)
	if err != nil {
		log.Error("%v", err)
		return false
	}

	return true
}

func (c *Session) Signal(signal []byte) error {
	select {
	case c.signal <- signal:
		return nil
	default:
		{
			cmd := binary.LittleEndian.Uint16(signal[2:4])
			return fmt.Errorf("[Signal] buffer full blocking connID:%v cmd:%v", c.ConnID, cmd)
		}
	}
}

func (c *Session) LastSignal(signal []byte) error {
	select {
	case c.lastSignal <- signal:
		return nil
	default:
		{
			cmd := binary.LittleEndian.Uint16(signal[2:4])
			return fmt.Errorf("[LastSignal] buffer full blocking connID:%v cmd:%v", c.ConnID, cmd)
		}
	}
}

func (c *Session) Verified() bool {
	return atomic.LoadInt32(&c.verify) != 0
}

func (c *Session) Verify() {
	atomic.CompareAndSwapInt32(&c.verify, 0, 1)
}

func (c *Session) IsClosed() bool {
	return atomic.LoadInt32(&c.closed) != 0
}

func (c *Session) IsShutdown() bool {
	return atomic.LoadInt32(&c.closed) == 1
}

func (c *Session) Close() {
	if atomic.CompareAndSwapInt32(&c.closed, 0, 1) {
		log.Info("[Session.Close] ")
	}
}

func (c *Session) Read(b []byte) (int, error) {
	return c.Conn.Read(b)
}

func (c *Session) Write(b []byte) (int, error) {
	return c.Conn.Write(b)
}

func (c *Session) LocalAddr() net.Addr {
	return c.Conn.LocalAddr()
}

func (c *Session) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Session) Reset() {
	if atomic.LoadInt32(&c.closed) == -1 {
		return
	}
	c.closed = -1
	c.verify = 0
	c.stopped = make(chan bool, 1)
	c.signal = make(chan interface{}, c.msgBuffSize)
	c.lastSignal = make(chan interface{}, 1)
	c.msgParser.reset()
}

// OnConnect ...
func (c *Session) OnConnect() {
	log.Error("[OnConnect] 建立连接 local:%s remote:%s", c.LocalAddr(), c.RemoteAddr())
}

func (c *Session) OnClose() {
	log.Error("[OnConnect] 断开连接 local:%s remote:%s", c.LocalAddr(), c.RemoteAddr())
}

func (c *Session) OnMessage(message *Message) {
	c.MessageHandler(context.Background(), message, c)
}
