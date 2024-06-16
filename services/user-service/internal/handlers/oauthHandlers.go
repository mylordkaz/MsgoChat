package handlers

import (
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/mylordkaz/MsgoChat/services/user-service/internal/models"
)

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
	err = db.DB.QueryRow(query, gothUser.UserID).Scan(&user.ID, &user.Email, &user.Name, &user.Provider)
}