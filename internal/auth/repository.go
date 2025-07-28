package auth

import "github.com/HugoMarinn/go-social-media-api/internal/models"

type Repository interface {
	CreateUser(user *models.User) (*models.User, error)
	EmailAlreadyTaken(email string) (bool, error)
}
