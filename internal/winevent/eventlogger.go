package winevent

import (
	"errors"
	"fmt"
	"golang.org/x/sys/windows/svc/eventlog"
	"strings"
	"sync"
	"syscall"
	"time"
)

const (
	ProviderName string = "PurrsomWatch"
)

type SubProvider int

const (
	System SubProvider = iota
	Detection
)

func (s SubProvider) String() string {
	return [...]string{"System", "Detection"}[s]
}

type WinEvent struct {
	Timestamp time.Time
	Message   string
	Severity  EventSeverity
	Type      SubProvider
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

var (
	globalEventLogger *EventLogger
	once              sync.Once
)

type EventLogger struct{}

func NewEventLogger() *EventLogger {
	return &EventLogger{}
}

// InstallWinEventProvider installs the event provider for the application TODO: Unexport function after testing
func (e *EventLogger) InstallWinEventProvider() error {
	for _, subProvider := range []SubProvider{System, Detection} {
		providerToInstall := fmt.Sprintf("%s-%s", ProviderName, subProvider.String())
		err := eventlog.InstallAsEventCreate(providerToInstall, eventlog.Info|eventlog.Warning|eventlog.Error)
		if err != nil {
			// Trying to parse access denied error here
			var errno syscall.Errno
			if errors.As(err, &errno) && errors.Is(errno, syscall.ERROR_ACCESS_DENIED) {
				fmt.Printf("Error installing winevent log provider %s. Insufficient permissions: %s\n", providerToInstall, syscall.ERROR_ACCESS_DENIED)
				return err
			} else {
				// Fall back to string of the error message if not a permission problem
				if strings.Contains(err.Error(), "registry key already exists") {
					fmt.Printf("It appears, that winevent log provider %s is already registered/installed: %s\n", providerToInstall, err.Error())
				} else {
					fmt.Printf("Unknown error registering/installing winevent log provider %s: %s\n", providerToInstall, err.Error())
				}
			}
		} else {
			fmt.Printf("Winevent log provider %s installed successfully.\n", providerToInstall)
		}
	}
	return nil
}

// Log logs an event to the Windows Event Log
func (e *EventLogger) Log(event WinEvent) error {
	source := fmt.Sprintf("%s-%s", ProviderName, event.Type.String())
	elog, err := eventlog.Open(source)
	if err != nil {
		fmt.Printf("Failed to open winevent log for provider %s: %s\n", source, err.Error())
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
		return fmt.Errorf("unknown event severity: %d", event.Severity)
	}

	if err != nil {
		fmt.Println("Failed to write winevent log event:", err)
		return err
	}
	fmt.Printf("Successfully logged event: %s\n", event.Message)
	return nil
}

// init initializes the global winevent logger
func init() {
	once.Do(func() {
		globalEventLogger = NewEventLogger()
	})
}

// GetLogger returns a handle to the winevent logger
func GetLogger() *EventLogger {
	return globalEventLogger
}
