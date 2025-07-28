package config

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

func setupDatabase() (*sqlx.DB, error) {
	db, err := connectDB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect DB: %v", err)
	}

	if err := runMigrations(db); err != nil {
		return nil, fmt.Errorf("failed to run DB migrations: %v", err)
	}

	return db, nil
}

func connectDB() (*sqlx.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func runMigrations(db *sqlx.DB) error {
	goose.SetDialect("postgres")
	return goose.Up(db.DB, "./migrations/production")
}
