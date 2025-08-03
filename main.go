package main

import (
	"fmt"
	"github.com/somenave/eventsCalendar/calendar"
)

func main() {
	event1, err1 := calendar.AddEvent("Event 1", "2025/08/10 20:00")
	if err1 != nil {
		fmt.Println("Error when adding new event:", err1)
	}
	event2, err2 := calendar.AddEvent("Event 2", "2025/08/12 15:00")
	if err2 != nil {
		fmt.Println("Error when adding new event:", err2)
	}
	calendar.ShowEvents()

	calendar.DeleteEvent(event1.ID)

	err3 := calendar.EditEvent(event2.ID, "Event, 2 Updated.", "2025/08/12 16:50")
	if err3 != nil {
		fmt.Println("Error when editing event:", err3)
	}

	err4 := calendar.EditEvent(event1.ID, "Should be error", "2025/08/12 20:00")
	if err4 != nil {
		fmt.Println("Error when editing event:", err4)
	}

	_, err5 := calendar.AddEvent("E", "2025/08/12 15:00")
	if err5 != nil {
		fmt.Println("Error when adding new event:", err5)
	}

	calendar.DeleteEvent(event1.ID)

	calendar.ShowEvents()
}
