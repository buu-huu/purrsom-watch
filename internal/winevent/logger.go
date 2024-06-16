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

package winevent

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
	once              sync.Once
)

type EventLogger struct{}

func NewEventLogger() *EventLogger {
	return &EventLogger{}
}

// InstallWinEventProvider installs the event provider for the application TODO: Unexport function after testing
func InstallWinEventProvider() error {
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

	switch event.Severity {
	case Info:
		err = elog.Info(uint32(event.Id), event.Message)
	case Warning:
		err = elog.Warning(uint32(event.Id), event.Message)
	case Error:
		err = elog.Error(uint32(event.Id), event.Message)
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
			return false, err
		}
		if !installed {
			fmt.Printf("Event provider %s is not installed.\n", subProvider.String())
			return false, nil
		}
	}
	return true, nil
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
