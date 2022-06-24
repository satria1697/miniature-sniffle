package util

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DbName     string
	DbHost     string
	DbPassword string
	DbUsername string
	DbPort     string
}

func InitConfig() Config {
	_ = godotenv.Load()
	return Config{
		DbName:     os.Getenv("POSTGRES_NAME"),
		DbHost:     os.Getenv("POSTGRES_HOST"),
		DbPassword: os.Getenv("POSTGRES_PASSWORD"),
		DbUsername: os.Getenv("POSTGRES_USER"),
		DbPort:     os.Getenv("POSTGRES_PORT"),
	}
}
