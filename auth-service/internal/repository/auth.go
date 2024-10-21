package repository

import (
	"auth-service/internal/models"
	"database/sql"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}

func (a authRepository) RegisterUser(user *models.User) error {
	//query := "INSERT INTO users ()"
}

func (a authRepository) GetUserByEmail(email string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a authRepository) GetUserByUsername(username string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a authRepository) GetUserByID(id uint) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (a authRepository) CreateSession(session *models.Session) error {
	//TODO implement me
	panic("implement me")
}

func (a authRepository) GetSessionByToken(token string) (*models.Session, error) {
	//TODO implement me
	panic("implement me")
}

func (a authRepository) DeleteAllSessionsByUserId(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (a authRepository) ClearExpiredSessions() error {
	//TODO implement me
	panic("implement me")
}
