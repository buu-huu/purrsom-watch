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
	"strings"
	"time"
)

type WinEvent struct {
	Id        EventId
	Timestamp time.Time
	Message   string
	Details   []string
	Severity  EventSeverity
	Type      SubProvider
}

type EventSeverity uint32

const (
	Info EventSeverity = iota
	Warning
	Error
)

func (e EventSeverity) String() string {
	return [...]string{"Info", "Warning", "Error"}[e]
}

// CreateEvent creates a WinEvent from a list of templates
func CreateEvent(id EventId, details ...interface{}) (WinEvent, error) {
	event, exists := eventTemplate[id]
	if !exists {
		return WinEvent{}, errors.New("event id does not exist")
	}
	event.Timestamp = time.Now()

	if len(details) > 0 {
		var detailStrings []string
		for _, d := range details {
			detailStrings = append(detailStrings, strings.TrimSuffix(fmt.Sprint(d), "\n"))
		}
		event.Details = detailStrings
		//event.Message = fmt.Sprintf("%s, %s", event.Message, strings.Join(detailStrings, ", "))
	}
	return event, nil
}

func (ev WinEvent) String() string {
	return fmt.Sprintf("%s - %s - %s - %s",
		ev.Timestamp.String(),
		ev.Severity.String(),
		ev.Message,
		strings.Join(ev.Details, ", "))
}
