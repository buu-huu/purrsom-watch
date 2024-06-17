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

// Package directory contains generic functions and helpers
package directory

import (
	"errors"
	"os/user"
	"path/filepath"
	"strings"
)

// CreateAbsoluteDirString creates an absolute directory string from userdir and the rest
// param username expects a string of the username
// param fileDir expects a string of the directory with a placeholder included
func CreateAbsoluteDirString(username string, fileDir string, placeholder string) (string, error) {
	var userDir string
	if username != "" {
		username, err := user.Lookup(username)
		if err != nil {
			return "", err
		}
		userDir = username.HomeDir
	}
	placeholderPrepared := placeholder + "/"
	dirStringSplit := strings.SplitN(fileDir, placeholderPrepared, 2)
	if len(dirStringSplit) != 2 {
		return "", errors.New("invalid fileDir format in configuration file")
	}
	fileDirSpecific := dirStringSplit[1]
	fileDirAbsolute := filepath.Join(userDir, fileDirSpecific)

	return fileDirAbsolute, nil
}
