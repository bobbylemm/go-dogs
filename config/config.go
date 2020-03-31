package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	)

type Config struct {
	AppPort string
	DBUrl string
}

func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		AppPort: os.Getenv("APP_PORT"),
		DBUrl:   os.Getenv("DB_URL"),
	}
}