package config

import (
	"log/slog"
)

type Config struct {
	logger *slog.Logger

	ServerPort int
}

func NewConfig(logger *slog.Logger) *Config {
	return &Config{logger: logger}
}

func (cfg *Config) LoadConfig() {
	cfg.ServerPort = cfg.readIntEnv("SERVER_PORT", 3000)
}
