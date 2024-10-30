package mocks

import (
	"github.com/stretchr/testify/mock"
	"user-service/internal/models"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user *models.User) error {
	args := m.Called(0)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserById(id uint64) (*models.User, error) {
	args := m.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUserById(id uint64) error {
	args := m.Called(id)
	return args.Error(0)
}
