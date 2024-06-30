package handlers

import (
	"net/http"

	"github.com/mylordkaz/MsgoChat/services/user-service/pkg/jwt"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request){
	cookie, err := r.Cookie("Token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	token := cookie.Value

	claims, err := jwt.VerifyJWT(token)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Welcome to the Dashboard, " + claims.Email + "!"))
}