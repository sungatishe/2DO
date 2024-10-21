package repository

import "auth-service/internal/models"

type AuthRepository interface {
	RegisterUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	CreateSession(session *models.Session) error
	GetSessionByToken(token string) (*models.Session, error)
	DeleteAllSessionsByUserId(id uint) error
	ClearExpiredSessions() error
}
