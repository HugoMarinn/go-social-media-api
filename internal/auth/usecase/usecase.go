package usecase

import (
	"errors"

	"github.com/HugoMarinn/go-social-media-api/internal/auth"
	"github.com/HugoMarinn/go-social-media-api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUnexpectedInRepo  = errors.New("we got an unexpected error consulting our repository")
	ErrEmailAlreadyTaken = errors.New("payload email has already taken")
)

type AuthUseCase struct {
	repo auth.Repository
}

func NewAuthUseCase(repo auth.Repository) auth.UseCase {
	return &AuthUseCase{repo}
}

func (uc *AuthUseCase) Register(payload *auth.RegisterRequestDTO) (*auth.RegisterResponseDTO, error) {
	emailAlreadyTaken, err := uc.repo.EmailAlreadyTaken(payload.Email)
	if err != nil {
		return nil, ErrUnexpectedInRepo
	}

	if emailAlreadyTaken {
		return nil, ErrEmailAlreadyTaken
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(payload.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	newUser, err := uc.repo.CreateUser(&models.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return nil, ErrUnexpectedInRepo
	}

	return &auth.RegisterResponseDTO{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
	}, nil
}
