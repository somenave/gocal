package events

import (
	"errors"
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
	"time"
)

type Event struct {
	ID       string
	Title    string
	StartAt  time.Time
	Priority Priority
}

func NewEvent(title string, startDate string, priority string) (*Event, error) {
	return buildEvent(getNextID(), title, startDate, priority)
}

func (e Event) Print() {
	fmt.Println(e.Title + " — " + e.StartAt.Format("02 Jan 2006 15:04, Mon") + " — " + string(e.Priority) + " — " + e.ID)
}

func (e *Event) Update(title string, dateStr string, priorityStr string) error {
	event, err := buildEvent(e.ID, title, dateStr, priorityStr)
	if err != nil {
		return err
	}

	e.Title = event.Title
	e.StartAt = event.StartAt
	return nil
}

func buildEvent(id string, title string, dateStr string, priorityStr string) (*Event, error) {
	if !IsValidTitle(title) {
		return &Event{}, errors.New("title is not valid")
	}

	startDate, err := ParseDate(dateStr)
	if err != nil {
		return &Event{}, errors.New("date is not valid")
	}

	priority, priorityErr := NewPriority(priorityStr)
	if priorityErr != nil {
		return &Event{}, priorityErr
	}

	return &Event{
		ID:       id,
		Title:    title,
		StartAt:  startDate,
		Priority: priority,
	}, nil
}

func ParseDate(date string) (time.Time, error) {
	dateParsed, err := dateparse.ParseAny(date)
	if err != nil {
		return time.Time{}, err
	}
	return dateParsed, nil
}

func getNextID() string {
	return uuid.New().String()
}
