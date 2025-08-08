package main

import (
	"fmt"
	"github.com/somenave/eventsCalendar/calendar"
	"github.com/somenave/eventsCalendar/storage"
)

func main() {
	s := storage.NewZipStorage("calendar.zip")
	c := calendar.NewCalendar(s)
	defer func() {
		err := c.Save()
		if err != nil {
			fmt.Println("Error saving calendar:", err)
		}
	}()

	err := c.Load()
	if err != nil {
		fmt.Println(err)
	}

	c.ShowEvents()

	event1, err1 := c.AddEvent("Meeting", "2025/06/12 12:00", "low")
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println(event1.Title, "added")
	}

	event2, err2 := c.AddEvent("One more meeting", "2025/06/12 16:00", "medium")
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println(event2.Title, "added")
	}

	err3 := c.EditEvent(event2.ID, "Call", "2025/06/12 16:50", "high")
	if err3 != nil {
		fmt.Println("Error:", err3)
	} else {
		fmt.Println("Event updated")
	}

	err4 := c.DeleteEvent(event1.ID)
	if err4 != nil {
		fmt.Println("Error:", err4)
	} else {
		fmt.Println("Event deleted")
	}

	c.ShowEvents()
}
