package config

import (
	"os"
	"strconv"
)

func (cfg Config) readIntEnv(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		cfg.logger.Warn(
			"The environment variable is unset, use fallback instead",
			"key", key,
			"fallback", fallback,
		)
		return fallback
	}
	valueAsInt, err := strconv.Atoi(value)
	if err != nil {
		cfg.logger.Warn(
			"The environment variable is invalid, use fallback instead",
			"key", key,
			"value", value,
			"fallback", fallback,
		)
		return fallback
	}
	return valueAsInt
}
