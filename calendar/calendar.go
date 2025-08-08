package calendar

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/somenave/eventsCalendar/events"
	"github.com/somenave/eventsCalendar/reminder"
	"github.com/somenave/eventsCalendar/storage"
	"time"
)

type Calendar struct {
	calendarEvents map[string]*events.Event
	storage        storage.Store
}

func NewCalendar(s storage.Store) *Calendar {
	data := make(map[string]*events.Event)
	return &Calendar{
		calendarEvents: data,
		storage:        s,
	}
}

func (calendar *Calendar) AddEvent(title string, date string, priority string) (*events.Event, error) {
	event, err := events.NewEvent(title, date, priority)
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

func (calendar *Calendar) EditEvent(id string, title string, startAt string, priority string) error {
	event, exist := calendar.calendarEvents[id]
	if !exist {
		return errors.New("there is no event with id " + id)
	}

	err := event.Update(title, startAt, priority)

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

func (calendar *Calendar) setEventReminder(id string, message string, at time.Time) error {
	event, exist := calendar.calendarEvents[id]
	if !exist {
		return errors.New("there is no event with id " + id)
	}
	event.Reminder = reminder.NewReminder(message, at)
	return nil
}

func (calendar *Calendar) Save() error {
	data, err := json.Marshal(calendar.calendarEvents)
	if err != nil {
		return err
	}
	err = calendar.storage.Save(data)
	return err
}

func (calendar *Calendar) Load() error {
	data, err := calendar.storage.Load()
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &calendar.calendarEvents)
	return err
}
