package events

import (
	"testing"
)

func TestPriority(t *testing.T) {
	priority := Priority("low")
	err := priority.Validate()
	if err != nil {
		t.Errorf("Exepected priority to be valid, got %v", err)
	}

	invalidPriority := Priority("asd")

	err = invalidPriority.Validate()
	if err == nil {
		t.Error("Expecting error for invalid priority, got none")
	}
}

func TestEventTitle(t *testing.T) {
	title := "Hello 123 Мир"
	isValid := IsValidTitle(title)
	if !isValid {
		t.Errorf("Exepected title '%v' to be valid", title)
	}

	shortTitle := "H"
	isValid = IsValidTitle(shortTitle)
	if isValid {
		t.Errorf("Exepected title '%v' to be invalid", shortTitle)
	}

	invalidTitle := "Hello!!!"
	isValid = IsValidTitle(invalidTitle)
	if isValid {
		t.Errorf("Exepected title '%v' to be invalid", invalidTitle)
	}
}

func TestEvent(t *testing.T) {
	event, err := NewEvent("Hello", "2025-01-02 10:20", "low")
	if err != nil {
		t.Errorf("Exepected event to be valid, got %v", err)
	}
	newTitle := "World"
	newPriority := Priority("high")
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

	_, err = NewEvent("E", "2025", "low")
	if err == nil {
		t.Errorf("Exepected title invalid error, got none")
	}

	_, err = NewEvent("Event", "2025/80", "low")
	if err == nil {
		t.Errorf("Exepected date invalid error, got none")
	}

	_, err = NewEvent("Event", "2025", "lowest")
	if err == nil {
		t.Errorf("Exepected priority invalid error, got none")
	}
}
