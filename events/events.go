package events

import (
	"errors"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
	"time"
)

type Event struct {
	ID      string
	Title   string
	StartAt time.Time
}

func NewEvent(title string, dateStr string) (Event, error) {
	if !IsValidTitle(title) {
		return Event{}, errors.New("title is not valid")
	}

	date, err := ParseDate(dateStr)
	if err != nil {
		return Event{}, errors.New("date is not valid")
	}

	return Event{
		ID:      getNextID(),
		Title:   title,
		StartAt: date,
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
