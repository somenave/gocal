package cmd

import (
	"fmt"
	"os"
)

func (c *Cmd) addEvent(args []string) {
	if len(args) != 3 {
		fmt.Println(addFormatMsg)
		c.logger.Add(LogOutput, addFormatMsg)
		return
	}
	title, date, priority := args[0], args[1], args[2]

	e, err := c.calendar.AddEvent(title, date, priority)
	if err != nil {
		fmt.Println(err)
		c.logger.Add(LogOutput, err.Error())
		return
	}
	msg := fmt.Sprintf("Event '%s' has been added, ID: %s", e.Title, e.ID)
	fmt.Println(msg)
	c.logger.Add(LogOutput, msg)
}

func (c *Cmd) removeEvent(args []string) {
	if len(args) != 1 {
		fmt.Println(removeFormatMsg)
		c.logger.Add(LogOutput, removeFormatMsg)
		return
	}
	eventId := args[0]
	err := c.calendar.DeleteEvent(eventId)
	if err != nil {
		fmt.Println(err)
		c.logger.Add(LogOutput, err.Error())
		return
	}
	fmt.Println(removeSuccessMsg)
	c.logger.Add(LogOutput, removeSuccessMsg)
}

func (c *Cmd) updateEvent(args []string) {
	if len(args) != 4 {
		fmt.Println(updateFormatMsg)
		c.logger.Add(LogOutput, updateFormatMsg)
		return
	}
	id, name, date, priority := args[0], args[1], args[2], args[3]
	err := c.calendar.EditEvent(id, name, date, priority)
	if err != nil {
		fmt.Println(err)
		c.logger.Add(LogOutput, err.Error())
		return
	}
	fmt.Println(updateSuccessMsg)
	c.logger.Add(LogOutput, updateSuccessMsg)
}

func (c *Cmd) setReminder(args []string) {
	if len(args) != 3 {
		fmt.Println(setReminderFormatMsg)
		c.logger.Add(LogOutput, setReminderFormatMsg)
		return
	}
	id, message, at := args[0], args[1], args[2]
	err := c.calendar.SetEventReminder(id, message, at)
	if err != nil {
		fmt.Println(err)
		c.logger.Add(LogOutput, err.Error())
		return
	}
	fmt.Println(setReminderSuccessMsg)
	c.logger.Add(LogOutput, setReminderSuccessMsg)
}

func (c *Cmd) removeReminder(args []string) {
	if len(args) != 1 {
		fmt.Println(removeReminderFormatMsg)
		c.logger.Add(LogOutput, removeReminderFormatMsg)
		return
	}
	eventId := args[0]
	err := c.calendar.RemoveEventReminder(eventId)
	if err != nil {
		fmt.Println(err)
		c.logger.Add(LogOutput, err.Error())
		return
	}
	fmt.Println(removeReminderSuccessMsg)
	c.logger.Add(LogOutput, removeReminderSuccessMsg)
}

func (c *Cmd) cancelReminder(args []string) {
	if len(args) != 1 {
		fmt.Println(cancelReminderFormatMsg)
		c.logger.Add(LogOutput, cancelReminderFormatMsg)
		return
	}
	eventId := args[0]
	err := c.calendar.CancelEventReminder(eventId)
	if err != nil {
		fmt.Println(err)
		c.logger.Add(LogOutput, err.Error())
		return
	}
	fmt.Println(cancelReminderSuccessMsg)
	c.logger.Add(LogOutput, cancelReminderSuccessMsg)
}

func (c *Cmd) showEvents() {
	events := c.calendar.GetEvents()
	if len(events) == 0 {
		fmt.Println(noEventsMsg)
		c.logger.Add(LogOutput, noEventsMsg)
		return
	}
	for _, event := range events {
		fmt.Println(event.String())
	}
}

func (c *Cmd) showLogs() {
	logs := c.logger.GetLogs()
	if len(logs) == 0 {
		fmt.Println("no logs")
		return
	}
	for _, log := range logs {
		logLn := fmt.Sprintf("[%s] %s — %s", log.At.Format("2006-01-02 15:04:05"), string(log.Type), log.Message)
		fmt.Println(logLn)
	}
}

func (c *Cmd) exit() {
	err := c.calendar.Save()
	if err != nil {
		msg := fmt.Sprintf("failed to save calendar: %v", err)
		fmt.Println(msg)
		c.logger.Add(LogError, msg)
	} else {
		fmt.Println(savedCalendarMsg)
		c.logger.Add(LogOutput, savedCalendarMsg)
		c.logger.Save()
	}
	close(c.calendar.Notification)
	os.Exit(0)
}
