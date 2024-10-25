package handlers

import (
	"api-gateway/internal/client"
	"api-gateway/internal/proto/user"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type UserHandlers struct {
	userClient *client.UserClient
}

func NewUserHandlers(userClient *client.UserClient) *UserHandlers {
	return &UserHandlers{userClient: userClient}
}

func (u *UserHandlers) CreateUser(rw http.ResponseWriter, r *http.Request) {
	var req user.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(rw, "Invalid request", http.StatusBadRequest)
		return
	}

	res, err := u.userClient.Client.CreateUser(context.Background(), &req)
	if err != nil {
		http.Error(rw, "Invalid request", http.StatusBadRequest)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(rw).Encode(res); err != nil {
		http.Error(rw, "Failed to encode response", http.StatusInternalServerError)
		return
	}

}

func (u *UserHandlers) GetUserById(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(rw, "Invalid request", http.StatusBadRequest)
		return
	}

	req := &user.GetUserByIdRequest{UserId: userID}

	res, err := u.userClient.Client.GetUserById(context.Background(), req)
	if err != nil {
		http.Error(rw, "Failed to get user by id", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(rw).Encode(res); err != nil {
		http.Error(rw, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (u *UserHandlers) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	var req user.UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(rw, "Invalid request", http.StatusBadRequest)
		return
	}

	res, err := u.userClient.Client.UpdateUser(context.Background(), &req)
	if err != nil {
		http.Error(rw, "Error in updating user", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(rw).Encode(res); err != nil {
		http.Error(rw, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (u *UserHandlers) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(rw, "Invalid request", http.StatusBadRequest)
		return
	}

	req := &user.DeleteUserRequest{UserId: userID}

	res, err := u.userClient.Client.DeleteUser(context.Background(), req)
	if err != nil {
		http.Error(rw, "Error in deleting user", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(rw).Encode(res); err != nil {
		http.Error(rw, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
