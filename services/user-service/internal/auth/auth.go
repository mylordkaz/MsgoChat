package auth

import (
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	// "github.com/markbates/goth/providers/github"
)

const (
	key = "randomKey"
	MaxAge = 86400 * 30
	IsProd = false  // set to true when https 
)

func NewAuth(router *mux.Router){
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	googleClientID := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	// githubClientID := os.Getenv("GITHUB_CLIENT_ID")
	// githubClientSecret := os.Getenv("GITHUB_CLIENT_SECRET")

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd

	gothic.Store = store

	goth.UseProviders(
		google.New(googleClientID, googleClientSecret, "http://localhost:3000/auth/google/callback"),
		//github.New(githubClientID, githubClientSecret, "http://localhost:3000/auth/github/callback"),
	)
}