package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string
	DB   *sqlx.DB
}

func New() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.New(".env file not found")
	}

	dsn := getEnvOrDefault("DATABASE_URL", "postgres://user:password@localhost:5432/go_social_media")
	db, err := setupDatabase(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed DB setup: %v", err)
	}

	return &Config{
		Port: getEnvOrDefault("PORT", "8080"),
		Env:  getEnvOrDefault("ENV", "development"),
		DB:   db,
	}, nil
}

func getEnvOrDefault(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}
