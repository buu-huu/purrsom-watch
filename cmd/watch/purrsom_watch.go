package main

import (
	"fmt"
	"github.com/buu-huu/purrsom-watch/configs"
	"github.com/buu-huu/purrsom-watch/pkg/watchcat"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: purrsom_watch <config_file_path>")
		return
	}

	configFilePath := os.Args[1]

	err := configs.InitConfig(configFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("CONFIGURATION:")
	err = configs.PrintConfig(configs.Configuration)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = watchcat.GenerateFile()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
