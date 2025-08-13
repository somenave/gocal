package cmd

import "fmt"

func (c *Cmd) addEvent(args []string) {
	if len(args) != 3 {
		fmt.Println("Format: add 'event name' 'date' 'priority'")
		return
	}
	title, date, priority := args[0], args[1], args[2]

	e, err := c.calendar.AddEvent(title, date, priority)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Event", e.Title, "has been added")
	}
}

func (c *Cmd) removeEvent(args []string) {
	if len(args) != 1 {
		fmt.Println("Format: remove 'event ID'")
		return
	}
	eventId := args[0]
	err := c.calendar.DeleteEvent(eventId)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Event has been removed")
	}
}

func (c *Cmd) updateEvent(args []string) {
	if len(args) != 4 {
		fmt.Println("Format: update 'event ID' 'name' 'date' 'priority'")
		return
	}
	id, name, date, priority := args[0], args[1], args[2], args[3]
	err := c.calendar.EditEvent(id, name, date, priority)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Event has been updated")
}

func (c *Cmd) setReminder(args []string) {
	if len(args) != 3 {
		fmt.Println("Format: reminder:set 'event ID' 'message' 'at'")
		return
	}
	id, message, at := args[0], args[1], args[2]
	err := c.calendar.SetEventReminder(id, message, at)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Reminder has been set")
}

func (c *Cmd) removeReminder(args []string) {
	if len(args) != 1 {
		fmt.Println("Format: reminder:remove 'event ID'")
		return
	}
	eventId := args[0]
	err := c.calendar.RemoveEventReminder(eventId)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Reminder has been removed")
}

func (c *Cmd) cancelReminder(args []string) {
	if len(args) != 1 {
		fmt.Println("Format: reminder:cancel 'event ID'")
		return
	}
	eventId := args[0]
	err := c.calendar.CancelEventReminder(eventId)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Reminder has been cancelled")
}
