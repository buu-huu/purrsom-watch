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

// Package main contains the main function and is therefore pretty slim
package main

import (
	"fmt"
	"github.com/buu-huu/purrsom-watch/configs"
	"github.com/buu-huu/purrsom-watch/internal/watchcat"
	"github.com/buu-huu/purrsom-watch/internal/winevent"
	"os"
	"time"
)

// main is the entry point for the application. It handles parsing of command line arguments and error
// handling. It calls the corresponding submodules of the program
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: purrsom_watch <config_file_path>")
		return
	}

	eventlogger := winevent.GetLogger()
	providerCheck, err := winevent.AreAllEventProvidersInstalled()
	if !providerCheck {
		fmt.Println("Winevent log providers not installed! Install providers first using install script.")
		return
	}
	if err != nil {
		fmt.Println("Error checking if winevent providers are installed: ", err)
		return
	}

	configFilePath := os.Args[1]
	err = configs.InitConfig(configFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("CONFIGURATION:")
	err = configs.PrintConfig(configs.Configuration)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = watchcat.GenerateDecoyFile(configs.Configuration)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//watchcat.Watch(configs.Configuration)

	event := winevent.WinEvent{
		Timestamp: time.Now(),
		Message:   "Decoy file created",
		Severity:  winevent.Info,
		Type:      winevent.System,
		Id:        winevent.System_Decoy_File_Created,
	}

	err = eventlogger.Log(event)
	if err != nil {
		fmt.Println(err)
	}
}
