package config

import (
	"log"
	"os"
)

type Config struct {
	DatabaseURL string
}

func NewConfig() *Config {
	connURL := os.Getenv("POSTGRESQL_URL")
	if connURL == "" {
		log.Fatal("Empty database connection string")
	}
	return &Config{DatabaseURL: connURL}
}
