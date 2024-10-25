package utils

import (
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
)

func ExtractUserIdFromToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie("jwt_token")
	if err != nil {
		return "", err
	}

	// Декодируем JWT токен
	tokenString := cookie.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем алгоритм
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		// Возвращаем ключ для проверки подписи
		return []byte(os.Getenv("JWT_KEY")), nil // замените на ваш секретный ключ
	})

	if err != nil {
		return "", err
	}

	// Извлекаем user_id из токена
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userID, ok := claims["sub"].(string); ok {
			return userID, nil
		}
	}

	return "", nil
}
