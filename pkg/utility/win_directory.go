package utility

import (
	"errors"
	"os/user"
	"path/filepath"
	"strings"
)

// CreateAbsoluteDirString creates an absolute directory string from userdir and the rest
func CreateAbsoluteDirString(username string, fileDir string) (string, error) {
	var userDir string
	if username != "" {
		username, err := user.Lookup(username)
		if err != nil {
			return "", err
		}
		userDir = username.HomeDir
	}
	dirStringSplit := strings.SplitN(fileDir, "%userdir%/", 2)
	if len(dirStringSplit) != 2 {
		return "", errors.New("invalid fileDir format in configuration file")
	}
	fileDirSpecific := dirStringSplit[1]
	fileDirAbsolute := filepath.Join(userDir, fileDirSpecific)

	//fmt.Println(fileDirAbsolute)
	return fileDirAbsolute, nil
}
