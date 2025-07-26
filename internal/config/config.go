package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	Env         string
	DatabaseURL string
}

func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.New(".env file not found")
	}

	return &Config{
		Port:        getOrDefault("PORT", "8080"),
		Env:         getOrDefault("ENV", "development"),
		DatabaseURL: getOrDefault("DATABASE_URL", "postgres://user:password@localhost:5432/go_social_media"),
	}, nil
}

func getOrDefault(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}
