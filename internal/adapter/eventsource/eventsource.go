package eventsource

import (
	"bytes"
	"context"
	"fmt"
	"grapefruit/kit/log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type EventMessage struct {
	id     string
	userId int64
	event  string
	data   string
}

func NewEventMessage(id string, userId int64, event string, data string) *EventMessage {
	return &EventMessage{
		id:     id,
		userId: userId,
		event:  event,
		data:   data,
	}

}

type RetryMessage struct {
	retry  time.Duration
	userId int64
	data   string
}

type eventSource struct {
	customHeadersFunc func(*http.Request) [][]byte

	sink           chan Message
	staled         chan *consumer
	add            chan *consumer
	close          chan struct{}
	idleTimeout    time.Duration
	retry          time.Duration
	timeout        time.Duration
	closeOnTimeout bool
	gzip           bool

	consumers sync.Map //userId->consumer
	ctx       context.Context
}

type Settings struct {
	// SetTimeout sets the write timeout for individual messages. The
	// default is 2 seconds.
	Timeout time.Duration

	// CloseOnTimeout sets whether a write timeout should close the
	// connection or just drop the Message.
	//
	// If the connection gets closed on a timeout, it's the client's
	// responsibility to re-establish a connection. If the connection
	// doesn't get closed, messages might get sent to a potentially dead
	// client.
	//
	// The default is true.
	CloseOnTimeout bool

	// Sets the timeout for an idle connection. The default is 30 minutes.
	IdleTimeout time.Duration

	// Gzip sets whether to use gzip Content-Encoding for clients which
	// support it.
	//
	// The default is false.
	Gzip bool
}

func DefaultSettings() *Settings {
	return &Settings{
		Timeout:        2 * time.Second,
		CloseOnTimeout: true,
		IdleTimeout:    30 * time.Minute,
		Gzip:           false,
	}
}

type EventSource interface {
	http.Handler
	//SendEventMessage send Message to  consumer
	SendEventMessage(userId int64, data, event, id string)

	//SendRetryMessage send retry Message to  consumer
	SendRetryMessage(userId int64, duration time.Duration)

	//Close close all consumer
	Close()
}

type Message interface {
	//PrepareMessage The Message to be sent to clients
	PrepareMessage() []byte
	//GetUserId get userId
	GetUserId() int64
	//GetData get data
	GetData() string
}

func (m *EventMessage) GetData() string {
	return m.data
}

func (m *EventMessage) GetUserId() int64 {
	return m.userId
}

func (m *EventMessage) PrepareMessage() []byte {
	var data bytes.Buffer
	if len(m.id) > 0 {
		data.WriteString(fmt.Sprintf("id: %s\n", strings.Replace(m.id, "\n", "", -1)))
	}
	if len(m.event) > 0 {
		data.WriteString(fmt.Sprintf("event: %s\n", strings.Replace(m.event, "\n", "", -1)))
	}
	if len(m.data) > 0 {
		lines := strings.Split(m.data, "\n")
		for _, line := range lines {
			data.WriteString(fmt.Sprintf("data: %s\n", line))
		}
	}
	data.WriteString("\n")
	return data.Bytes()
}

func controlProcess(es *eventSource) {
	for {
		select {
		case em := <-es.sink:
			msg := em.PrepareMessage()
			func() {
				if e, exist := es.consumers.Load(em.GetUserId()); exist {
					c := e.(*consumer)
					// Only send this msg if the consumer isn't staled
					if !c.staled {
						select {
						case c.in <- msg:
						default:
						}
					}
				}
			}()
		case <-es.close:
			close(es.sink)
			close(es.add)
			close(es.staled)
			close(es.close)
			es.consumers.Range(func(key, value any) bool {
				c := value.(*consumer)
				close(c.in)
				return true
			})
			return
		case c := <-es.add:
			func() {
				es.consumers.Store(c.GetUerId(), c)
			}()
		case c := <-es.staled:
			es.consumers.Delete(c.GetUerId())
			close(c.in)
		}
	}
}

// New creates new EventSource instance.
func New(settings *Settings, customHeadersFunc func(*http.Request) [][]byte) EventSource {
	if settings == nil {
		settings = DefaultSettings()
	}

	es := new(eventSource)
	es.customHeadersFunc = customHeadersFunc
	es.sink = make(chan Message, 1)
	es.close = make(chan struct{})
	es.staled = make(chan *consumer, 1)
	es.add = make(chan *consumer)
	es.timeout = settings.Timeout
	es.idleTimeout = settings.IdleTimeout
	es.closeOnTimeout = settings.CloseOnTimeout
	es.gzip = settings.Gzip
	es.ctx = context.Background()
	go controlProcess(es)
	return es
}

// ServeHTTP implements http.Handler interface.
func (es *eventSource) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	cons, err := newConsumer(resp, req, es)
	if err != nil {
		log.Info("Can't create connection to a consumer: ", err)
		return
	}
	es.add <- cons
}

func (es *eventSource) sendMessage(m Message) {
	es.sink <- m
}

func (es *eventSource) SendEventMessage(userId int64, data, event, id string) {
	em := &EventMessage{id, userId, event, data}
	es.sendMessage(em)
}

func (m *RetryMessage) PrepareMessage() []byte {
	return []byte(fmt.Sprintf("retry: %d\n\n", m.retry/time.Millisecond))
}

func (m *RetryMessage) GetUserId() int64 {
	return m.userId
}

func (m *RetryMessage) GetData() string {
	return m.data
}

func (es *eventSource) SendRetryMessage(userId int64, t time.Duration) {
	es.sendMessage(&RetryMessage{t, userId, ""})
}

func (es *eventSource) Close() {
	es.close <- struct{}{}
}
