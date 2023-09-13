package service

import (
	"course-api/application/repository"
	"course-api/domain/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.AuthRepository
}

// NewAuthService create instance for auth service
//
// Params:
//   - repo *repository.AuthRepository
//
// Return:
//   - *AuthService
func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

// Login is method to check user data and send
//
// Params:
//   - repo *repository.AuthRepository
//
// Return:
//   - *string
//   - error
func (as *AuthService) Login(creds model.LoginRequest) (*string, error) {
	user, err := as.repo.Login(creds)
	if err != nil {
		return nil, err
	}

	// check password matching
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Pass))
	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	secret := os.Getenv("SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &t, nil
}
