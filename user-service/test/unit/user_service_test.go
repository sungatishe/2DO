package unit

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"user-service/internal/models"
	"user-service/internal/proto"
	"user-service/internal/service"
	"user-service/test/mocks"
)

func TestCreateUser(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userService := service.NewUserService(mockRepo)

	req := &proto.CreateUserRequest{
		Username:    "test",
		Email:       "test@test.com",
		Avatar:      "avatar.jpg",
		Description: "test",
	}

	mockRepo.On("CreateUser", mock.Anything).Return(nil)

	resp, err := userService.CreateUser(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "User created successfully", resp.Message)
	assert.Equal(t, "test", resp.User.Username)

	mockRepo.AssertExpectations(t)
}

func TestGetUserById(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userService := service.NewUserService(mockRepo)

	user := models.User{
		Model:       gorm.Model{ID: 1},
		Username:    "test",
		Email:       "test@test.com",
		Avatar:      "avatar.jpg",
		Description: "test",
	}

	mockRepo.On("GetUserById", uint64(user.ID)).Return(&user, nil)

	req := &proto.GetUserByIdRequest{UserId: 1}
	resp, err := userService.GetUserById(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "test", resp.User.Username)

	mockRepo.AssertExpectations(t)
}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userService := service.NewUserService(mockRepo)

	user := models.User{
		Model:       gorm.Model{ID: 1},
		Username:    "test",
		Email:       "test@test.com",
		Avatar:      "avatar.jpg",
		Description: "test",
	}

	mockRepo.On("GetUserById", uint64(user.ID)).Return(&user, nil)
	mockRepo.On("UpdateUser", &user).Return(nil)

	req := &proto.UpdateUserRequest{
		UserId:      1,
		Username:    "updateUser",
		Email:       "updateduser@example.com",
		Avatar:      "new_avatar.png",
		Description: "Updated User",
	}

	resp, err := userService.UpdateUser(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "User updated successfully", resp.Message)

	mockRepo.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userService := service.NewUserService(mockRepo)

	mockRepo.On("DeleteUserById", uint64(1)).Return(nil)

	req := &proto.DeleteUserRequest{UserId: 1}
	resp, err := userService.DeleteUser(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "User deleted successfully", resp.Message)

	mockRepo.AssertExpectations(t)
}
