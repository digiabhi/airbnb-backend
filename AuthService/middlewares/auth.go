package middlewares

import (
	env "AuthService/config/env"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Authorization header must start with 'Bearer '", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			http.Error(w, "Token is required", http.StatusUnauthorized)
			return
		}

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.GetString("JWT_SECRET", "TOKEN")), nil
		})

		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		userId, okId := claims["id"]

		email, okEmail := claims["email"].(string)

		fmt.Println(claims, userId, okId)

		if !okId || !okEmail {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "userId", userId)
		ctx = context.WithValue(ctx, "email", email)
		fmt.Println("Authenticated user ID:", userId, "Email:", email)
		// For demonstration purposes, we assume the token is valid.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
