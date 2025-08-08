package cmd

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/google/shlex"
	"github.com/somenave/eventsCalendar/calendar"
	"os"
	"strings"
)

type Cmd struct {
	calendar *calendar.Calendar
}

func NewCmd(c *calendar.Calendar) *Cmd {
	return &Cmd{calendar: c}
}

func (c *Cmd) executor(input string) {
	parts, err := shlex.Split(input)
	if err != nil {
		fmt.Println("invalid input")
		return
	}
	if len(parts) < 1 {
		fmt.Println("input is empty, type command")
		return
	}

	cmd := strings.ToLower(parts[0])
	switch cmd {
	case "add":
		if len(parts) != 4 {
			fmt.Println("Format: add 'event name' 'date' 'priority'")
			return
		}
		title := parts[1]
		date := parts[2]
		priority := parts[3]

		e, err := c.calendar.AddEvent(title, date, priority)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Event", e.Title, "has been added")
		}
	case "list":
		c.calendar.ShowEvents()
	case "remove":
		if len(parts) != 2 {
			fmt.Println("Format: remove 'event ID'")
			return
		}
		eventId := parts[1]
		err := c.calendar.DeleteEvent(eventId)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Event has been removed")
		}
	case "update":
		if len(parts) != 5 {
			fmt.Println("Format: update 'event ID' 'name' 'date' 'priority'")
			return
		}
		id, name, date, priority := parts[1], parts[2], parts[3], parts[4]
		err := c.calendar.EditEvent(id, name, date, priority)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Event has been updated")
	case "help":
		fmt.Println("supported commands:")
		fmt.Println(" 'add' >> format: add 'event name' 'date' 'priority'")
		fmt.Println(" 'list'")
		fmt.Println(" 'remove' >> format: remove 'event ID'")
		fmt.Println(" 'update' >> format: update 'event ID' 'name' 'date' 'priority'")
		fmt.Println(" 'help'")
	case "exit":
		err := c.calendar.Save()
		if err != nil {
			fmt.Println("Error saving calendar:", err)
		} else {
			fmt.Println("Calendar saved")
		}
		os.Exit(0)
	default:
		fmt.Println("invalid command")
		fmt.Println("type 'help' for list of supported commands.")
	}
}

func (c *Cmd) completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: "add", Description: "add new event"},
		{Text: "list", Description: "list all events"},
		{Text: "remove", Description: "remove event"},
		{Text: "update", Description: "update event"},
		{Text: "help", Description: "show help"},
		{Text: "exit", Description: "exit program"},
	}
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}

func (c *Cmd) Run() {
	err := c.calendar.Load()
	if err != nil {
		fmt.Println(err)
	}

	p := prompt.New(
		c.executor,
		c.completer,
		prompt.OptionPrefix("> "),
	)
	p.Run()
}
