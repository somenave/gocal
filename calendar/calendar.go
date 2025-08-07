package calendar

import (
	"errors"
	"fmt"
	"github.com/somenave/eventsCalendar/events"
)

type Calendar struct {
	calendarEvents map[string]*events.Event
}

func NewCalendar() *Calendar {
	return &Calendar{
		calendarEvents: make(map[string]*events.Event),
	}
}

func (calendar *Calendar) AddEvent(title string, date string) (*events.Event, error) {
	event, err := events.NewEvent(title, date)
	if err != nil {
		return &events.Event{}, err
	}

	for _, existingEvent := range calendar.calendarEvents {
		if existingEvent.StartAt.Equal(event.StartAt) {
			return &events.Event{}, errors.New("there is already an event at this time")
		}
	}

	calendar.calendarEvents[event.ID] = event
	return event, nil
}

func (calendar *Calendar) EditEvent(id string, title string, startAt string) error {
	event, exist := calendar.calendarEvents[id]
	if !exist {
		return errors.New("there is no event with id " + id)
	}

	err := event.Update(title, startAt)

	return err
}

func (calendar *Calendar) DeleteEvent(id string) error {
	_, exist := calendar.calendarEvents[id]
	if exist {
		delete(calendar.calendarEvents, id)
		return nil
	} else {
		return errors.New("there is no event with id " + id)
	}
}

func (calendar *Calendar) ShowEvents() {
	fmt.Println("---")
	for _, event := range calendar.calendarEvents {
		event.Print()
	}
	fmt.Println("---")
}
