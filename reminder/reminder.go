package reminder

import (
	"fmt"
	"time"
)

type Reminder struct {
	Message string
	At      time.Time
	Sent    bool
}

func NewReminder(message string, at time.Time) *Reminder {
	return &Reminder{
		Message: message,
		At:      at,
		Sent:    false,
	}
}

func (r *Reminder) Send() {
	if r.Sent {
		return
	}
	fmt.Println("Reminder", r.Message)
	r.Sent = true
}

func (r *Reminder) Stop() {
	//
}
