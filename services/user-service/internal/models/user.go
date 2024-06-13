package models

import "time"

type User struct {
	ID            int64
    Email         string
    PasswordHash  string
    GoogleID      *string // Pointer to handle null values
    Name          string
    AvatarURL     string
    Provider      string // "google" or "local"
    AccessToken   *string
    RefreshToken  *string
    TokenExpiry   *time.Time
    CreatedAt     time.Time
    UpdatedAt     time.Time
}