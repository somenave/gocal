package events

import (
	"regexp"
)

func IsValidTitle(title string) bool {
	matched, err := regexp.MatchString("^[a-zA-Zа-яА-я0-9 ,.]{2,100}$", title)
	if err != nil {
		return false
	}
	return matched
}
