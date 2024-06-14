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

// Package watchcat contains the program logic for ransomware detection. This includes decoy file generation and
// the detection engine itself
package watchcat

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/buu-huu/purrsom-watch/configs"
	"github.com/buu-huu/purrsom-watch/data/decoy"
	"github.com/buu-huu/purrsom-watch/pkg/util/directory"
	"github.com/buu-huu/purrsom-watch/pkg/util/ransomware/cryptoanalysis"
	"os"
	"path/filepath"
	"time"
)

// DecoyFileHandle is an exported pointer to the decoy file
var DecoyFileHandle *DecoyFile

// DecoyFile struct holds metadata and a handle to the actual file
type DecoyFile struct {
	File        *os.File
	FilePath    string
	Entropy     float64
	Data        *decoy.DataOriginal
	SizeKB      float64
	CreatedTime time.Time
}

// writeDataToFile writes data to the decoy file
func (file *DecoyFile) writeDataToFile() error {
	if file.File == nil {
		return errors.New("DecoyFile is nil")
	}
	data, err := hex.DecodeString(file.Data.DataHex)
	if err != nil {
		return err
	}
	_, err = file.File.Write(data)
	if err != nil {
		return err
	}
	fmt.Printf("Wrote hex data to file %s\n", file.File.Name())
	file.Entropy = cryptoanalysis.Entropy(data)

	return nil
}

// GenerateDecoyFile generates a decoy file with parameters from the config
func GenerateDecoyFile(config *configs.Config) error {
	if !configs.IsConfigParsed(config) {
		return errors.New("config not parsed")
	}

	if DecoyFileHandle == nil {
		DecoyFileHandle = &DecoyFile{}
	}

	fileDir, err := directory.CreateAbsoluteDirString(
		configs.Configuration.PurrEngine.DecoyFile.Location.Username,
		configs.Configuration.PurrEngine.DecoyFile.Location.FileDir,
		configs.PlaceholderUserdir)
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
	DecoyFileHandle.FilePath = filepath.Join(
		fileDir,
		fmt.Sprintf("%s.%s", config.PurrEngine.DecoyFile.FileName, config.PurrEngine.DecoyFile.FileExtension))

	DecoyFileHandle.Data = decoy.Data01
	file, err := os.Create(DecoyFileHandle.FilePath)
	defer file.Close()

	if err != nil {
		return err
	}
	DecoyFileHandle.File = file
	fmt.Println("Created file")

	err = DecoyFileHandle.writeDataToFile()
	fileStats, err := DecoyFileHandle.File.Stat()
	if err != nil {
		return err
	}

	DecoyFileHandle.SizeKB = float64(fileStats.Size()) / (1 << 10)
	DecoyFileHandle.CreatedTime = time.Now()

	return nil
}
