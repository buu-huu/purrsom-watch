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

package main

import (
	"fmt"
	"github.com/buu-huu/purrsom-watch/configs"
	"github.com/buu-huu/purrsom-watch/internal/watchcat"
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

	err = watchcat.GenerateDecoyFile(configs.Configuration)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}