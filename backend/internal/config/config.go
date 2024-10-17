package config

import (
	"log/slog"
)

type Config struct {
	logger *slog.Logger

	ServerPort       int
	DatabasePassword string
}

func NewConfig(logger *slog.Logger) *Config {
	return &Config{logger: logger}
}

func (cfg *Config) LoadConfig() {
	cfg.ServerPort = cfg.readIntEnv("SERVER_PORT", 3000)
	cfg.DatabasePassword = cfg.readStringEnv("DATABASE_PASSWORD", "postgrespassword")
}
