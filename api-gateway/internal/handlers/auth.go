package handlers

import (
	"api-gateway/internal/client"
	"api-gateway/internal/proto/auth"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type AuthHandlers struct {
	authClient *client.AuthClient
}

func NewAuthHandlers(authClient *client.AuthClient) *AuthHandlers {
	return &AuthHandlers{authClient: authClient}
}

func (a *AuthHandlers) Register(rw http.ResponseWriter, r *http.Request) {
	var req auth.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(rw, "Invalid request", http.StatusBadRequest)
		return
	}

	res, err := a.authClient.Client.Register(context.Background(), &req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(rw).Encode(res); err != nil {
		http.Error(rw, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (a *AuthHandlers) Login(rw http.ResponseWriter, r *http.Request) {
	var req auth.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(rw, "Invalid request", http.StatusBadRequest)
		return
	}

	res, err := a.authClient.Client.Login(context.Background(), &req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(rw, &http.Cookie{
		Name:     "jwt_token",
		Value:    res.Token,
		HttpOnly: true,                    // Ограничение доступа к cookie только на стороне сервера
		Secure:   true,                    // Только для HTTPS-соединений (на время разработки может быть установлено в false)
		SameSite: http.SameSiteStrictMode, // Предотвращение отправки cookie в кросс-сайтовых запросах
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour), // Установка времени жизни токена (например, 24 часа)
	})

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(res); err != nil {
		http.Error(rw, "Failed to encode response", http.StatusInternalServerError)
	}
}
