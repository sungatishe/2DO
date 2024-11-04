package integration

import (
	"api-gateway/internal/client"
	"api-gateway/internal/handlers"
	"api-gateway/internal/proto/auth"
	"api-gateway/test/mocks"
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func setupAuthTodoRouter() (*chi.Mux, *mocks.MockAuthClient) {
	mockClient := &mocks.MockAuthClient{}
	authHandlers := handlers.NewAuthHandlers(&client.AuthClient{Client: mockClient})
	router := chi.NewRouter()
	router.Post("/register", authHandlers.Register)
	router.Post("/login", authHandlers.Login)
	router.Post("/logout", authHandlers.Logout)
	return router, mockClient
}

func TestRegister(t *testing.T) {
	router, mockClient := setupAuthTodoRouter()

	mockClient.On("Register", mock.Anything, mock.AnythingOfType("*auth.RegisterRequest")).
		Return(&auth.RegisterResponse{Message: "User registered successfully"}, nil)

	reqBody, _ := json.Marshal(map[string]string{
		"Username": "testuser",
		"Email":    "test@test.com",
		"Password": "password123",
	})

	req, _ := http.NewRequest("POST", "/register", bytes.NewReader(reqBody))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	var res auth.RegisterResponse
	err := json.NewDecoder(rr.Body).Decode(&res)
	require.NoError(t, err)
	require.Equal(t, "User registered successfully", res.Message)

	mockClient.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	router, mockClient := setupAuthTodoRouter()

	mockClient.On("Login", mock.Anything, mock.AnythingOfType("*auth.LoginRequest")).
		Return(&auth.LoginResponse{Token: "testtoken"}, nil)

	reqBody, _ := json.Marshal(map[string]string{
		"Email":    "test@test.com",
		"Password": "test",
	})
	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(reqBody))
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	var res auth.LoginResponse
	err := json.NewDecoder(rr.Body).Decode(&res)
	require.NoError(t, err)
	require.Equal(t, "testtoken", res.Token)

	cookie := rr.Result().Cookies()
	require.NotNil(t, cookie)
	require.Equal(t, "jwt_token", cookie[0].Name)
	require.Equal(t, "testtoken", cookie[0].Value)
}

func TestLogout(t *testing.T) {
	router, _ := setupAuthTodoRouter()

	req, _ := http.NewRequest("POST", "/logout", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	var res map[string]string
	err := json.NewDecoder(rr.Body).Decode(&res)
	require.NoError(t, err)
	require.Equal(t, "Logout successful", res["message"])

	cookie := rr.Result().Cookies()
	require.NotNil(t, cookie)
	require.Equal(t, "jwt_token", cookie[0].Name)
	require.True(t, cookie[0].Expires.Before(time.Now()))
}
