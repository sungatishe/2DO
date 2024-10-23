package repository

import (
	"auth-service/internal/models"
	"errors"
	"gorm.io/gorm"
)

var ErrNotFound = errors.New("not found")

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (a *authRepository) RegisterUser(user *models.User) error {
	if err := a.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (a *authRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := a.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (a *authRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := a.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (a *authRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := a.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}
