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
	event, err := buildEvent(getNextID(), title, startDate, priority)
	if err != nil {
		return &Event{}, fmt.Errorf("can't create event: %w", errors.New(err.Error()))
	}
	return event, nil
}

func (e *Event) Update(title string, dateStr string, priority Priority) error {
	event, err := buildEvent(e.ID, title, dateStr, priority)
	if err != nil {
		return fmt.Errorf("can't update event: %w", errors.New(err.Error()))
	}

	e.Title = event.Title
	e.StartAt = event.StartAt
	e.Priority = event.Priority
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

func (e Event) String() string {
	return e.ID + ": " + e.Title + " — " + e.StartAt.Format("02 Jan 2006 15:04, Mon") + " — " + string(e.Priority)
}

func getNextID() string {
	return uuid.New().String()
}

func (e *Event) AddReminder(message string, at string, notify func(string)) error {
	r, err := reminder.NewReminder(message, at, notify)
	if err != nil {
		return fmt.Errorf("can't create reminder: %w", err)
	}
	e.Reminder = r
	r.Start()
	return nil
}

func (e *Event) StopReminder() error {
	if e.Reminder == nil {
		return errors.New("reminder doesn't exist")
	}
	err := e.Reminder.Stop()
	if err != nil {
		return fmt.Errorf("can't stop reminder: %w", errors.New(err.Error()))
	}
	return nil
}

func (e *Event) RemoveReminder() error {
	if e.Reminder == nil {
		return errors.New("reminder doesn't exist")
	}
	e.Reminder = nil
	return nil
}
