package unit

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
	"time"
	"todo-service/internal/models"
	"todo-service/internal/proto/proto"
	"todo-service/internal/service"
	"todo-service/test/mocks"
)

func TestCreateTodo(t *testing.T) {
	mockRepo := new(mocks.MockTodoRepository)
	todoService := service.NewTodoService(mockRepo)

	req := &proto.CreateTodoRequest{
		UserId:      1,
		Title:       "title",
		Description: "test",
		IsDone:      false,
		Deadline:    time.Now().Format(time.RFC3339),
	}

	mockRepo.On("CreateTodo", mock.Anything).Return(nil)

	resp, err := todoService.CreateTodo(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "Todo created successfully", resp.Message)

	mockRepo.AssertExpectations(t)
}

func TestGetTodoById(t *testing.T) {
	mockRepo := new(mocks.MockTodoRepository)
	todoService := service.NewTodoService(mockRepo)

	todo := &models.Todo{
		Model:       gorm.Model{ID: 1},
		UserID:      1,
		Title:       "title",
		Description: "test",
		IsDone:      false,
		Deadline:    time.Now(),
	}

	mockRepo.On("GetTodoById", todo.ID).Return(todo, nil)

	req := &proto.GetTodosByIdRequest{Id: 1}
	resp, err := todoService.GetTodoById(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "title", resp.Todo.Title)

	mockRepo.AssertExpectations(t)
}

func TestUpdateTodo(t *testing.T) {
	mockRepo := new(mocks.MockTodoRepository)
	todoService := service.NewTodoService(mockRepo)

	todo := &models.Todo{
		Model:       gorm.Model{ID: 1},
		UserID:      1,
		Title:       "title",
		Description: "test",
		IsDone:      false,
		Deadline:    time.Now(),
	}

	mockRepo.On("GetTodoById", todo.ID).Return(todo, nil)
	mockRepo.On("UpdateTodo", todo).Return(nil)

	req := &proto.UpdateTodoRequest{
		Id:          1,
		UserId:      1,
		Title:       "updatetitle",
		Description: "updatetest",
		IsDone:      false,
		Deadline:    time.Now().Format(time.RFC3339),
	}

	resp, err := todoService.UpdateTodo(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, req.Title, resp.Todo.Title)

	mockRepo.AssertExpectations(t)
}

func TestDeleteTodo(t *testing.T) {
	mockRepo := new(mocks.MockTodoRepository)
	todoService := service.NewTodoService(mockRepo)

	mockRepo.On("DeleteTodoById", uint(1)).Return(nil)

	req := &proto.DeleteTodoRequest{Id: 1}

	resp, err := todoService.DeleteTodo(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "Todo deleted successfully", resp.Message)

	mockRepo.AssertExpectations(t)
}

func TestListTodo(t *testing.T) {
	mockRepo := new(mocks.MockTodoRepository)
	todoService := service.NewTodoService(mockRepo)

	todos := []models.Todo{
		{
			Model:       gorm.Model{ID: 1},
			UserID:      1,
			Title:       "title 1",
			Description: "test 1",
			IsDone:      false,
			Deadline:    time.Now(),
		},
		{
			Model:       gorm.Model{ID: 2},
			UserID:      1,
			Title:       "title 2",
			Description: "test 2",
			IsDone:      false,
			Deadline:    time.Now(),
		},
	}

	mockRepo.On("ListTodoByUserId", uint(1)).Return(todos, nil)

	req := &proto.ListTodoRequest{UserId: uint64(1)}

	resp, err := todoService.ListTodo(context.Background(), req)

	assert.NoError(t, err)
	assert.Len(t, resp.Todos, 2)
	assert.Equal(t, "title 1", resp.Todos[0].Title)
	assert.Equal(t, "title 2", resp.Todos[1].Title)
	assert.Equal(t, "test 1", resp.Todos[0].Description)
	assert.Equal(t, "test 2", resp.Todos[1].Description)

	mockRepo.AssertExpectations(t)
}

func TestGetTodosByDeadline(t *testing.T) {
	mockRepo := new(mocks.MockTodoRepository)
	todoService := service.NewTodoService(mockRepo)

	deadline := time.Now().Add(48 * time.Hour).Format(time.RFC3339)

	todos := []models.Todo{
		{
			Model:       gorm.Model{ID: 1},
			UserID:      1,
			Title:       "title 1",
			Description: "test 1",
			IsDone:      false,
			Deadline:    time.Now().Add(24 * time.Hour),
		},
		{
			Model:       gorm.Model{ID: 2},
			UserID:      1,
			Title:       "title 2",
			Description: "test 2",
			IsDone:      false,
			Deadline:    time.Now().Add(36 * time.Hour),
		},
	}

	mockRepo.On("GetTodosByDeadline", deadline).Return(todos, nil)

	req := &proto.GetTodosByDeadlineRequest{Deadline: deadline}

	resp, err := todoService.GetTodosByDeadline(context.Background(), req)

	assert.NoError(t, err)
	assert.Len(t, resp.Todos, 2)
	assert.Equal(t, "title 1", resp.Todos[0].Title)
	assert.Equal(t, "title 2", resp.Todos[1].Title)
	assert.Equal(t, "test 1", resp.Todos[0].Description)
	assert.Equal(t, "test 2", resp.Todos[1].Description)

	mockRepo.AssertExpectations(t)

}
