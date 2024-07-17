package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/markbates/goth"
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
	
	if user.Provider != "local" {
        return "", errors.New("this account uses OAuth, please login with the appropriate provider")
    }


    if !hash.CheckPasswordHash(password, user.Password) {
        return "", errors.New("invalid credentials")
    }

    return jwt.GenerateToken(user.ID, 24*time.Hour)
}

func (s *AuthService) LoginOrRegisterOAuthUser(gothUser goth.User) (string, error) {
	var user *models.User
	var err error
	user, err = s.repo.GetUserByProviderId(gothUser.Provider, gothUser.UserID)
	if err != nil {
		user, err = s.repo.GetUserByEmail(gothUser.Email)
		if err != nil {
			newUser := &models.User{
				Username: gothUser.NickName,
				Email: gothUser.Email,
				Provider: gothUser.Provider,
			}

			switch gothUser.Provider {
			case "google":
				newUser.GoogleID = gothUser.UserID
			case "github":
				newUser.GithubID = gothUser.UserID
			}

			err = s.repo.CreateUser(newUser)
			if err != nil {
				return "", err
			}
			user = newUser
		} else {
			return "", fmt.Errorf("email already in use with another account")
		}
	}
	return jwt.GenerateToken(user.ID, 24*time.Hour)
}
