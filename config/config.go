package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	BasePath string
	Port int
}

func (config *Config) Default() {
	config.BasePath = "/"
	config.Port = 8080
}

func (config *Config) LoadFromEnvironmentVariables() error {
	if s, exists := os.LookupEnv("PORT"); exists {
		var err error
		if config.Port, err = strconv.Atoi(s); err != nil {
			return fmt.Errorf("PORT error: %v", err)
		}
	}

	if s, exists := os.LookupEnv("BASE_PATH"); exists {
		config.BasePath = s
	}
	return nil
}