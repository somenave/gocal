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
	Notification   chan string
}

func NewCalendar(s storage.Store) *Calendar {
	data := make(map[string]*events.Event)
	return &Calendar{
		calendarEvents: data,
		storage:        s,
		Notification:   make(chan string),
	}
}

func (c *Calendar) ShowEvents() {
	if len(c.calendarEvents) == 0 {
		fmt.Println("No events found")
		return
	}
	fmt.Println("---")
	for _, event := range c.calendarEvents {
		event.Print()
	}
	fmt.Println("---")
}

func (c *Calendar) AddEvent(title string, date string, priority string) (*events.Event, error) {
	event, err := events.NewEvent(title, date, events.Priority(priority))
	if err != nil {
		return &events.Event{}, err
	}

	for _, existingEvent := range c.calendarEvents {
		if existingEvent.StartAt.Equal(event.StartAt) {
			return &events.Event{}, errors.New("there is already an event at this time")
		}
	}

	c.calendarEvents[event.ID] = event
	return event, nil
}

func (c *Calendar) EditEvent(id string, title string, startAt string, priority string) error {
	event, existErr := c.checkEventExist(id)
	if existErr != nil {
		return existErr
	}

	return event.Update(title, startAt, events.Priority(priority))
}

func (c *Calendar) DeleteEvent(id string) error {
	_, existErr := c.checkEventExist(id)
	if existErr != nil {
		return existErr
	}

	delete(c.calendarEvents, id)
	return nil
}

func (c *Calendar) SetEventReminder(id string, message string, at string) error {
	event, existErr := c.checkEventExist(id)
	if existErr != nil {
		return existErr
	}
	return event.AddReminder(message, at, c.Notify)
}

func (c *Calendar) RemoveEventReminder(id string) error {
	event, existErr := c.checkEventExist(id)
	if existErr != nil {
		return existErr
	}
	return event.RemoveReminder()
}

func (c *Calendar) CancelEventReminder(id string) error {
	event, existErr := c.checkEventExist(id)
	if existErr != nil {
		return existErr
	}
	err := event.StopReminder()
	return err
}

func (c *Calendar) Notify(msg string) {
	c.Notification <- msg
}

func (c *Calendar) Save() error {
	data, err := json.Marshal(c.calendarEvents)
	if err != nil {
		return err
	}
	return c.storage.Save(data)
}

func (c *Calendar) Load() error {
	data, err := c.storage.Load()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &c.calendarEvents)
}

func (c *Calendar) checkEventExist(id string) (*events.Event, error) {
	event, exist := c.calendarEvents[id]
	if exist {
		return event, nil
	}
	return nil, errors.New("there is no event with id " + id)
}
