package repository

import (
	"todo-service/internal/models"
)

type TodoRepository interface {
	CreateTodo(todo *models.Todo) error
	GetTodoById(id uint64) (*models.Todo, error)
	UpdateTodo(todo *models.Todo) error
	DeleteTodoById(id uint64) error
	ListTodoByUserId(id uint64) ([]models.Todo, error)
	GetTodosByDeadline(deadline string) ([]models.Todo, error)
}
