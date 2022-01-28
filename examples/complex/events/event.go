package events

type EventType string

const (
	OnTaskCreated EventType = "TASK_CREATED"
	OnTaskUpdated EventType = "TASK_UPDATED"
	OnTaskDeleted EventType = "TASK_DELETED"
)

type Event struct {
	EventID  string
	TaskID   int
	TaskName string
	Action   string
	User     struct {
		Name     string
		Phone    string
		Email    string
		DeviceID string
	}
}
