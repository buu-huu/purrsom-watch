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
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/buu-huu/purrsom-watch/configs"
	"github.com/buu-huu/purrsom-watch/data/decoy"
	"github.com/buu-huu/purrsom-watch/pkg/utility"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type DecoyFile struct {
	File    *os.File
	Entropy float64
	SizeKB  float64
}

var DecoyFileHandle DecoyFile

func GenerateDecoyFile(config *configs.Config) error {
	if !configs.IsConfigParsed(config) {
		return errors.New("config not parsed")
	}

	fileDir, err := CreateAbsoluteDirString(config)
	if err != nil {
		return err
	}

	// Create watch dir if not exists
	_, err = os.Stat(fileDir)
	if os.IsNotExist(err) {
		// Create Directory
		fmt.Println("Provided directory doesn't exist. Creating...")
		err := os.MkdirAll(fileDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Create decoy file
	file, err := os.Create(filepath.Join(fileDir, config.PurrEngine.FileName))
	if err != nil {
		return err
	}
	DecoyFileHandle.File = file
	fmt.Println("Created file")

	err = WriteDecoyFile(&DecoyFileHandle)
	if err != nil {
		return err
	}
	return nil
}

func WriteDecoyFile(fileWithEntropy *DecoyFile) error {
	if fileWithEntropy.File == nil {
		return errors.New("DecoyFile is nil")
	}
	data, err := hex.DecodeString(decoy.HexDecoy)
	if err != nil {
		return err
	}

	_, err = fileWithEntropy.File.Write(data)
	if err != nil {
		return err
	}

	fmt.Println("Wrote hex to decoy file")
	fileWithEntropy.Entropy = utility.Entropy(data)
	fmt.Println("Entropy:", fileWithEntropy.Entropy)
	return nil
}

func CreateAbsoluteDirString(config *configs.Config) (string, error) {
	var userDir string
	if configs.Configuration.PurrEngine.Username != "" {
		username, err := user.Lookup(config.PurrEngine.Username)
		if err != nil {
			return "", err
		}
		userDir = username.HomeDir
	}
	dirStringSplit := strings.SplitN(config.PurrEngine.FileDir, "%userdir%/", 2)
	if len(dirStringSplit) != 2 {
		return "", errors.New("invalid fileDir format in configuration file")
	}
	fileDirSpecific := dirStringSplit[1]
	fileDirAbsolute := filepath.Join(userDir, fileDirSpecific)

	//fmt.Println(fileDirAbsolute)
	return fileDirAbsolute, nil
}
