package handlers

import (
	"net/http"

	"github.com/mylordkaz/MsgoChat/services/auth-service/pkg/database"
)

type AuthHandler struct {
	db *database.Database
}

func NewAuthHandler(db *database.Database) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request){
	var user models.user
	
}