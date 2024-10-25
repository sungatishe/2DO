package repository

import "todo-service/internal/models"

type TodoRepository interface {
	CreateTodo(todo *models.Todo) error
	GetTodoById(id uint) (*models.Todo, error)
	UpdateTodo(todo *models.Todo) error
	DeleteTodoById(id uint) error
	ListTodoByUserId(id uint) ([]models.Todo, error)
}
