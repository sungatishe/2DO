package middleware

import (
	"api-gateway/internal/client"
	"api-gateway/internal/proto"
	"context"
	"net/http"
	"strings"
	"time"
)

func AuthMiddleware(authClient client.AuthClient) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header required", http.StatusUnauthorized)
				return
			}

			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
				return
			}

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			res, err := authClient.Client.ValidateToken(ctx, &proto.ValidateTokenRequest{Token: tokenParts[1]})
			if err != nil || !res.IsValid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			ctx = context.WithValue(ctx, "userId", res.UserId)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
