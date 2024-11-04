package integration

import (
	"api-gateway/internal/client"
	"api-gateway/internal/handlers"
	"api-gateway/internal/proto/todo"
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

func setupTodoTestRouter() (*chi.Mux, *mocks.MockTodoClient) {
	mockClient := &mocks.MockTodoClient{}
	todoHandlers := handlers.NewTodoHandlers(&client.TodoClient{Client: mockClient})
	router := chi.NewRouter()
	router.Post("/todo", todoHandlers.CreateTodo)
	router.Get("/todo/{id}", todoHandlers.GetTodoById)
	router.Get("/todo", todoHandlers.GetUsersListTodo)
	router.Put("/todo", todoHandlers.UpdateTodo)
	router.Delete("/todo/{id}", todoHandlers.DeleteTodo)
	return router, mockClient
}

func TestCreateTodo(t *testing.T) {
	router, mockClient := setupTodoTestRouter()
	utils.ExtractUserIdFromToken = func(r *http.Request) (string, error) {
		return "1", nil
	}

	mockClient.On("CreateTodo", mock.Anything, mock.AnythingOfType("*todo.CreateTodoRequest")).
		Return(&todo.CreateTodoResponse{
			Message: "Todo created successfully",
			Todo:    &todo.Todo{Id: 1, UserId: 1, Title: "Test Todo"},
		}, nil)

	reqBody, _ := json.Marshal(map[string]string{
		"Title": "Test Todo",
	})

	req, _ := http.NewRequest("POST", "/todo", bytes.NewReader(reqBody))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusCreated, rr.Code)
	var res todo.CreateTodoResponse
	err := json.NewDecoder(rr.Body).Decode(&res)
	require.NoError(t, err)
	require.NotNil(t, res.Todo)
	require.Equal(t, uint64(1), res.Todo.Id)
	require.Equal(t, "Test Todo", res.Todo.Title)

	mockClient.AssertExpectations(t)
}

func TestGetTodo(t *testing.T) {
	router, mockClient := setupTodoTestRouter()

	mockClient.On("GetTodoById", mock.Anything, &todo.GetTodosByIdRequest{Id: 1}).
		Return(&todo.GetTodosByIdResponse{Todo: &todo.Todo{
			Id:          1,
			UserId:      1,
			Title:       "Test Todo",
			Description: "test",
			IsDone:      false,
			Deadline:    "",
		}}, nil)

	req, _ := http.NewRequest("GET", "/todo/1", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	var res todo.GetTodosByIdResponse
	err := json.NewDecoder(rr.Body).Decode(&res)
	require.NoError(t, err)
	require.NotNil(t, res.Todo)
	require.Equal(t, uint64(1), res.Todo.Id)
	require.Equal(t, "Test Todo", res.Todo.Title)

	mockClient.AssertExpectations(t)
}

func TestUpdateTodo(t *testing.T) {
	router, mockClient := setupTodoTestRouter()
	utils.ExtractUserIdFromToken = func(r *http.Request) (string, error) {
		return "1", nil
	}

	mockClient.On("UpdateTodo", mock.Anything, &todo.UpdateTodoRequest{
		Id:          1,
		UserId:      1,
		Title:       "updateTest",
		Description: "updateTest",
		IsDone:      false,
		Deadline:    "",
	}).Return(&todo.UpdateTodoResponse{Todo: &todo.Todo{
		Id:          1,
		UserId:      1,
		Title:       "updateTest",
		Description: "updateTest",
		IsDone:      false,
		Deadline:    "",
	}}, nil)

	reqBody, _ := json.Marshal(map[string]interface{}{
		"Id":          1,
		"Title":       "updateTest",
		"Description": "updateTest",
		"IsDone":      false,
		"Deadline":    "",
	})

	req, _ := http.NewRequest("PUT", "/todo", bytes.NewReader(reqBody))
	req.Header.Set("Authorization", "Bearer testtoken")
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	var res todo.UpdateTodoResponse
	err := json.NewDecoder(rr.Body).Decode(&res)
	require.NoError(t, err)
	require.NotNil(t, res.Todo)
	require.Equal(t, uint64(1), res.Todo.Id)
	require.Equal(t, "updateTest", res.Todo.Title)
	require.Equal(t, "updateTest", res.Todo.Description)

	mockClient.AssertExpectations(t)
}

func TestDeleteTodo(t *testing.T) {
	router, mockClient := setupTodoTestRouter()

	mockClient.On("DeleteTodo", mock.Anything, &todo.DeleteTodoRequest{Id: 1}).
		Return(&todo.DeleteTodoResponse{Message: "Todo deleted successfully"}, nil)

	req, _ := http.NewRequest("DELETE", "/todo/1", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	var res todo.DeleteTodoResponse
	err := json.NewDecoder(rr.Body).Decode(&res)
	require.NoError(t, err)
	require.Equal(t, "Todo deleted successfully", res.Message)
}

func TestGetUsersListTodo(t *testing.T) {
	router, mockClient := setupTodoTestRouter()

	utils.ExtractUserIdFromToken = func(r *http.Request) (string, error) {
		return "1", nil
	}

	mockClient.On("ListTodo", mock.Anything, &todo.ListTodoRequest{UserId: 1}).
		Return(&todo.ListTodoResponse{Todos: []*todo.Todo{
			{
				Id:          1,
				UserId:      1,
				Title:       "First Todo",
				Description: "First description",
				IsDone:      false,
				Deadline:    "2024-11-30",
			},
			{
				Id:          2,
				UserId:      1,
				Title:       "Second Todo",
				Description: "Second description",
				IsDone:      true,
				Deadline:    "2024-12-15",
			},
		}}, nil)

	req, _ := http.NewRequest("GET", "/todo", nil)
	req.Header.Set("Authorization", "Bearer testtoken")
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	require.Equal(t, http.StatusOK, rr.Code)
	var res todo.ListTodoResponse
	err := json.NewDecoder(rr.Body).Decode(&res)
	require.NoError(t, err)
	require.Len(t, res.Todos, 2)

	require.Equal(t, uint64(1), res.Todos[0].Id)
	require.Equal(t, "First Todo", res.Todos[0].Title)
	require.Equal(t, "First description", res.Todos[0].Description)
	require.False(t, res.Todos[0].IsDone)
	require.Equal(t, "2024-11-30", res.Todos[0].Deadline)

	require.Equal(t, uint64(2), res.Todos[1].Id)
	require.Equal(t, "Second Todo", res.Todos[1].Title)
	require.Equal(t, "Second description", res.Todos[1].Description)
	require.True(t, res.Todos[1].IsDone)
	require.Equal(t, "2024-12-15", res.Todos[1].Deadline)

	mockClient.AssertExpectations(t)
}
