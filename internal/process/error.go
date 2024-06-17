package process

import (
	"fmt"
	"github.com/buu-huu/purrsom-watch/internal/eventlog"
	"strings"
)

func HandleError(err error, details ...string) {
	if err == nil {
		fmt.Println("Error is nil. Nothing to handle")
		return
	}
	detailsJoined := strings.Join(details, ", ")
	ev, err := eventlog.CreateEvent(eventlog.System_App_Error, err, detailsJoined)
	if err != nil {
		fmt.Println("Failed to log event:", ev)
	}
	eventlog.GetLogger().Log(ev)
}
