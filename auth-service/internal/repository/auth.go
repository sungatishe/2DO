package repository

import (
	"auth-service/internal/models"
	"database/sql"
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}

func (a *authRepository) RegisterUser(user *models.User) error {
	query := "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id"
	err := a.db.QueryRow(query, user.Username, user.Email, user.PasswordHash).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}

	return nil
}

func (a *authRepository) GetUserByEmail(email string) (*models.User, error) {
	query := "SELECT id, username, email, password_hash FROM users WHERE email = $1"
	row := a.db.QueryRow(query, email)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return user, nil
}

func (a *authRepository) GetUserByUsername(username string) (*models.User, error) {
	query := "SELECT id, username, email, password_hash FROM users WHERE username = $1"
	row := a.db.QueryRow(query, username)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}
	return user, nil
}

func (a *authRepository) GetUserByID(id uint) (*models.User, error) {
	query := "SELECT id, username, email, password_hash FROM users WHERE id = $1"
	row := a.db.QueryRow(query, id)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return user, nil
}
