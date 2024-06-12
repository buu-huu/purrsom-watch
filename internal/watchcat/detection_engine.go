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

package watchcat

import (
	"fmt"
	"github.com/buu-huu/purrsom-watch/configs"
	"strconv"
	"time"
)

func Watch(config *configs.Config) {
	interval, err := strconv.Atoi(config.PurrEngine.PurrInterval)
	if err != nil {
		fmt.Println("Could not parse PurrInterval. Defaulting to 10s")
		interval = 10
	}
	ticker := time.Tick(time.Duration(interval) * time.Second)
	for range ticker {
		doSomething()
	}
}

func doSomething() {
	fmt.Println("Doing something...")
	// Your code to execute goes here
}
