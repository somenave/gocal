package calendar

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/somenave/eventsCalendar/events"
	"github.com/somenave/eventsCalendar/storage"
	"strings"
)

var ErrEventAtThisTimeExist = errors.New("there is already an event at this time")
var ErrEventDoesNotExist = errors.New("there is no event with given id")

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

func (c *Calendar) GetEvents() map[string]*events.Event {
	return c.calendarEvents
}

func (c *Calendar) FindEventByTitle(title string) *events.Event {
	for _, e := range c.calendarEvents {
		if strings.Contains(strings.ToLower(e.Title), strings.ToLower(title)) {
			return e
		}
	}
	return nil
}

func (c *Calendar) AddEvent(title string, date string, priority string) (*events.Event, error) {
	event, err := events.NewEvent(title, date, events.Priority(priority))
	if err != nil {
		return &events.Event{}, err
	}

	for _, existingEvent := range c.calendarEvents {
		if existingEvent.StartAt.Equal(event.StartAt) {
			return &events.Event{}, ErrEventAtThisTimeExist
		}
	}

	c.calendarEvents[event.ID] = event
	return event, nil
}

func (c *Calendar) EditEvent(id string, title string, startAt string, priority string) error {
	event, existErr := c.checkEventExist(id)
	if existErr != nil {
		return fmt.Errorf("can't edit event: %w", existErr)
	}

	return event.Update(title, startAt, events.Priority(priority))
}

func (c *Calendar) DeleteEvent(id string) error {
	_, existErr := c.checkEventExist(id)
	if existErr != nil {
		return fmt.Errorf("can't delete event: %w", existErr)
	}

	delete(c.calendarEvents, id)
	return nil
}

func (c *Calendar) SetEventReminder(id string, message string, at string) error {
	event, existErr := c.checkEventExist(id)
	if existErr != nil {
		return fmt.Errorf("can't set event reminder: %w", existErr)
	}
	return event.AddReminder(message, at, c.Notify)
}

func (c *Calendar) RemoveEventReminder(id string) error {
	event, existErr := c.checkEventExist(id)
	if existErr != nil {
		return fmt.Errorf("can't remove event reminder: %w", existErr)
	}
	return event.RemoveReminder()
}

func (c *Calendar) CancelEventReminder(id string) error {
	event, existErr := c.checkEventExist(id)
	if existErr != nil {
		return fmt.Errorf("can't cancel event reminder: %w", existErr)
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
		return fmt.Errorf("can't save calendar: %w", err)
	}
	return c.storage.Save(data)
}

func (c *Calendar) Load() error {
	data, err := c.storage.Load()
	if err != nil {
		return fmt.Errorf("can't load calendar: %w", err)
	}
	return json.Unmarshal(data, &c.calendarEvents)
}

func (c *Calendar) checkEventExist(id string) (*events.Event, error) {
	event, exist := c.calendarEvents[id]
	if exist {
		return event, nil
	}
	return nil, ErrEventDoesNotExist
}
