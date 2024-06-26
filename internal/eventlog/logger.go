// MIT License
//
// Copyright (c) 2024 buuhuu
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package eventlog contains data and program logic for winevent logging
package eventlog

import (
	"errors"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"golang.org/x/sys/windows/svc/eventlog"
	"strings"
	"sync"
	"syscall"
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

var (
	globalEventLogger *EventLogger
	loggerOnce        sync.Once
)

type EventLogger struct{}

func NewEventLogger() *EventLogger {
	return &EventLogger{}
}

// InstallWinEventProvider installs the event provider for the application
func InstallWinEventProvider() error {
	for _, subProvider := range []SubProvider{System, Detection} {
		providerToInstall := fmt.Sprintf("%s-%s", ProviderName, subProvider.String())
		err := eventlog.InstallAsEventCreate(providerToInstall, eventlog.Info|eventlog.Warning|eventlog.Error)
		if err != nil {
			// Trying to parse access denied error here
			var errno syscall.Errno
			if errors.As(err, &errno) && errors.Is(errno, syscall.ERROR_ACCESS_DENIED) {
				fmt.Printf("Error installing eventlog log provider %s. Insufficient permissions: %s\n",
					providerToInstall,
					syscall.ERROR_ACCESS_DENIED)
				return err
			} else {
				// Fall back to string of the error message if not a permission problem
				if strings.Contains(err.Error(), "registry key already exists") {
					fmt.Printf("It appears, that eventlog log provider %s is already registered/installed: %s\n",
						providerToInstall,
						err.Error())
				} else {
					fmt.Printf("Unknown error registering/installing eventlog log provider %s: %s\n",
						providerToInstall,
						err.Error())
				}
			}
		} else {
			fmt.Printf("Winevent log provider %s installed successfully.\n", providerToInstall)
		}
	}
	return nil
}

// Log accepts an event and forwards it to the logging function
// Todo: Queueing and error handling here
func (e *EventLogger) Log(ev WinEvent) {
	fmt.Printf("Trying to log event %d\n", ev.Id)
	if err := e.logNow(ev); err != nil {
		fmt.Println("Failed to log event:", ev, err)
		// Todo: Queue for another try and error handling
	}
}

// logNow logs an event to the Windows Event
func (e *EventLogger) logNow(ev WinEvent) error {
	source := fmt.Sprintf("%s-%s", ProviderName, ev.Type.String())
	elog, err := eventlog.Open(source)
	if err != nil {
		fmt.Printf("Failed to open eventlog log for provider %s: %s\n", source, err.Error())
		return err
	}
	defer elog.Close()

	logMessage := fmt.Sprintf("%s - %s - %s - %s",
		ev.Timestamp.String(),
		ev.Severity.String(),
		ev.Message,
		strings.Join(ev.Details, ", "))

	switch ev.Severity {
	case Info:
		err = elog.Info(uint32(ev.Id), logMessage)
	case Warning:
		err = elog.Warning(uint32(ev.Id), logMessage)
	case Error:
		err = elog.Error(uint32(ev.Id), logMessage)
	default:
		err = elog.Info(uint32(ev.Id), logMessage)
		fmt.Printf("unknown ev severity: %d, defaulting to info severity\n", ev.Severity)
	}

	if err != nil {
		fmt.Println("Failed to write eventlog log ev:", err)
		return err
	}
	fmt.Printf("Successfully logged event: %s\n", ev.String())
	return nil
}

// IsEventProviderInstalled checks if the specified event provider is installed
func IsEventProviderInstalled(subProvider SubProvider) (bool, error) {
	source := fmt.Sprintf("%s-%s", ProviderName, subProvider.String())
	regKey := fmt.Sprintf(`SYSTEM\CurrentControlSet\Services\EventLog\Application\%s`, source)

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, regKey, registry.QUERY_VALUE)
	if err != nil {
		if errors.Is(err, registry.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	defer key.Close()

	return true, nil
}

// AreAllEventProvidersInstalled checks if all event providers are installed
func AreAllEventProvidersInstalled() (bool, error) {
	for _, subProvider := range []SubProvider{System, Detection} {
		installed, err := IsEventProviderInstalled(subProvider)
		if err != nil {
			ev, err := CreateEvent(System_App_Error, err)
			if err != nil {
				fmt.Println("Failed to create event:", err)
			}
			globalEventLogger.Log(ev)
			return false, err
		}
		if !installed {
			ev, err := CreateEvent(System_EventProviderNotInstalled, subProvider.String())
			if err != nil {
				fmt.Println("Failed to create event:", err)
			}
			globalEventLogger.Log(ev)
			return false, nil
		}
	}
	return true, nil
}

// init initializes the global eventlog logger
func init() {
	loggerOnce.Do(func() {
		globalEventLogger = NewEventLogger()
	})
}

// GetLogger returns a handle to the eventlog logger
func GetLogger() *EventLogger {
	return globalEventLogger
}
