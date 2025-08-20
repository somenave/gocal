package main

import (
	"fmt"
	"github.com/somenave/eventsCalendar/calendar"
	"github.com/somenave/eventsCalendar/cmd"
	"github.com/somenave/eventsCalendar/logger"
	"github.com/somenave/eventsCalendar/storage"
)

func main() {
	cs := storage.NewJsonStorage("calendar.json")
	c := calendar.NewCalendar(cs)

	ls := storage.NewJsonStorage("calendar.log")
	l := cmd.NewLogger(ls)

	logFile, err := logger.Init("app.log")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer logFile.Close()

	cli := cmd.NewCmd(c, l)
	cli.Run()
}
