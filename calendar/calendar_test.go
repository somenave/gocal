package calendar

import (
	"testing"
)

var testCalendar = NewCalendar(nil)

func TestAddEvent(t *testing.T) {
	_, err := testCalendar.AddEvent("Hello", "2025-01-02 10:20", "low")
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
	eventsCount := len(testCalendar.GetEvents())
	if eventsCount != 1 {
		t.Errorf("Expected 1 event but got %v", eventsCount)
	}

	_, err = testCalendar.AddEvent("Hello2", "2025-01-02 10:20", "low")
	if err == nil {
		t.Errorf("Expected error about event existence with same time, got nil")
	}
}

func TestRemoveEvent(t *testing.T) {
	e, _ := testCalendar.AddEvent("Hello", "2025-01-02 10:20", "low")
	err := testCalendar.DeleteEvent(e.ID)
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
	eventsCount := len(testCalendar.GetEvents())
	if eventsCount != 0 {
		t.Errorf("Expected 0 events but got %v", eventsCount)
	}
	err = testCalendar.DeleteEvent("no-existing-id")
	if err == nil {
		t.Errorf("Expected error about event does not exist, got nil")
	}
}
