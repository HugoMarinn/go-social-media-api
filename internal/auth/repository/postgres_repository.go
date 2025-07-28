package repository

import (
	"database/sql"
	"errors"

	"github.com/HugoMarinn/go-social-media-api/internal/auth"
	"github.com/HugoMarinn/go-social-media-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type PostgresAuthRepository struct {
	db *sqlx.DB
}

func NewPostgresAuthRepository(db *sqlx.DB) auth.Repository {
	return &PostgresAuthRepository{db}
}

func (r *PostgresAuthRepository) EmailAlreadyTaken(email string) (bool, error) {
	var foundEmail string
	err := r.db.Get(
		&foundEmail,
		`SELECT email FROM users WHERE email = $1`,
		email,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil // email has not already taken
		}
		return false, err
	}

	return true, nil // email has already taken
}

func (r *PostgresAuthRepository) CreateUser(user *models.User) (*models.User, error) {
	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	err := r.db.QueryRowx(query, user.Name, user.Email, user.Password).
		Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}
