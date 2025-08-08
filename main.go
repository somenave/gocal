package main

import (
	"github.com/somenave/eventsCalendar/calendar"
	"github.com/somenave/eventsCalendar/cmd"
	"github.com/somenave/eventsCalendar/storage"
)

func main() {
	s := storage.NewJsonStorage("calendar.json")
	c := calendar.NewCalendar(s)

	cli := cmd.NewCmd(c)
	cli.Run()
}
