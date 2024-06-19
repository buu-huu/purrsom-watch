package configs

import (
	"os"
	"testing"
)

// Global variable for the valid configuration JSON
var validConfigJSON = `{
	"purrEngine": {
		"purrInterval": "4",
		"decoyFile": {
			"fileName": "meow",
			"fileExtension": "docx",
			"location": {
				"fileDir": "%userdir%/Documents/purr",
				"username": "peter"
			}
		}
	},
	"winEventProvider": {
		"eventId": "7700"
	}
}`

// Helper function to create a temporary config file
func createTempConfigFile(t *testing.T, content string) *os.File {
	t.Helper()
	tmpFile, err := os.CreateTemp("", "config-*.json")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tmpFile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatal(err)
	}
	return tmpFile
}

// Test cases for InitConfig
func TestInitConfig(t *testing.T) {
	invalidConfig := `{
		"purrEngine": {
			"purrInterval": "4",
			"decoyFile": {
				"fileName": "meow",
				"fileExtension": "docx",
				"location": {
					"fileDir": "invalidDir/Documents/purr",
					"username": "peter"
				}
			}
		},
		"winEventProvider": {
			"eventId": "7700"
		}
	}`

	tests := []struct {
		name          string
		configContent string
		expectError   bool
	}{
		{"ValidConfig", validConfigJSON, false},
		{"InvalidConfig", invalidConfig, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpfile := createTempConfigFile(t, tt.configContent)
			defer os.Remove(tmpfile.Name())

			err := InitConfig(tmpfile.Name())
			if (err != nil) != tt.expectError {
				t.Errorf("InitConfig() error = %v, expectError %v", err, tt.expectError)
			}
		})
	}
}

// Test cases for ParseConfig
func TestParseConfig(t *testing.T) {
	tmpfile := createTempConfigFile(t, validConfigJSON)
	defer os.Remove(tmpfile.Name())

	config, err := ParseConfig(tmpfile.Name())
	if err != nil {
		t.Fatalf("ParseConfig() error = %v", err)
	}

	if config.PurrEngine.PurrInterval != "4" {
		t.Errorf("expected PurrInterval to be '4', got %s", config.PurrEngine.PurrInterval)
	}
	if config.PurrEngine.DecoyFile.Location.FileDir != "%userdir%/Documents/purr" {
		t.Errorf("expected FileDir to be '%%userdir%%/Documents/purr', got %s", config.PurrEngine.DecoyFile.Location.FileDir)
	}
	if config.WinEventProvider.EventId != "7700" {
		t.Errorf("expected EventId to be '7700', got %s", config.WinEventProvider.EventId)
	}
}

// Test cases for IsConfigFileLegitimate
func TestIsConfigFileLegitimate(t *testing.T) {
	validConfig := &Config{
		PurrEngine: struct {
			PurrInterval string `json:"purrInterval"`
			DecoyFile    struct {
				FileName      string `json:"fileName"`
				FileExtension string `json:"fileExtension"`
				Location      struct {
					FileDir  string `json:"fileDir"`
					Username string `json:"username"`
				} `json:"Location"`
			} `json:"decoyFile"`
		}{
			PurrInterval: "4",
			DecoyFile: struct {
				FileName      string `json:"fileName"`
				FileExtension string `json:"fileExtension"`
				Location      struct {
					FileDir  string `json:"fileDir"`
					Username string `json:"username"`
				} `json:"Location"`
			}{
				FileName:      "meow",
				FileExtension: "docx",
				Location: struct {
					FileDir  string `json:"fileDir"`
					Username string `json:"username"`
				}{
					FileDir:  "%userdir%/Documents/purr",
					Username: "peter",
				},
			},
		},
		WinEventProvider: struct {
			EventId string `json:"eventId"`
		}{
			EventId: "7700",
		},
	}

	invalidConfig := &Config{
		PurrEngine: struct {
			PurrInterval string `json:"purrInterval"`
			DecoyFile    struct {
				FileName      string `json:"fileName"`
				FileExtension string `json:"fileExtension"`
				Location      struct {
					FileDir  string `json:"fileDir"`
					Username string `json:"username"`
				} `json:"Location"`
			} `json:"decoyFile"`
		}{
			PurrInterval: "4",
			DecoyFile: struct {
				FileName      string `json:"fileName"`
				FileExtension string `json:"fileExtension"`
				Location      struct {
					FileDir  string `json:"fileDir"`
					Username string `json:"username"`
				} `json:"Location"`
			}{
				FileName:      "meow",
				FileExtension: "docx",
				Location: struct {
					FileDir  string `json:"fileDir"`
					Username string `json:"username"`
				}{
					FileDir:  "invalidDir/Documents/purr",
					Username: "peter",
				},
			},
		},
		WinEventProvider: struct {
			EventId string `json:"eventId"`
		}{
			EventId: "7700",
		},
	}

	tests := []struct {
		name     string
		config   *Config
		expected bool
	}{
		{"ValidConfig", validConfig, true},
		{"InvalidConfig", invalidConfig, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isLegitimate, err := IsConfigFileLegitimate(tt.config)
			if isLegitimate != tt.expected {
				t.Errorf("IsConfigFileLegitimate() = %v, expected %v, error = %v", isLegitimate, tt.expected, err)
			}
		})
	}
}
