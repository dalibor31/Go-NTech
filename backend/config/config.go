package config

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	DBPath string `json:"db_path"`
	Port   string `json:"port"`
}

var Config AppConfig

func LoadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}

	return nil
}

func CreateDefaultConfig(path string) error {
	defaultConfig := AppConfig{
		DBPath: "NTech.db",
		Port:   "8080",
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(defaultConfig)
}
