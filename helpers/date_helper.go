package helpers

import (
	"github.com/araddon/dateparse"
	"time"
)

func ParseDate(date string) (time.Time, error) {
	dateParsed, err := dateparse.ParseLocal(date)
	if err != nil {
		return time.Time{}, err
	}
	return dateParsed, nil
}
