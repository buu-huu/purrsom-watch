package configs

import (
	"encoding/json"
	"os"
)

// Config struct defines the structure of the JSON configuration
type Config struct {
	PurrEngine struct {
		PurrInterval  string `json:"purrInterval"`
		FileDir       string `json:"fileDir"`
		FileName      string `json:"fileName"`
		FileExtension string `json:"fileExtension"`
	} `json:"purrEngine"`
	WinEventProvider struct {
		EventId string `json:"eventId"`
	} `json:"winEventProvider"`
}

// ParseConfig parses the JSON configuration file and returns a Config struct
func ParseConfig(filePath string) (Config, error) {
	configFile, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}

	defer func() {
		if cerr := configFile.Close(); cerr != nil {
			err = cerr
		}
	}()

	var config Config
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
