package notifications

import (
	"fmt"
	event_dispatcher "github.com/YReshetko/event-dispatcher"
	"github.com/YReshetko/event-dispatcher/examples/complex/events"
)

type Push struct{}

func NewPush(d *event_dispatcher.EventDispatcher[events.EventType, events.Event]) *Push {
	e := &Push{}
	d.AddEventListener(events.OnTaskCreated, e.handleCreated)
	return e
}

func (e *Push) handleCreated(event events.Event) {
	fmt.Println("Push to", event.User.Name)
	fmt.Println("Device ID", event.User.DeviceID)
	fmt.Println("Message: New task ID:", event.TaskID, ",", event.TaskName, ".")
}
