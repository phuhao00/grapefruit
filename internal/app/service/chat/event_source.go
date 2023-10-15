package chat

import (
	"grapefruit/internal/adapter/eventsource"
	"grapefruit/kit/log"
	"net/http"
)

type EventSource struct {
	real eventsource.EventSource
}

func NewEventSource() *EventSource {
	return &EventSource{real: nil}
}

func (e *EventSource) RunAsClient() {

}

func (e *EventSource) RunAsServer() {
	e.real = eventsource.New(nil, nil)
	defer e.real.Close()
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.Handle("/events", e.real)
	log.Info("Open URL http://localhost:8081/ in your browser.")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (e *EventSource) Start() {
	go e.RunAsServer()
	go e.RunAsClient()
}
func (e *EventSource) Stop() {

}
