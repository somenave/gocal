package reminder

import (
	"errors"
	"github.com/somenave/eventsCalendar/helpers"
	"strings"
	"time"
)

var ErrEmptyMessage = errors.New("message is empty")
var ErrTimerNotExist = errors.New("timer doesn't exist")

type Reminder struct {
	Message string    `json:"message"`
	At      time.Time `json:"at"`
	timer   *time.Timer
	Sent    bool `json:"sent"`
	notify  func(string)
}

func NewReminder(message string, at string, notify func(string)) (*Reminder, error) {
	if len(strings.TrimSpace(message)) == 0 {
		return nil, ErrEmptyMessage
	}
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

func (r *Reminder) Send() {
	if r.Sent || r.notify == nil {
		return
	}
	r.notify(r.Message)
	r.Sent = true
}

func (r *Reminder) Start() {
	duration := time.Until(r.At)
	r.timer = time.AfterFunc(duration, r.Send)
}

func (r *Reminder) Stop() error {
	if r.timer == nil {
		return ErrTimerNotExist
	}
	stopped := r.timer.Stop()
	if !stopped {
		return errors.New("reminder has already expired or been stopped")
	}
	return nil
}
