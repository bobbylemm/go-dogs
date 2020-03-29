package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	)

type Config struct {
	AppPort string
	DbUrl string
}

func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		AppPort: os.Getenv("APP_PORT"),
		DbUrl:   os.Getenv("DB_URL"),
	}
}