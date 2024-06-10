package watchcat

import (
	"errors"
	"fmt"
	"github.com/buu-huu/purrsom-watch/configs"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func GenerateFile() error {
	if !configs.IsConfigParsed(configs.Configuration) {
		return errors.New("config not parsed")
	}

	fileDir, err := CreateAbsoluteDirString()
	if err != nil {
		return err
	}

	_, err = os.Stat(fileDir)
	if err == nil {
		fmt.Println(fileDir, "exists")
	} else if os.IsNotExist(err) {
		return errors.New(fileDir + ": directory does not exist")
	} else {
		return err
	}

	return nil
}

func CreateAbsoluteDirString() (string, error) {
	var userDir string
	if configs.Configuration.PurrEngine.Username != "" {
		username, err := user.Lookup(configs.Configuration.PurrEngine.Username)
		if err != nil {
			return "", err
		}
		userDir = username.HomeDir
	}
	dirStringSplit := strings.SplitN(configs.Configuration.PurrEngine.FileDir, "%userdir%/", 2)
	if len(dirStringSplit) != 2 {
		return "", errors.New("invalid fileDir format in configuration file")
	}
	fileDirSpecific := dirStringSplit[1]
	fileDirAbsolute := filepath.Join(userDir, fileDirSpecific)

	//fmt.Println(fileDirAbsolute)
	return fileDirAbsolute, nil
}
