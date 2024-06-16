package main

import (
	"fmt"
	"github.com/buu-huu/purrsom-watch/internal/eventlog"
)

func main() {
	fmt.Println("Installing eventlog providers for purrsom watch...")
	err := eventlog.InstallWinEventProvider()
	if err != nil {
		fmt.Println("Error installing eventlog provider.")
		return
	}
}
