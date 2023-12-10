package config

import (
	"errors"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const ENV_PREFIX = "gocmdclient"

type Config struct {
	BinPth string `envconfig:"BIN_PATH"`
}

// Recieve configuration values from env variables
func InitConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("Error with config initialization")
	}

	var cfg Config
	if err := envconfig.Process(ENV_PREFIX, &cfg); err != nil {
		return nil, errors.New("Error with config initialization")
	}

	return &cfg, nil
}
