package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/markbates/goth/gothic"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/models"
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
	query := `SELECT id, email, name, provider FROM users WHERE google_id = $1`
	err = db.QueryRow(query, gothUser.UserID).Scan(&user.ID, &user.Email, &user.Name, &user.Provider)
	if err != nil {
		if err == sql.ErrNoRows {
			user = models.User{
				Email:       gothUser.Email,
                GoogleID:    &gothUser.UserID,
                Name:        gothUser.Name,
                AvatarURL:   &gothUser.AvatarURL,
                Provider:    "google",
                AccessToken: &gothUser.AccessToken,
                RefreshToken: &gothUser.RefreshToken,
                TokenExpiry: &gothUser.ExpiresAt,
                CreatedAt:   time.Now(),
                UpdatedAt:   time.Now(),
			}
			query = `INSERT INTO users (email, google_id, name, avatar_url, provider, access_token, refresh_token, token_expiry, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
			RETURNING id
			`
			err = db.QueryRow(query,user.Email, user.GoogleID, user.Name, user.AvatarURL, user.Provider, user.AccessToken, user.RefreshToken, user.TokenExpiry, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
			if err != nil {
				http.Error(w, "Error saing user", http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "Error retriving user", http.StatusInternalServerError)
			return
		}
	}
	fmt.Fprintf(w, "User: %v", user)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request){
	gothic.Logout(w, r)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}