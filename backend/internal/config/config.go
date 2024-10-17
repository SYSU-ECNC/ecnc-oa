package config

import (
	"log/slog"
)

type Config struct {
	logger *slog.Logger

	DatabasePassword string
	JWTSecret        string
	ServerPort       int
}

func NewConfig(logger *slog.Logger) *Config {
	return &Config{logger: logger}
}

func (cfg *Config) LoadConfig() {
	cfg.DatabasePassword = cfg.readStringEnv("DATABASE_PASSWORD", "postgrespassword")
	cfg.JWTSecret = cfg.readStringEnv("JWT_SECRET", "d4c900b4994de272126318ffa6eb90be")
	cfg.ServerPort = cfg.readIntEnv("SERVER_PORT", 3000)
}
