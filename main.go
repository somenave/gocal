package main

import (
	"github.com/somenave/eventsCalendar/calendar"
	"github.com/somenave/eventsCalendar/cmd"
	"github.com/somenave/eventsCalendar/storage"
)

func main() {
	cs := storage.NewJsonStorage("calendar.json")
	c := calendar.NewCalendar(cs)

	ls := storage.NewJsonStorage("calendar.log")
	l := cmd.NewLogger(ls)

	cli := cmd.NewCmd(c, l)
	cli.Run()
}
