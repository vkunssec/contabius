package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	ServerName = "contabius"
)

func Env(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env file not exists")
	}
	return os.Getenv(key)
}
