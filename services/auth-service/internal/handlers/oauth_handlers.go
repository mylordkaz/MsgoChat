package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"github.com/mylordkaz/MsgoChat/services/auth-service/internal/config"
	"github.com/mylordkaz/MsgoChat/services/auth-service/internal/service"
)

func InitializeOAuth(router *mux.Router, cfg *config.Config, authService *service.AuthService) {
	goth.UseProviders(
		google.New(cfg.GoogleClientID, cfg.GoogleClientSecret, fmt.Sprintf("%s/auth/google/callback", cfg.ServerURL)),
		github.New(cfg.GitHubClientID, cfg.GitHubClientSecret, fmt.Sprintf("%s/auth/github/callback", cfg.ServerURL)),
	)

	router.HandleFunc("/auth/{provider}", func(w http.ResponseWriter, r *http.Request) {
		gothic.BeginAuthHandler(w, r)
	})

	router.HandleFunc("/auth/{provider}/callback", func(w http.ResponseWriter, r *http.Request) {
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		token, err := authService.LoginOrRegisterOAuthUser(user)
		if err != nil {
			http.Error(w, "Failed to authenticate user", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	})
}