package configs

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/vkunssec/contabius/utils/logger"
)

const (
	ServerName = "contabius"
)

var (
	Host = GetEnvOrDefault("HOST", "localhost:8080")
)

func Env(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Logger.Error().Err(err).Msg(".env file not exists")
	}
	return os.Getenv(key)
}

func GetEnvOrDefault(key, defaultValue string) string {
	if value := Env(key); value != "" {
		return value
	}
	return defaultValue
}
