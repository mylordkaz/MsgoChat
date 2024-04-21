package auth

import (
    "github.com/markbates/goth"
    "github.com/markbates/goth/providers/google"
    
)

func init() {
    goth.UseProviders(
        google.New("your-client-id", "your-client-secret", "http://localhost:8080/auth/google/callback"),
    )
}