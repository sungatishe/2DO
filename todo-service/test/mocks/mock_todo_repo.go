package mocks

import (
	"github.com/stretchr/testify/mock"
	"todo-service/internal/models"
)

type MockTodoRepository struct {
	mock.Mock
}

func (m *MockTodoRepository) CreateTodo(todo *models.Todo) error {
	args := m.Called(0)
	return args.Error(0)
}

func (m *MockTodoRepository) GetTodoById(id uint) (*models.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Todo), args.Error(1)
}

func (m *MockTodoRepository) UpdateTodo(todo *models.Todo) error {
	args := m.Called(todo)
	return args.Error(0)
}

func (m *MockTodoRepository) DeleteTodoById(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTodoRepository) ListTodoByUserId(id uint) ([]models.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockTodoRepository) GetTodosByDeadline(deadline string) ([]models.Todo, error) {
	//TODO implement me
	panic("implement me")
}
