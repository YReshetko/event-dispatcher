package service

import (
	"math/rand"
	"strconv"
	"time"

	event_dispatcher "github.com/YReshetko/event-dispatcher"
	"github.com/YReshetko/event-dispatcher/examples/complex/events"
)

type TaskService struct {
	Disp *event_dispatcher.EventDispatcher[events.EventType, events.Event]
}

func (t *TaskService) Create(task Task) {
	// Do what you need with the task
	t.Disp.DispatchEvent(events.OnTaskCreated, event(task))
}

func (t *TaskService) Update(task Task) {
	// Do what you need with the task
	t.Disp.DispatchEvent(events.OnTaskUpdated, event(task))
}

func (t *TaskService) Delete(task Task) {
	// Do what you need with the task
	t.Disp.DispatchEvent(events.OnTaskDeleted, event(task))
}

func event(task Task) events.Event {
	return events.Event{
		EventID:  strconv.Itoa(int(rand.Int63n(1000000))),
		TaskID:   task.ID,
		TaskName: task.Name,
		Action:   "Complete it by " + task.Deadline.Format(time.RFC3339),
		User: struct {
			Name     string
			Phone    string
			Email    string
			DeviceID string
		}{
			Name:     task.User.FirstName + " " + task.User.LastName,
			Phone:    task.User.PhoneNumber,
			Email:    task.User.Email,
			DeviceID: task.User.DeviceID,
		},
	}
}
