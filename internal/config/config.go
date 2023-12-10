package config

import (
	"errors"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	ENV_PREFIX = "gocmdclient"

	// TODO: Move to yml config
	TASK_TIMEOUT = time.Second * 3
)

type Config struct {
	BinPth      string `envconfig:"BIN_PATH"`
	TaskTimeout time.Duration
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
	cfg.TaskTimeout = TASK_TIMEOUT

	return &cfg, nil
}
