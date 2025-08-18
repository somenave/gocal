package tests

import (
	"github.com/somenave/eventsCalendar/events"
	"testing"
)

func TestPriority(t *testing.T) {
	priority := events.Priority("low")
	err := priority.Validate()
	if err != nil {
		t.Errorf("Exepected priority to be valid, got %v", err)
	}

	invalidPriority := events.Priority("asd")

	err = invalidPriority.Validate()
	if err == nil {
		t.Error("Expecting error for invalid priority, got none")
	}
}

func TestEventTitle(t *testing.T) {
	title := "Hello 123 Мир"
	isValid := events.IsValidTitle(title)
	if !isValid {
		t.Errorf("Exepected title '%v' to be valid", title)
	}

	shortTitle := "H"
	isValid = events.IsValidTitle(shortTitle)
	if isValid {
		t.Errorf("Exepected title '%v' to be invalid", shortTitle)
	}

	invalidTitle := "Hello!!!"
	isValid = events.IsValidTitle(invalidTitle)
	if isValid {
		t.Errorf("Exepected title '%v' to be invalid", invalidTitle)
	}
}

func TestEvent(t *testing.T) {
	event, err := events.NewEvent("Hello", "2025-01-02 10:20", "low")
	if err != nil {
		t.Errorf("Exepected event to be valid, got %v", err)
	}
	newTitle := "World"
	newPriority := events.Priority("high")
	err = event.Update(newTitle, "2025-01-02 10:20", newPriority)
	if err != nil {
		t.Errorf("Exepected event successfully update, got %v", err)
	}
	if event.Title != newTitle {
		t.Errorf("Exepted event title is '%s' after update, got title is '%s'", newTitle, event.Title)
	}
	if event.Priority != newPriority {
		t.Errorf("Exepected event priority is '%s' after update, got priority is '%s'", newPriority, event.Priority)
	}

	_, err = events.NewEvent("E", "2025", "low")
	if err == nil {
		t.Errorf("Exepected title invalid error, got none")
	}

	_, err = events.NewEvent("Event", "2025/80", "low")
	if err == nil {
		t.Errorf("Exepected date invalid error, got none")
	}

	_, err = events.NewEvent("Event", "2025", "lowest")
	if err == nil {
		t.Errorf("Exepected priority invalid error, got none")
	}
}
