package repository

import (
	"gorm.io/gorm"
	"todo-service/internal/models"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (t *todoRepository) CreateTodo(todo *models.Todo) error {
	return t.db.Create(todo).Error
}

func (t *todoRepository) GetTodoById(id uint) (*models.Todo, error) {
	var todo models.Todo
	err := t.db.First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (t *todoRepository) UpdateTodo(todo *models.Todo) error {
	return t.db.Save(todo).Error
}

func (t *todoRepository) DeleteTodoById(id uint) error {
	return t.db.Delete(&models.Todo{}, id).Error
}

func (t *todoRepository) ListTodoByUserId(id uint) ([]models.Todo, error) {
	var todos []models.Todo
	err := t.db.Where("user_id = ?", id).Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (t *todoRepository) GetTodosByDeadline(deadline string) ([]models.Todo, error) {
	var todos []models.Todo
	err := t.db.Where("deadline = ?", deadline).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}
