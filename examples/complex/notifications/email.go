package notifications

import (
	"fmt"
	event_dispatcher "github.com/YReshetko/event-dispatcher"
	"github.com/YReshetko/event-dispatcher/examples/complex/events"
)

type Email struct{}

func NewEmail(d *event_dispatcher.EventDispatcher[events.EventType, events.Event]) *Email {
	e := &Email{}
	d.AddEventListener(events.OnTaskCreated, e.handleCreated)
	d.AddEventListener(events.OnTaskUpdated, e.handleUpdated)
	d.AddEventListener(events.OnTaskDeleted, e.handleDeleted)
	return e
}

func (e *Email) handleCreated(event events.Event) {
	fmt.Println("Sending Email to", event.User.Name)
	fmt.Println("From go@example.com")
	fmt.Println("To", event.User.Email)
	fmt.Println("Message: Ypu have a new task ID:", event.TaskID, ",", event.TaskName, ".\nTODO:", event.Action)
}

func (e *Email) handleUpdated(event events.Event) {
	fmt.Println("Sending Email to", event.User.Name)
	fmt.Println("From go@example.com")
	fmt.Println("To", event.User.Email)
	fmt.Println("Message: Check your task ID:", event.TaskID, ",", event.TaskName, ", it was updated.\nTODO:", event.Action)
}

func (e *Email) handleDeleted(event events.Event) {
	fmt.Println("Sending Email to", event.User.Name)
	fmt.Println("From go@example.com")
	fmt.Println("To", event.User.Email)
	fmt.Println("Message: Deleted task ID:", event.TaskID, ",", event.TaskName, ".")
}
