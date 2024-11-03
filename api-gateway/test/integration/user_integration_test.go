package integration

import (
	"api-gateway/internal/client"
	"api-gateway/internal/handlers"
	"api-gateway/internal/proto/user"
	"api-gateway/pgk/utils"
	"api-gateway/test/mocks"
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestRouter() (*chi.Mux, *mocks.MockUserClient) {
	mockClient := &mocks.MockUserClient{}
	userHandlers := handlers.NewUserHandlers(&client.UserClient{Client: mockClient})
	router := chi.NewRouter()
	router.Post("/user", userHandlers.CreateUser)
	router.Get("/user", userHandlers.GetUserById)
	router.Put("/user", userHandlers.UpdateUser)
	router.Delete("/user/{id}", userHandlers.DeleteUser)
	return router, mockClient
}

func TestCreateUser(t *testing.T) {
	router, mockClient := setupTestRouter()

	mockClient.On("CreateUser", mock.Anything, mock.AnythingOfType("*user.CreateUserRequest")).
		Return(&user.CreateUserResponse{
			User:    &user.User{Id: 1},
			Message: "User created successfully",
		}, nil)

	reqBody, _ := json.Marshal(map[string]string{
		"Username":    "testusername",
		"Email":       "test@test.com",
		"Avatar":      "avatar",
		"Description": "test",
	})
	req, _ := http.NewRequest("POST", "/user", bytes.NewReader(reqBody))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	require.Equal(t, http.StatusCreated, rr.Code)
	var res user.CreateUserResponse
	err := json.NewDecoder(rr.Body).Decode(&res)
	require.NoError(t, err)
	require.NotNil(t, res.User)
	require.Equal(t, uint64(1), res.User.Id)
}

func TestGetUserById(t *testing.T) {
	originalFunc := utils.ExtractUserIdFromToken
	defer func() { utils.ExtractUserIdFromToken = originalFunc }()

	utils.ExtractUserIdFromToken = func(r *http.Request) (string, error) {
		return "1", nil
	}

	router, mockClient := setupTestRouter()

	mockClient.On("GetUserById", mock.Anything, &user.GetUserByIdRequest{UserId: uint64(1)}).
		Return(&user.GetUserByIdResponse{User: &user.User{Username: "test"}}, nil)

	req, _ := http.NewRequest("GET", "/user", nil)
	req.Header.Set("Authorization", "Bearer testtoken")
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	var res user.GetUserByIdResponse
	err := json.NewDecoder(rr.Body).Decode(&res)
	require.NoError(t, err)
	require.Equal(t, "test", res.User.Username)
}

func TestUpdateUser(t *testing.T) {
	router, mockClient := setupTestRouter()

	mockClient.On("UpdateUser", mock.Anything, &user.UpdateUserRequest{
		UserId:      uint64(1),
		Username:    "testupdate",
		Email:       "update@test.com",
		Avatar:      "upd",
		Description: "upd",
	}).Return(&user.UpdateUserResponse{Message: "User updated successfully"}, nil)

	reqBody, _ := json.Marshal(map[string]interface{}{
		"user_id":     1,
		"Username":    "testupdate",
		"Email":       "update@test.com",
		"Avatar":      "upd",
		"Description": "upd",
	})
	req, _ := http.NewRequest("PUT", "/user", bytes.NewReader(reqBody))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	var res user.UpdateUserResponse
	err := json.NewDecoder(rr.Body).Decode(&res)
	require.NoError(t, err)
	require.Equal(t, "User updated successfully", res.Message)
}

func TestDeleteUser(t *testing.T) {
	router, mockClient := setupTestRouter()

	mockClient.On("DeleteUser", mock.Anything, &user.DeleteUserRequest{UserId: uint64(1)}).
		Return(&user.DeleteUserResponse{Message: "User deleted successfully"}, nil)

	req, _ := http.NewRequest("DELETE", "/user/1", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	var res user.UpdateUserResponse
	err := json.NewDecoder(rr.Body).Decode(&res)
	require.NoError(t, err)
	require.Equal(t, "User deleted successfully", res.Message)
}
