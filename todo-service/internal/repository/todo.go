package repository

import (
	"database/sql"
	"fmt"
	"time"
	"todo-service/internal/models"
)

type todoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) TodoRepository {
	return &todoRepository{db: db}
}

func (t *todoRepository) CreateTodo(todo *models.Todo) error {
	query := "INSERT INTO todos (user_id, title, description, is_done, deadline) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := t.db.QueryRow(query, todo.UserID, todo.Title, todo.Description, todo.IsDone, todo.Deadline).Scan(&todo.ID)
	if err != nil {
		return fmt.Errorf("failed to create todo: %w", err)
	}

	return nil
}

func (t *todoRepository) GetTodoById(id uint64) (*models.Todo, error) {
	query := "SELECT id, user_id, title, description, is_done, deadline FROM todos WHERE id = $1"
	row := t.db.QueryRow(query, id)

	todo := &models.Todo{}
	err := row.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.IsDone, &todo.Deadline)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to retrieve todo: %w", err)
	}

	return todo, nil
}

func (t *todoRepository) UpdateTodo(todo *models.Todo) error {
	query := "UPDATE todos SET title = $1, description = $2, is_done = $3, deadline = $4 WHERE id = $5"
	_, err := t.db.Exec(query, todo.Title, todo.Description, todo.IsDone, todo.Deadline, todo.ID)
	if err != nil {
		return fmt.Errorf("failed to update todo: %w", err)
	}

	return nil
}

func (t *todoRepository) DeleteTodoById(id uint64) error {
	query := "DELETE FROM todos WHERE id = $1"
	_, err := t.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}

	return nil
}

func (t *todoRepository) ListTodoByUserId(id uint64) ([]models.Todo, error) {
	query := "SELECT id, user_id, title, description, is_done, deadline FROM todos WHERE user_id = $1"
	rows, err := t.db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to list todos: %w", err)
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		todo := models.Todo{}
		err := rows.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.IsDone, &todo.Deadline)
		if err != nil {
			return nil, fmt.Errorf("failed to scan todo: %w", err)
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (t *todoRepository) GetTodosByDeadline(deadline string) ([]models.Todo, error) {
	query := "SELECT id, user_id, title, description, is_done, deadline FROM todos WHERE deadline <= $1"
	deadlineTime, _ := time.Parse(time.RFC3339, deadline)
	deadlineTime = deadlineTime.UTC()
	rows, err := t.db.Query(query, deadlineTime)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve todos by deadline: %w", err)
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		todo := models.Todo{}
		err := rows.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.IsDone, &todo.Deadline)
		if err != nil {
			return nil, fmt.Errorf("failed to scan todo: %w", err)
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
