package repository

import (
	"gorm.io/gorm"
	"user-service/internal/models"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (u *userRepo) CreateUser(user *models.User) error {
	return u.db.Create(user).Error
}

func (u *userRepo) GetUserById(id uint64) (*models.User, error) {
	var user models.User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) UpdateUser(user *models.User) error {
	return u.db.Save(user).Error
}

func (u *userRepo) DeleteUserById(id uint64) error {
	return u.db.Delete(&models.User{}, id).Error
}
