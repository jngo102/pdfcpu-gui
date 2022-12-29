package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

const settingsFile = "settings.json"

// Application settings that are saved to and loaded from the local disk
type Settings struct {
	// Application language
	Language string `json:"language"`
	// Application theme
	Theme string `json:"theme"`
	// Opened PDFs
	OpenedPDFs []string `json:"openedPDFs"`
}

// Create new settings with default values
func NewSettings() Settings {
	return Settings{
		Language:   "en",
		Theme:      "light",
		OpenedPDFs: []string{},
	}
}

// Load application settings from local disk
func LoadSettings() (Settings, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Println("Failed to get user config dir: ", err)
		return Settings{}, err
	}
	settingsPath := filepath.Join(configDir, configFolder, settingsFile)
	file, err := os.OpenFile(settingsPath, os.O_RDONLY, 0666)
	if err != nil {
		log.Println("Failed to open/create settings file: ", err)
		return Settings{}, err
	}
	defer file.Close()

	// Decode settings from JSON
	decoder := json.NewDecoder(file)
	var settings Settings
	err = decoder.Decode(&settings)
	if err != nil {
		log.Println("Failed to decode settings file: ", err)
		return Settings{}, err
	}

	return settings, nil
}

// Save application settings to local disk
func (s *Settings) Save() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Failed to get user config dir: ", err)
		return err
	}
	settingsPath := filepath.Join(configDir, configFolder, settingsFile)
	settingsFile, err := os.OpenFile(settingsPath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Failed to open/create settings file: ", err)
		return err
	}
	defer settingsFile.Close()

	// Encode settings to JSON
	encoder := json.NewEncoder(settingsFile)
	encoder.SetIndent("", "\t")
	return encoder.Encode(s)
}
