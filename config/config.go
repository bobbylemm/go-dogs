package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	)

type Config struct {
	AppPort string
	DBUrl string
}

func GetConfig() *Config {
	return &Config{
		AppPort: os.Getenv("APP_PORT"),
		DBUrl:   os.Getenv("DB_URL"),
	}
}