package calendar

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/somenave/eventsCalendar/events"
	"github.com/somenave/eventsCalendar/storage"
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

func (calendar *Calendar) ShowEvents() {
	if len(calendar.calendarEvents) == 0 {
		fmt.Println("No events found")
		return
	}
	fmt.Println("---")
	for _, event := range calendar.calendarEvents {
		event.Print()
	}
	fmt.Println("---")
}

func (calendar *Calendar) AddEvent(title string, date string, priority string) (*events.Event, error) {
	event, err := events.NewEvent(title, date, events.Priority(priority))
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
	event, existErr := calendar.checkEventExist(id)
	if existErr != nil {
		return existErr
	}

	return event.Update(title, startAt, events.Priority(priority))
}

func (calendar *Calendar) DeleteEvent(id string) error {
	_, existErr := calendar.checkEventExist(id)
	if existErr != nil {
		return existErr
	}

	delete(calendar.calendarEvents, id)
	return nil
}

func (calendar *Calendar) SetEventReminder(id string, message string, at string) error {
	event, existErr := calendar.checkEventExist(id)
	if existErr != nil {
		return existErr
	}
	return event.AddReminder(message, at)
}

func (calendar *Calendar) RemoveEventReminder(id string) error {
	event, existErr := calendar.checkEventExist(id)
	if existErr != nil {
		return existErr
	}
	event.RemoveReminder()
	return nil
}

func (calendar *Calendar) Save() error {
	data, err := json.Marshal(calendar.calendarEvents)
	if err != nil {
		return err
	}
	return calendar.storage.Save(data)
}

func (calendar *Calendar) Load() error {
	data, err := calendar.storage.Load()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &calendar.calendarEvents)
}

func (calendar *Calendar) checkEventExist(id string) (*events.Event, error) {
	event, exist := calendar.calendarEvents[id]
	if exist {
		return event, nil
	}
	return nil, errors.New("there is no event with id " + id)
}
