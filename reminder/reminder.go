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
	notify  func(string)
}

func NewReminder(message string, at string, notify func(string)) (*Reminder, error) {
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
		notify:  notify,
	}, nil
}

func (r *Reminder) sendMessage() error {
	if r.Sent {
		return errors.New("reminder already sent")
	}
	if r.notify != nil {
		r.notify(r.Message)
		r.Sent = true
		return nil
	}
	return errors.New("reminder not sent")
}

func (r *Reminder) Send() {
	err := r.sendMessage()
	if err != nil {
		fmt.Println(err)
	}
}

func (r *Reminder) Start() {
	duration := time.Until(r.At)
	r.timer = time.AfterFunc(duration, r.Send)
}

func (r *Reminder) Stop() error {
	if r.timer == nil {
		return errors.New("could not stop the reminder, timer doesn't exist")
	}
	stopped := r.timer.Stop()
	if stopped {
		fmt.Println("Reminder stopped")
	} else {
		return errors.New("reminder has already expired or been stopped")
	}
	return nil
}
