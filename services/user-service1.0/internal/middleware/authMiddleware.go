package middleware

import (
	"context"
	"net/http"

	"github.com/mylordkaz/MsgoChat/services/user-service/contextkeys"
	"github.com/mylordkaz/MsgoChat/services/user-service/pkg/jwt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request)  {
		cookie, err := r.Cookie("Token")
		if err != nil || cookie.Value == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		claims, err := jwt.VerifyJWT(cookie.Value)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// store user email in context
		ctx := context.WithValue(r.Context(), contextkeys.UserContextKey, claims.Email)
        next.ServeHTTP(w, r.WithContext(ctx))
	})
}