package winevent

import (
	"errors"
	"fmt"
	"golang.org/x/sys/windows/svc/eventlog"
	"strings"
	"syscall"
	"time"
)

const (
	WinEventSourceName string = "PurrsomWatch"
)

type SubSource int

const (
	System SubSource = iota
	Detection
)

func (s SubSource) String() string {
	return [...]string{"System", "Detection"}[s]
}

type WinEvent struct {
	Timestamp time.Time
	Message   string
	Severity  EventSeverity
	Type      SubSource
}

type EventSeverity uint32

const (
	Info EventSeverity = iota
	Warning
	Error
)

type EventID uint32

const (
	EventIDInfo    EventID = 7705
	EventIDWarning EventID = 7706
	EventIDError   EventID = 7707
)

// InstallWinEventSource TODO: Unexport function after testing
func InstallWinEventSource() error {
	for _, subSource := range []SubSource{System, Detection} {
		sourceToInstall := fmt.Sprintf("%s-%s", WinEventSourceName, subSource.String())
		err := eventlog.InstallAsEventCreate(sourceToInstall, eventlog.Info|eventlog.Warning|eventlog.Error)
		if err != nil {
			// Trying to parse errors here to generic errors for proper logging...
			var errno syscall.Errno
			if errors.As(err, &errno) && errors.Is(errno, syscall.ERROR_ACCESS_DENIED) {
				fmt.Printf("Error installing Wineventlog subsource %s. Insufficient permissions: %s\n", sourceToInstall, syscall.ERROR_ACCESS_DENIED)
				return err
			} else {
				// Fall back to string of the error message if not a permission problem
				if strings.Contains(err.Error(), "registry key already exists") {
					fmt.Printf("It appears, that Wineventlog subsource %s is already registered/installed: %s\n", sourceToInstall, err.Error())
				} else {
					fmt.Printf("Unknown error registering/installing Wineventlog subsource %s: %s\n", sourceToInstall, err.Error())
				}
			}
		} else {
			fmt.Printf("Wineventlog subsource %s installed successfully.\n", sourceToInstall)
		}
	}
	return nil
}

func Log(event WinEvent) error {
	source := fmt.Sprintf("%s-%s", WinEventSourceName, event.Type.String())
	elog, err := eventlog.Open(source)
	if err != nil {
		fmt.Printf("Failed to open winevent log for source %s: %s\n", source, err.Error())
		return err
	}
	defer elog.Close()

	var id EventID
	switch event.Severity {
	case Info:
		id = EventIDInfo
		err = elog.Info(uint32(id), event.Message)
	case Warning:
		id = EventIDWarning
		err = elog.Warning(uint32(id), event.Message)
	case Error:
		id = EventIDError
		err = elog.Error(uint32(id), event.Message)
	default:
		err = fmt.Errorf("unknown event severity: %d", event.Severity)
	}

	if err != nil {
		fmt.Println("Failed to write winevent log event:", err)
	}
	fmt.Printf("Successfully logged event: %s\n", event.Message)
	return nil
}
