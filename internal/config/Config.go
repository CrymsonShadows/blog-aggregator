package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUserName = name
	err := write(cfg)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", homeDir, configFileName), nil
}

func Read() (Config, error) {
	configDir, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	file, err := os.Open(configDir)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func write(cfg *Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	configDir, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(configDir, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
