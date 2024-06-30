package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/markbates/goth/gothic"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/models"
	"github.com/mylordkaz/MsgoChat/services/user-service/pkg/hash"
	"github.com/mylordkaz/MsgoChat/services/user-service/pkg/jwt"
)

type RegisterRequest struct{
	Email 		string `json:"email"`
	Password 	string `json:"password"`
	Name 		string `json:"name"`
}
type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request){
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	hashedPassword, err := hash.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Email: req.Email,
		PasswordHash: hashedPassword,
		Name: req.Name,
		Provider: "local",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := `
		INSERT INTO users (email, password_hash, name, provider, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	err = db.QueryRow(query, user.Email, user.PasswordHash, user.Name, user.Provider, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	token, err := jwt.GenerateJWT(user.ID, user.Email)
	if err != nil {
		log.Println("Error generating token:", err)
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}
	
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name: "Token",
		Value: token,
		Expires: expirationTime,
		HttpOnly: true,
		Secure: false, // use true in prod
		Path: "/",
	})

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	// w.Header().Set("content-type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(user)
}

func LoginHandler(w http.ResponseWriter, r *http.Request){
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var user models.User
	query := `SELECT id, email, password_hash, name, provider FROM users WHERE email = $1`
	err := db.QueryRow(query, req.Email).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.Provider)
	if err != nil {
		if err == sql.ErrNoRows{
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		} else {
			http.Error(w, "Error retriving user", http.StatusInternalServerError)
		}
		return 
	}

	if !hash.CheckPasswordHash(req.Password, user.PasswordHash){
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}
	token, err := jwt.GenerateJWT(user.ID, user.Email)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}
	
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name: "Token",
		Value: token,
		Expires: expirationTime,
		HttpOnly: true,
		Secure: false, // true in prod
		Path: "/",
	})

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	// w.Header().Set("content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(user)
}


func AuthHandlers(w http.ResponseWriter, r *http.Request){
	gothic.BeginAuthHandler(w, r)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request){
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil{
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	var user models.User
	query := `SELECT id, email, name, provider FROM users WHERE google_id = $1 OR github_id = $1`
	err = db.QueryRow(query, gothUser.UserID).Scan(&user.ID, &user.Email, &user.Name, &user.Provider)
	if err != nil {
		if err == sql.ErrNoRows {
			user = models.User{
				Email:       gothUser.Email,
                GoogleID:    &gothUser.UserID,
                Name:        gothUser.Name,
                AvatarURL:   &gothUser.AvatarURL,
                Provider:    gothUser.Provider,
                AccessToken: &gothUser.AccessToken,
                RefreshToken: &gothUser.RefreshToken,
                TokenExpiry: &gothUser.ExpiresAt,
                CreatedAt:   time.Now(),
                UpdatedAt:   time.Now(),
			}

			if gothUser.Provider == "google"{
				user.GoogleID = &gothUser.UserID
			} else if gothUser.Provider == "github"{
				user.GithubID = &gothUser.UserID
			}

			query = `INSERT INTO users (email, google_id, github_id, name, avatar_url, provider, access_token, refresh_token, token_expiry, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
			RETURNING id
			`
			err = db.QueryRow(query,user.Email, user.GoogleID, user.GithubID, user.Name, user.AvatarURL, user.Provider, user.AccessToken, user.RefreshToken, user.TokenExpiry, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
			if err != nil {
				http.Error(w, "Error saing user", http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "Error retriving user", http.StatusInternalServerError)
			return
		}
	}

	token, err := jwt.GenerateJWT(user.ID, user.Email)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}
	
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name: "Token",
		Value: token,
		Expires: expirationTime,
		HttpOnly: true,
		Secure: false, // true in prod
		Path: "/",
	})

	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	fmt.Fprintf(w, "User: %v", user)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request){
	gothic.Logout(w, r)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}