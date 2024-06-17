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

// Package process provides process level features such as OS signal handling
package process

import (
	"fmt"
	"github.com/buu-huu/purrsom-watch/internal/eventlog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	ExitDelay = 2000
)

var (
	logger = eventlog.GetLogger()
)

// HandleProcessTermination handles OS signals which ask the process to shut down (e.g. when the system shuts down)
// This closes all relevant queues, etc.
func HandleProcessTermination() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	go func() {
		sig := <-sigChan
		ev, err := eventlog.CreateEvent(eventlog.System_App_Shutdown_Signal, fmt.Sprintf("Signal: %s", sig.String()))
		if err != nil {
			fmt.Println("Failed to create event.")
		}
		logger.Log(ev)
		cleanup()
		time.Sleep(ExitDelay * time.Millisecond)
		os.Exit(0)
	}()
}
