package configs

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
)

const (
	FileDirRegex = `^%userdir%[/\\].*$`
)

// Config struct defines the structure of the JSON configuration
type Config struct {
	PurrEngine struct {
		PurrInterval  string `json:"purrInterval"`
		Username      string `json:"username"`
		FileDir       string `json:"fileDir"`
		FileName      string `json:"fileName"`
		FileExtension string `json:"fileExtension"`
	} `json:"purrEngine"`
	WinEventProvider struct {
		EventId string `json:"eventId"`
	} `json:"winEventProvider"`
}

var Configuration *Config

func InitConfig(filePath string) error {
	Configuration, err := ParseConfig(filePath)
	if err != nil {
		return err
	}

	configLegitimate, err := IsConfigFileLegitimate(Configuration)
	if !configLegitimate {
		return errors.New("config file is not legitimate")
	}

	return nil
}

// ParseConfig parses the JSON configuration file and sets the config variable
func ParseConfig(filePath string) (*Config, error) {
	configFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	err = json.NewDecoder(configFile).Decode(&Configuration)
	if err != nil {
		return nil, err
	}

	return Configuration, nil
}

// PrintConfig prints the given configuration attributes
func PrintConfig(c *Config) error {
	if c == nil {
		return errors.New("configuration was not parsed")
	}

	configJSON, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(configJSON))

	return nil
}

// IsConfigParsed returns a boolean depending on if the config file was parsed
func IsConfigParsed(c *Config) bool {
	return c != nil
}

// IsConfigFileLegitimate checks validity of the provided config file
func IsConfigFileLegitimate(c *Config) (bool, error) {
	if !IsConfigParsed(c) {
		return false, errors.New("config file was not parsed")
	}
	re := regexp.MustCompile(FileDirRegex)
	if !re.MatchString(c.PurrEngine.FileDir) {
		return false, errors.New(fmt.Sprint("Config attribute FileDir does not match regex ", FileDirRegex))
	}
	return true, nil
}
