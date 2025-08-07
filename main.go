package main

import (
	"fmt"
	"github.com/somenave/eventsCalendar/calendar"
)

func main() {
	c := calendar.NewCalendar()

	event1, err1 := c.AddEvent("Meeting", "2025/06/12 12:00")
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println(event1.Title, "added")
	}

	event2, err2 := c.AddEvent("One more meeting", "2025/06/12 16:00")
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println(event2.Title, "added")
	}

	err3 := c.EditEvent(event2.ID, "Call", "2025/06/12 16:50")
	if err3 != nil {
		fmt.Println("Error:", err3)
	} else {
		fmt.Println("Event updated")
	}

	err := c.DeleteEvent(event1.ID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Event deleted")
	}

	c.ShowEvents()
}
