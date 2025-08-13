package reminder

import (
	"errors"
	"fmt"
	"github.com/somenave/eventsCalendar/helpers"
	"time"
)

type Reminder struct {
	Message string    `json:"message"`
	At      time.Time `json:"at"`
	timer   *time.Timer
	Sent    bool `json:"sent"`
}

func NewReminder(message string, at string) (*Reminder, error) {
	date, err := helpers.ParseDate(at)
	if err != nil {
		return nil, err
	}
	if date.Before(time.Now()) {
		return nil, errors.New("reminder date must be in the future")
	}
	return &Reminder{
		Message: message,
		At:      date,
		timer:   nil,
		Sent:    false,
	}, nil
}

func (r *Reminder) Send() {
	if r.Sent {
		return
	}
	fmt.Println("Reminder:", r.Message)
	r.Sent = true
}

func (r *Reminder) Start() {
	duration := time.Until(r.At)
	r.timer = time.AfterFunc(duration, r.Send)
}

func (r *Reminder) Stop() {
	stopped := r.timer.Stop()
	if stopped {
		fmt.Println("Reminder stopped")
	} else {
		fmt.Println("Reminder has already expired or been stopped")
	}
}
