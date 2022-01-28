package main

import (
	"fmt"
	"time"

	"github.com/YReshetko/event-dispatcher"
)

type EventType string

const (
	Create EventType = "CREATE"
	Update EventType = "UPDATE"
	Delete EventType = "DELETE"
)

type Event struct {
	target string
}

func main() {
	dispatcher := event_dispatcher.NewEventDispatcher[EventType, Event](10)

	dispatcher.AddEventListener(Create, func(event Event) {
		fmt.Println("Create Listener 1; event target:", event.target)
	})
	dispatcher.AddEventListener(Create, func(event Event) {
		fmt.Println("Create Listener 2; event target:", event.target)
	})

	dispatcher.AddEventListener(Update, func(event Event) {
		fmt.Println("Update Listener 1; event target:", event.target)
	})
	dispatcher.AddEventListener(Update, func(event Event) {
		fmt.Println("Update Listener 2; event target:", event.target)
	})

	dispatcher.AddEventListener(Delete, func(event Event) {
		fmt.Println("Delete Listener 1; event target:", event.target)
	})
	dispatcher.AddEventListener(Delete, func(event Event) {
		fmt.Println("Delete Listener 2; event target:", event.target)
	})

	dispatcher.DispatchEvent(Create, Event{target: "first-target"})
	dispatcher.DispatchEvent(Update, Event{target: "first-target"})
	dispatcher.DispatchEvent(Delete, Event{target: "first-target"})

	time.Sleep(time.Second)
	//dispatcher.Close()
}
