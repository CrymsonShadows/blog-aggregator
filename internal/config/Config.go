package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbURL string `json:"db_url`
}

func Read() Config {
	configDir, err := getConfigFilePath()
	if err != nil {
		fmt.Println(err.Error())
		return Config{}
	}
	file, err := os.Open(configDir)
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
		return Config{}
	}
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err.Error())
		return Config{}
	}
	return config
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", homeDir, ".gatorconfig.json"), nil
}
