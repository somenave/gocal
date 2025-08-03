package calendar

import (
	"errors"
	"fmt"
	"github.com/somenave/eventsCalendar/events"
)

var calendarEvents = make(map[string]events.Event)

func AddEvent(title string, date string) (events.Event, error) {
	event, err := events.NewEvent(title, date)
	if err != nil {
		return events.Event{}, err
	}

	for _, existingEvent := range calendarEvents {
		if existingEvent.StartAt.Equal(event.StartAt) {
			return events.Event{}, errors.New("there is already an event at this time")
		}
	}

	calendarEvents[event.ID] = event
	return event, nil
}

func EditEvent(id string, title string, startAt string) error {
	event, exist := calendarEvents[id]
	if !exist {
		return errors.New("There is no event with id " + id)
	}
	if !events.IsValidTitle(title) {
		return errors.New("title is not valid")
	}

	date, err := events.ParseDate(startAt)
	if err != nil {
		return errors.New("date is not valid")
	}

	event.Title = title
	event.StartAt = date

	calendarEvents[id] = event

	return nil
}

func DeleteEvent(id string) {
	event, exist := calendarEvents[id]
	if exist {
		delete(calendarEvents, id)
		fmt.Println("Deleted event:", getEventStr(event))
	} else {
		fmt.Println("There is no event with id:", id)
	}
}

func ShowEvents() {
	fmt.Println("---")
	for _, event := range calendarEvents {
		fmt.Println(getEventStr(event))
	}
	fmt.Println("---")
}

func getEventStr(event events.Event) string {
	return event.Title + " — " + event.StartAt.Format("02 Jan 2006 15:04, Mon") + " — " + event.ID
}
