package repository

import (
	"user-service/internal/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserById(id uint64) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUserById(id uint64) error
}
