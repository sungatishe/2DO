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
