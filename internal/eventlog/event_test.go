package eventlog

import (
	"testing"
	"time"
)

func TestCreateEvent(t *testing.T) {
	id := System_App_Start
	details := []interface{}{"detail1", "detail2"}

	event, err := CreateEvent(id, details...)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if event.Id != id {
		t.Errorf("expected event ID %d, got %d", id, event.Id)
	}

	if event.Message != "Application started" {
		t.Errorf("expected message 'Application started', got %s", event.Message)
	}

	expectedDetails := []string{"detail1", "detail2"}
	for i, detail := range event.Details {
		if detail != expectedDetails[i] {
			t.Errorf("expected detail %s, got %s", expectedDetails[i], detail)
		}
	}

	if event.Severity != Info {
		t.Errorf("expected severity Info, got %v", event.Severity)
	}
}

func TestCreateEventInvalidId(t *testing.T) {
	id := EventId(9999) // Invalid ID

	_, err := CreateEvent(id)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	expectedError := "event id does not exist"
	if err.Error() != expectedError {
		t.Errorf("expected error message '%s', got '%s'", expectedError, err.Error())
	}
}

func TestWinEventString(t *testing.T) {
	event := WinEvent{
		Id:        System_App_Start,
		Timestamp: time.Date(2024, time.June, 23, 12, 0, 0, 0, time.UTC),
		Message:   "Application started",
		Details:   []string{"detail1", "detail2"},
		Severity:  Info,
		Type:      System,
	}

	expectedString := "2024-06-23 12:00:00 +0000 UTC - Info - Application started - detail1, detail2"
	if event.String() != expectedString {
		t.Errorf("expected '%s', got '%s'", expectedString, event.String())
	}
}

func TestEventSeverityString(t *testing.T) {
	tests := []struct {
		severity EventSeverity
		expected string
	}{
		{Info, "Info"},
		{Warning, "Warning"},
		{Error, "Error"},
	}

	for _, test := range tests {
		if test.severity.String() != test.expected {
			t.Errorf("expected '%s', got '%s'", test.expected, test.severity.String())
		}
	}
}
