package cmd

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/google/shlex"
	"github.com/somenave/eventsCalendar/calendar"
	"github.com/somenave/eventsCalendar/logger"
	"strings"
)

type Cmd struct {
	calendar *calendar.Calendar
	logger   *Logger
}

func NewCmd(c *calendar.Calendar, l *Logger) *Cmd {
	return &Cmd{calendar: c, logger: l}
}

func (c *Cmd) Run() {
	err := c.logger.Load()
	if err != nil {
		fmt.Printf("Cannot load logger: %v", err.Error())
	}

	err = c.calendar.Load()
	if err != nil {
		fmt.Println(err)
		c.logger.Add(LogError, err.Error())
		logger.Error(err.Error())
	}

	p := prompt.New(
		c.executor,
		c.completer,
		prompt.OptionPrefix("> "),
	)

	go func() {
		for msg := range c.calendar.Notification {
			fmt.Printf("%s\n", msg)
			c.logger.Add(LogNotification, msg)
			logger.Info(fmt.Sprintf("Notification: %s", msg))
		}
	}()

	p.Run()
}

func (c *Cmd) executor(input string) {
	input = strings.TrimSpace(input)
	if input == "" {
		return
	}

	c.logger.Add(LogInput, input)
	logger.Info(fmt.Sprintf("Input: %s", input))

	parts, err := shlex.Split(input)
	if err != nil {
		fmt.Println(invalidInputErrMsg)
		c.logger.Add(LogError, invalidInputErrMsg)
		logger.Info(invalidInputErrMsg)
		return
	}

	cmd := strings.ToLower(parts[0])
	args := parts[1:]
	switch cmd {
	case "add":
		c.addEvent(args)
	case "list":
		c.showEvents()
	case "remove":
		c.removeEvent()
	case "update":
		c.updateEvent(args)
	case "reminder:set":
		c.setReminder(args)
	case "reminder:remove":
		c.removeReminder(args)
	case "reminder:cancel":
		c.cancelReminder(args)
	case "help":
		fmt.Println("supported commands:")
		fmt.Println(" 'add' >>", addFormatMsg)
		fmt.Println(" 'list'")
		fmt.Println(" 'remove' >>", chooseMsg)
		fmt.Println(" 'update' >>", updateFormatMsg)
		fmt.Println(" 'reminder:set' >>", setReminderFormatMsg)
		fmt.Println(" 'reminder:remove' >>", removeReminderFormatMsg)
		fmt.Println(" 'reminder:cancel' >>", cancelReminderFormatMsg)
		fmt.Println(" 'help'")
	case "logs":
		c.showLogs()
	case "exit":
		c.exit()
	default:
		fmt.Println(invalidCommandErrMsg)
		c.logger.Add(LogOutput, invalidCommandErrMsg)
		logger.Info(invalidCommandErrMsg)
	}
}

func (c *Cmd) completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: "help", Description: "show help"},
		{Text: "add", Description: "add new event"},
		{Text: "list", Description: "list all events"},
		{Text: "remove", Description: "remove event"},
		{Text: "update", Description: "update event"},
		{Text: "reminder:set", Description: "set reminder for event"},
		{Text: "reminder:remove", Description: "remove reminder for event"},
		{Text: "reminder:cancel", Description: "cancel reminder for event"},
		{Text: "logs", Description: "show logs"},
		{Text: "exit", Description: "exit program"},
	}
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}
