package main

import (
	"github.com/YReshetko/event-dispatcher"
	"github.com/YReshetko/event-dispatcher/examples/complex/events"
	"github.com/YReshetko/event-dispatcher/examples/complex/notifications"
	"github.com/YReshetko/event-dispatcher/examples/complex/service"
	"math/rand"
	"time"
)

func main() {
	d := event_dispatcher.NewEventDispatcher[events.EventType, events.Event](3)

	s := service.TaskService{d}

	notifications.NewEmail(d)
	notifications.NewPush(d)
	notifications.NewSms(d)

	actions := []func(service.Task){s.Create, s.Update, s.Delete}

	for i := 0; i < 100; i++ {
		a := actions[rand.Intn(len(actions))]
		go a(newTask())
	}

	time.Sleep(time.Second * 10)

}

var taskNames = []string{"Create service N", "Update service N", "Integrate with subsystem N"}

func newTask() service.Task {
	return service.Task{
		ID:       rand.Int(),
		Name:     taskNames[rand.Intn(len(taskNames))],
		Deadline: time.Now().Add(time.Hour * time.Duration(rand.Int63n(100000))),
		User: service.User{
			ID:          "user-1",
			FirstName:   "John",
			LastName:    "Doe",
			Email:       "john.doe@example.com",
			PhoneNumber: "+000(00)-00-00-000",
			DeviceID:    "1234567890",
		},
	}
}
