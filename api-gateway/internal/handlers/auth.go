package handlers

import (
	"api-gateway/internal/client"
	"api-gateway/internal/proto"
	"context"
	"encoding/json"
	"net/http"
)

type AuthHandlers struct {
	authClient *client.AuthClient
}

func NewAuthHandlers(authClient *client.AuthClient) *AuthHandlers {
	return &AuthHandlers{authClient: authClient}
}

func (a *AuthHandlers) Register(rw http.ResponseWriter, r *http.Request) {
	var req proto.RegisterRequest
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
	var req proto.LoginRequest
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

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(res); err != nil {
		http.Error(rw, "Failed to encode response", http.StatusInternalServerError)
	}
}
