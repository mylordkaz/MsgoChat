package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mylordkaz/MsgoChat/services/auth-service/internal/models"
	"github.com/mylordkaz/MsgoChat/services/auth-service/pkg/database"
	"github.com/mylordkaz/MsgoChat/services/auth-service/pkg/hash"
)

type AuthHandler struct {
	db *database.Database
}

func NewAuthHandler(db *database.Database) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request){
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return 
	}

	hashedPassword, err := hash.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// save user to db
	if err := h.db.CreateUser(&user); err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}