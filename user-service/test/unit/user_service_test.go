package unit

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"user-service/internal/models"
	"user-service/internal/proto"
	"user-service/internal/service"
	"user-service/test/mocks"
)

type TestUserService struct {
	mockRepo *mocks.MockUserRepository
}

func (u *TestUserService) TestCreateUser(t *testing.T) {
	userService := service.NewUserService(u.mockRepo)

	req := &proto.CreateUserRequest{
		Username:    "test",
		Email:       "test@test.com",
		Avatar:      "avatar.jpg",
		Description: "test",
	}

	u.mockRepo.On("CreateUser", mock.Anything).Return(nil)

	resp, err := userService.CreateUser(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "User created successfully", resp.Message)
	assert.Equal(t, "test", resp.User.Username)

	u.mockRepo.AssertExpectations(t)
}

func (u *TestUserService) TestGetUserById(t *testing.T) {
	userService := service.NewUserService(u.mockRepo)

	user := models.User{
		ID:          1,
		Username:    "test",
		Email:       "test@test.com",
		Avatar:      "avatar.jpg",
		Description: "test",
	}

	u.mockRepo.On("GetUserById", uint64(user.ID)).Return(user, nil)

	req := &proto.GetUserByIdRequest{UserId: 1}
	resp, err := userService.GetUserById(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "test", resp.User.Username)

	u.mockRepo.AssertExpectations(t)
}

func (u *TestUserService) TestUpdateUser(t *testing.T) {
	userService := service.NewUserService(u.mockRepo)

	user := models.User{
		ID:          1,
		Username:    "test",
		Email:       "test@test.com",
		Avatar:      "avatar.jpg",
		Description: "test",
	}

	u.mockRepo.On("GetUserById", uint64(user.ID)).Return(user, nil)
	u.mockRepo.On("UpdateUser", user).Return(nil)

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

	u.mockRepo.AssertExpectations(t)
}

func (u *TestUserService) TestDeleteUser(t *testing.T) {
	userService := service.NewUserService(u.mockRepo)

	u.mockRepo.On("DeleteUserById", uint64(1)).Return(nil)

	req := &proto.DeleteUserRequest{UserId: 1}
	resp, err := userService.DeleteUser(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "User deleted successfully", resp.Message)

	u.mockRepo.AssertExpectations(t)
}
