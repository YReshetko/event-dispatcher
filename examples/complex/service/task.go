package service

import "time"

type Task struct {
	ID       int
	Name     string
	Deadline time.Time
	User     User
}

type User struct {
	ID          string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	DeviceID    string
}
