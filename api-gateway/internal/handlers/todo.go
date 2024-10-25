package handlers

import (
	"api-gateway/internal/client"
	"api-gateway/internal/proto/todo"
	"api-gateway/pgk/utils"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

type TodoHandlers struct {
	client *client.TodoClient
}

func NewTodoHandlers(client *client.TodoClient) *TodoHandlers {
	return &TodoHandlers{client: client}
}

func (h *TodoHandlers) CreateTodo(rw http.ResponseWriter, r *http.Request) {
	userID, err := utils.ExtractUserIdFromToken(r)
	if err != nil || userID == "" {
		http.Error(rw, "Unauthorized", http.StatusUnauthorized)
		return
	}
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
		return
	}

	var req todo.CreateTodoRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(rw, "Invalid request body", http.StatusBadRequest)
		log.Println(err)
		return
	}

	req.UserId = id

	res, err := h.client.Client.CreateTodo(context.Background(), &req)
	if err != nil {
		http.Error(rw, "Error in creating todo", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(rw).Encode(res); err != nil {
		http.Error(rw, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *TodoHandlers) GetTodoById(rw http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "id")

	id, err := strconv.ParseUint(todoId, 10, 64)
	if err != nil {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
		return
	}

	req := &todo.GetTodosByIdRequest{Id: id}

	res, err := h.client.Client.GetTodoById(context.Background(), req)
	if err != nil {
		http.Error(rw, "Error in getting todo", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	err = json.NewEncoder(rw).Encode(res)
	if err != nil {
		http.Error(rw, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *TodoHandlers) UpdateTodo(rw http.ResponseWriter, r *http.Request) {
	userID, err := utils.ExtractUserIdFromToken(r)
	if err != nil || userID == "" {
		http.Error(rw, "Unauthorized", http.StatusUnauthorized)
		return
	}
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
		return
	}

	var req todo.UpdateTodoRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(rw, "Invalid request", http.StatusBadRequest)
		return
	}

	req.UserId = id

	res, err := h.client.Client.UpdateTodo(context.Background(), &req)
	if err != nil {
		http.Error(rw, "Error in updating todo", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	err = json.NewEncoder(rw).Encode(res)
	if err != nil {
		http.Error(rw, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *TodoHandlers) DeleteTodo(rw http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(todoId, 10, 64)
	if err != nil {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
		return
	}

	req := &todo.DeleteTodoRequest{Id: id}

	res, err := h.client.Client.DeleteTodo(context.Background(), req)
	if err != nil {
		http.Error(rw, "Error in deleting todo", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	err = json.NewEncoder(rw).Encode(res)
	if err != nil {
		http.Error(rw, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *TodoHandlers) GetUsersListTodo(rw http.ResponseWriter, r *http.Request) {
	userID, err := utils.ExtractUserIdFromToken(r)
	if err != nil || userID == "" {
		http.Error(rw, "Unauthorized", http.StatusUnauthorized)
		return
	}
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
		return
	}

	req := &todo.ListTodoRequest{UserId: id}

	res, err := h.client.Client.ListTodo(context.Background(), req)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	err = json.NewEncoder(rw).Encode(res)
	if err != nil {
		http.Error(rw, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
