package usecase

import (
	"testing"
	"time"

	"github.com/HugoMarinn/go-social-media-api/internal/auth"
	"github.com/HugoMarinn/go-social-media-api/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockAuthRepository struct {
	mock.Mock
}

func (m *mockAuthRepository) EmailAlreadyTaken(email string) (bool, error) {
	args := m.Called(email)
	return args.Bool(0), args.Error(1)
}

func (m *mockAuthRepository) CreateUser(user *models.User) (*models.User, error) {
	args := m.Called(user)
	return args.Get(0).(*models.User), args.Error(1)
}

func TestAuthUseCase_Register(t *testing.T) {
	repo := new(mockAuthRepository)
	usecase := NewAuthUseCase(repo)
	responseDTOCreatedAt := time.Now()
	reponseDTOID := uuid.New()

	testCases := []struct {
		name           string
		payload        *auth.RegisterRequestDTO
		expectedReturn *auth.RegisterResponseDTO
		expectedErr    error
	}{
		{
			name: "fail when pass a email already taken",
			payload: &auth.RegisterRequestDTO{
				Name:     "John Doe",
				Email:    "email@gmail.com",
				Password: "something like this",
			},
			expectedReturn: nil,
			expectedErr:    ErrEmailAlreadyTaken,
		},
		{
			name: "register user successfully",
			payload: &auth.RegisterRequestDTO{
				Name:     "A person",
				Email:    "email@outlook.com",
				Password: "another thing like that",
			},
			expectedReturn: &auth.RegisterResponseDTO{
				ID:        reponseDTOID,
				Name:      "A person",
				Email:     "email@outlook.com",
				CreatedAt: responseDTOCreatedAt,
			},
			expectedErr: nil,
		},
	}

	repo.On("EmailAlreadyTaken", "email@gmail.com").Return(true, nil)
	repo.On("EmailAlreadyTaken", "email@outlook.com").Return(false, nil)
	repo.On("CreateUser", mock.AnythingOfType("*models.User")).Return(&models.User{
		ID:        reponseDTOID,
		Name:      "A person",
		Email:     "email@outlook.com",
		CreatedAt: responseDTOCreatedAt,
	}, nil)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := usecase.Register(tc.payload)
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expectedReturn, resp)
		})
	}
}
