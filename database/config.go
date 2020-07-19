package database

import (
	"encoding/json"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"os"
)

type Config struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	SSLMode  string `json:"sslmode"`
}

func (c *Config) DatabaseConfig() *Config {
	return c
}

// Loads files optionally from the json file stored
// at path, then will override those values based on the envconfig
// struct tags. The envPrefix is how we prefix our environment
// variables.
func LoadConfig(path, envPrefix string, config interface{}) error {
	if path != "" {
		err := LoadFile(path, config)
		if err != nil {
			return errors.Wrap(err, "Error loading config file")
		}
	}
	err := envconfig.Process(envPrefix, config)
	return errors.Wrap(err, "Error loading config from env")
}

// Unmarshalls a json file into a config struct
func LoadFile(path string, config interface{}) error {
	configFile, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "Failed to read config file")
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	if err = decoder.Decode(config); err != nil {
		return errors.Wrap(err, "Failed to decode config file")
	}
	return nil
}
