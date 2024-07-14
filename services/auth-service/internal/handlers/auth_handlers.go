package handlers

import (
	"encoding/json"
	"net/http"
	

	"github.com/mylordkaz/MsgoChat/services/auth-service/internal/models"
	"github.com/mylordkaz/MsgoChat/services/auth-service/internal/service"
	
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
    return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request){
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return 
	}

	if err := h.authService.RegisterUser(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// generate token for newly registred user
	token, err := h.authService.LoginUser(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Error logging in after registration", http.StatusInternalServerError)
		return
	}

	response := struct {
		Message 	string `json:"message"`
		Token 		string 	`json:"token"`
		TokenType 	string 	`json:"token_type"`
		ExpiresIn	int64	`json:"expires_in"`
	}{
		Message: "User registered successfully",
		Token: token,
		TokenType: "Bearer",
		ExpiresIn: 24 * 60 * 60,
	}

	
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request){
	var loginRequest models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.authService.LoginUser(loginRequest.Username, loginRequest.Password)
    if err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

	response := struct {
		Token 		string 	`json:"token"`
		TokenType 	string 	`json:"token_type"`
		ExpiresIn	int64	`json:"expires_in"`
	}{
		Token: token,
		TokenType: "Bearer",
		ExpiresIn: 24 * 60 * 60,
	}

	w.Header().Set("Content-Type", "application/json")	
	json.NewEncoder(w).Encode(response)
}


// fetch user from database
// user, err := h.db.GetUserByUsername(loginRequest.Username)
// if err != nil {
// 	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
// 	return
// }

// check password
// if !hash.CheckPasswordHash(loginRequest.Password, user.Password){
// 	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
// 	return
// }

// generate JWT
// token, err := jwt.GenerateToken(user.ID, 24*time.Hour)
// if err != nil {
// 	http.Error(w, "Error generating token", http.StatusInternalServerError)
// 	return
// }