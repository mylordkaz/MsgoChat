package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/mylordkaz/MsgoChat/services/auth-service/pkg/utils/jwt"
)


func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return 
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			http.Error(w, "Invlid token format", http.StatusUnauthorized)
			return
		}

		claims, err := jwt.VerifyToken(bearerToken[1])
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// add user ID to the request context
		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
		
	}
}