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
	args := parts[1:]
	switch cmd {
	case "add":
		c.addEvent(args)
	case "list":
		c.calendar.ShowEvents()
	case "remove":
		c.removeEvent(args)
	case "update":
		c.updateEvent(args)
	case "reminder:set":
		c.setReminder(args)
	case "reminder:remove":
		c.removeReminder(args)
	case "help":
		fmt.Println("supported commands:")
		fmt.Println(" 'add' >> format: add 'event name' 'date' 'priority'")
		fmt.Println(" 'list'")
		fmt.Println(" 'remove' >> format: remove 'event ID'")
		fmt.Println(" 'update' >> format: update 'event ID' 'name' 'date' 'priority'")
		fmt.Println(" 'reminder:set' >> format: reminder:set 'event ID' 'message' 'date'")
		fmt.Println(" 'reminder:remove' >> format: reminder:remove 'event ID'")
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
		{Text: "reminder:set", Description: "set reminder for event"},
		{Text: "reminder:remove", Description: "remove reminder for event"},
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
