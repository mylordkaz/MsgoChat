package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"regexp"
	"strings"

	"github.com/mylordkaz/MsgoChat/services/user-service/internal/models"
)

func GenerateRandomID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func ValidateUser(user *models.User) error {
	// check username
	if strings.TrimSpace(user.Username) == "" {
		return errors.New("username cannot be empty")
	}
	if len(user.Username) < 3 || len(user.Username) > 30 {
		return errors.New("username must be between 3 to 30 characters")
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(user.Username) {
		return errors.New("username can only contain alphanumeric characters and underscore")
	}

	// check email
	if strings.TrimSpace(user.Email) == "" {
		return errors.New("email cannot be empty")
	}
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(user.Email) {
		return errors.New("invalid email format")
	}

	return nil
}