package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Logger `yaml:"logger"`
	Db     `yaml:"postgres"`
}

type Logger struct {
	LogLevel string `yaml:"log-level"`
	LogFile  string `yaml:"log-file"`
}

type Db struct {
}

func ReadConfig() (*Config, error) {
	cfg := Config{}
	err := cleanenv.ReadConfig("config/config.yaml", &cfg)
	if err != nil {
		return nil, fmt.Errorf("Read config error %v", err.Error)
	}
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, fmt.Errorf("Read env error %v", err.Error)
	}

	return &cfg, nil
}
