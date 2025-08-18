package tests

import (
	"github.com/somenave/eventsCalendar/reminder"
	"testing"
	"time"
)

func notifyForTest(s string) {}

func TestReminder(t *testing.T) {
	reminderValidTime := time.Now().Add(time.Second * 10)
	r, err := reminder.NewReminder("!", reminderValidTime.String(), notifyForTest)
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}
	if r.Sent == true {
		t.Error("Expected reminder not to be sent after creation")
	}
	r.Send()
	if r.Sent == false {
		t.Error("Expected reminder to be sent")
	}

	_, err = reminder.NewReminder("", time.Now().String(), notifyForTest)
	if err == nil {
		t.Errorf("Expected error about date should be in future but got nil")
	}
}
