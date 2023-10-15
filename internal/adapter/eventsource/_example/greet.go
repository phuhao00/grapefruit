package main

import (
	"grapefruit/internal/adapter/eventsource"
	"log"
	"net/http"
	"time"
)

func main() {
	es := eventsource.New(nil, nil)
	defer es.Close()
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.Handle("/events", es)
	go func() {
		for {
			es.SendEventMessage(0, "hello", "", "")
			time.Sleep(2 * time.Second)
		}
	}()
	log.Print("Open URL http://localhost:8080/ in your browser.")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
