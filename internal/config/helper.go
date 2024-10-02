package config

import (
	"encoding/json"
	"errors"
	"os"
)

func writeConfig(config *Config) error {
	dat, err := json.Marshal(config)
	if err != nil {
		return errors.New("config file data wasn't marshaled")
	}
	return os.WriteFile(configFileName, dat, 0644)
}
