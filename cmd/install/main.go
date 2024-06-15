package main

import (
	"fmt"
	"github.com/buu-huu/purrsom-watch/internal/winevent"
)

func main() {
	fmt.Println("Installing winevent providers for purrsom watch...")
	err := winevent.InstallWinEventProvider()
	if err != nil {
		fmt.Println("Error installing winevent provider.")
		return
	}
}
