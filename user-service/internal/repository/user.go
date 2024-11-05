package repository

import (
	"database/sql"
	"fmt"
	"user-service/internal/models"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (u *userRepo) CreateUser(user *models.User) error {
	query := "INSERT INTO users (username, email, avatar, description) VALUES ($1, $2, $3, $4) RETURNING id"
	err := u.db.QueryRow(query, user.Username, user.Email, user.Avatar, user.Description).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("Failed to create user: %w", err)
	}
	return nil
}

func (u *userRepo) GetUserById(id uint64) (*models.User, error) {
	query := "SELECT id, username, email, avatar, description FROM users WHERE id = $1"
	row := u.db.QueryRow(query, id)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Avatar, &user.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("Failed to retreive user: %w", err)
	}
	return user, nil
}

func (u *userRepo) UpdateUser(user *models.User) error {
	query := "UPDATE users SET username = $1, email = $2, avatar = $3, description = $4 WHERE id = $5"
	_, err := u.db.Exec(query, user.Username, user.Email, user.Avatar, user.Description, user.ID)
	if err != nil {
		return fmt.Errorf("Failed to update user: %w", err)
	}

	return nil
}

func (u *userRepo) DeleteUserById(id uint64) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := u.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Failed to delete user: %w", err)
	}

	return nil
}
