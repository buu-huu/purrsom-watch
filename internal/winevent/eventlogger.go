package winevent

import (
	"errors"
	"fmt"
	"golang.org/x/sys/windows/svc/eventlog"
	"strings"
	"syscall"
)

const (
	WinEventSourceName = "PurrsomWatch"
)

type SubSource int

const (
	System SubSource = iota
	Detection
)

func (s SubSource) String() string {
	return [...]string{"System", "Detection"}[s]
}

// InstallWinEventSource TODO: Unexport function after testing
func InstallWinEventSource() error {
	for _, subSource := range []SubSource{System, Detection} {
		err := eventlog.InstallAsEventCreate(subSource.String(), eventlog.Info|eventlog.Warning|eventlog.Error)
		if err != nil {
			// Trying to parse errors here to generic errors for proper logging...
			var errno syscall.Errno
			if errors.As(err, &errno) && errors.Is(errno, syscall.ERROR_ACCESS_DENIED) {
				fmt.Printf("Error installing Wineventlog subsource %s. Insufficient permissions: %s\n", subSource.String(), syscall.ERROR_ACCESS_DENIED)
				return err
			} else {
				// Fall back to string of the error message if not a permission problem
				if strings.Contains(err.Error(), "registry key already exists") {
					fmt.Printf("It appears, that Wineventlog subsource %s is already registered/installed: %s\n", subSource.String(), err.Error())
				} else {
					fmt.Printf("Unknown error registering/installing Wineventlog subsource %s: %s\n", subSource.String(), err.Error())
				}
			}
		} else {
			fmt.Printf("Wineventlog subsource %s installed successfully.\n", subSource.String())
		}
	}
	return nil
}
