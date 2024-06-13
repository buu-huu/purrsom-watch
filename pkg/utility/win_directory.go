package utility

import (
	"errors"
	"github.com/buu-huu/purrsom-watch/configs"
	"os/user"
	"path/filepath"
	"strings"
)

// CreateAbsoluteDirString creates an absolute directory string from userdir and the rest
func CreateAbsoluteDirString(config *configs.Config) (string, error) {
	var userDir string
	if config.PurrEngine.DecoyFile.Location.Username != "" {
		username, err := user.Lookup(config.PurrEngine.DecoyFile.Location.Username)
		if err != nil {
			return "", err
		}
		userDir = username.HomeDir
	}
	dirStringSplit := strings.SplitN(config.PurrEngine.DecoyFile.Location.FileDir, "%userdir%/", 2)
	if len(dirStringSplit) != 2 {
		return "", errors.New("invalid fileDir format in configuration file")
	}
	fileDirSpecific := dirStringSplit[1]
	fileDirAbsolute := filepath.Join(userDir, fileDirSpecific)

	//fmt.Println(fileDirAbsolute)
	return fileDirAbsolute, nil
}
