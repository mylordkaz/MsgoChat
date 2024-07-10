package jwt

import "github.com/golang-jwt/jwt/v5"

var secretKey = []byte("secret-key") // in prod: NO HARDCODED

type claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}