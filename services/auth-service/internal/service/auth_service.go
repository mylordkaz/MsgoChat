package service

import (
	"errors"
	"time"

	"github.com/mylordkaz/MsgoChat/services/auth-service/internal/models"
	"github.com/mylordkaz/MsgoChat/services/auth-service/internal/repository"
	"github.com/mylordkaz/MsgoChat/services/auth-service/pkg/utils/hash"
	"github.com/mylordkaz/MsgoChat/services/auth-service/pkg/utils/jwt"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) RegisterUser(user *models.User) error {
	hashedPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.repo.CreateUser(user)
}

func (s *AuthService) LoginUser(username, password string) (string, error) {
    user, err := s.repo.GetUserByUsername(username)
    if err != nil {
        return "", err
    }

    if !hash.CheckPasswordHash(password, user.Password) {
        return "", errors.New("invalid credentials")
    }

    return jwt.GenerateToken(user.ID, 24*time.Hour)
}

// add other service e.g : getuserbyid