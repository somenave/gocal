package events

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/somenave/eventsCalendar/helpers"
	"github.com/somenave/eventsCalendar/reminder"
	"time"
)

type Event struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	StartAt  time.Time `json:"start_at"`
	Priority Priority
	Reminder *reminder.Reminder
}

func NewEvent(title string, startDate string, priority Priority) (*Event, error) {
	return buildEvent(getNextID(), title, startDate, priority)
}

func (e *Event) Update(title string, dateStr string, priority Priority) error {
	event, err := buildEvent(e.ID, title, dateStr, priority)
	if err != nil {
		return err
	}

	e.Title = event.Title
	e.StartAt = event.StartAt
	return nil
}

func buildEvent(id string, title string, dateStr string, priority Priority) (*Event, error) {
	if !IsValidTitle(title) {
		return &Event{}, errors.New("title is not valid")
	}

	startDate, err := helpers.ParseDate(dateStr)
	if err != nil {
		return &Event{}, errors.New("date is not valid")
	}

	priorityErr := Priority.Validate(priority)
	if priorityErr != nil {
		return &Event{}, priorityErr
	}

	return &Event{
		ID:       id,
		Title:    title,
		StartAt:  startDate,
		Priority: priority,
		Reminder: nil,
	}, nil
}

func (e Event) Print() {
	fmt.Println(e.ID + ": " + e.Title + " — " + e.StartAt.Format("02 Jan 2006 15:04, Mon") + " — " + string(e.Priority))
}

func getNextID() string {
	return uuid.New().String()
}

func (e *Event) AddReminder(message string, at string) error {
	r, err := reminder.NewReminder(message, at)
	if err != nil {
		return err
	}
	e.Reminder = r
	r.Start()
	return nil
}

func (e *Event) RemoveReminder() error {
	if e.Reminder != nil {
		e.Reminder.Stop()
		e.Reminder = nil
		return nil
	}
	return errors.New("reminder is not set")
}
