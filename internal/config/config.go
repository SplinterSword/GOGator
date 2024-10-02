package config

import (
	"encoding/json"
	"errors"
	"os"
)

const configFileName = "gatorconfig.json"

type Config struct {
	DBURL       string `json:"db_url"`
	CurrentUser string `json:"current_user_name"`
}

func Read() (Config, error) {
	dat, err := os.ReadFile(configFileName)
	if err != nil {
		return Config{}, errors.New("config file not found")
	}
	var config Config
	err = json.Unmarshal(dat, &config)
	if err != nil {
		return Config{}, errors.New("config file data wasn't unmarshaled")
	}
	return config, nil
}

func (config *Config) SetUser(user string) error {
	config.CurrentUser = user
	return writeConfig(config)
}
