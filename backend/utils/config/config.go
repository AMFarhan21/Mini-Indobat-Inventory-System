package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConnectionString string
	Port               string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error on loading env: %v", err.Error())
	}
	cfg := &Config{
		DBConnectionString: os.Getenv("DB_CONNECTION_STRING"),
		Port:               os.Getenv("PORT"),
	}

	return cfg
}
