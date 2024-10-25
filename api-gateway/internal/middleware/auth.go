package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// Ваш секретный ключ для подписи JWT
var jwtSecret = []byte(os.Getenv("JWT_KEY"))

// AuthMiddleware - middleware для аутентификации на основе JWT из куки
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// Извлечение токена из куки
		cookie, err := r.Cookie("jwt_token")
		if err != nil {
			http.Error(rw, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenStr := cookie.Value

		// Проверка и декодирование JWT токена
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// Проверка метода подписи токена
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(rw, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Извлечение полезной нагрузки и передача данных пользователя в контекст запроса
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Проверка времени истечения токена
			if exp, ok := claims["exp"].(float64); ok && time.Unix(int64(exp), 0).Before(time.Now()) {
				http.Error(rw, "Token expired", http.StatusUnauthorized)
				return
			}

			// Извлечение идентификатора пользователя из claims
			userID := claims["sub"].(string) // Используем "sub" как идентификатор пользователя
			ctx := context.WithValue(r.Context(), "user_id", userID)
			next.ServeHTTP(rw, r.WithContext(ctx))
		} else {
			http.Error(rw, "Unauthorized", http.StatusUnauthorized)
			return
		}
	})
}
