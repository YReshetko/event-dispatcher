package notifications

import (
	"fmt"
	event_dispatcher "github.com/YReshetko/event-dispatcher"
	"github.com/YReshetko/event-dispatcher/examples/complex/events"
)

type Sms struct{}

func NewSms(d *event_dispatcher.EventDispatcher[events.EventType, events.Event]) *Sms {
	e := &Sms{}
	d.AddEventListener(events.OnTaskDeleted, e.handleDeleted)
	return e
}

func (e *Sms) handleDeleted(event events.Event) {
	fmt.Println("SMS to", event.User.Name)
	fmt.Println("Phone number", event.User.Phone)
	fmt.Println("Message: task ", event.TaskID, "was deleted. Please stop your work!")
}
