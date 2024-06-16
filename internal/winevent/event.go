package winevent

import "time"

type WinEvent struct {
	Timestamp time.Time
	Message   string
	Severity  EventSeverity
	Type      SubProvider
	Id        EventId
}

type EventSeverity uint32

const (
	Info EventSeverity = iota
	Warning
	Error
)

type EventId uint32

const (
	System_App_Start          EventId = 7705
	System_App_Error          EventId = 7706
	System_Decoy_File_Created EventId = 7707
	System_Decoy_File_Deleted EventId = 7708
)
