package service

import (
	"context"
	"user-service/internal/models"
	"user-service/internal/proto"
	"user-service/internal/repository"
)

type userService struct {
	repo repository.UserRepository
	proto.UnimplementedUserServiceServer
}

func NewUserService(repo repository.UserRepository) proto.UserServiceServer {
	return &userService{repo: repo}
}

func (u userService) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	user := models.User{
		Username:    req.Username,
		Email:       req.Username,
		Avatar:      req.Avatar,
		Description: req.Description,
	}

	err := u.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return &proto.CreateUserResponse{
		Message: "User created successfully",
		User: &proto.User{
			Id:          user.ID,
			Username:    user.Username,
			Email:       user.Email,
			Description: user.Description,
			Avatar:      user.Avatar,
		},
	}, nil
}

func (u userService) GetUserById(ctx context.Context, req *proto.GetUserByIdRequest) (*proto.GetUserByIdResponse, error) {
	user, err := u.repo.GetUserById(req.UserId)
	if err != nil {
		return nil, err
	}

	return &proto.GetUserByIdResponse{User: &proto.User{
		Id:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Description: user.Description,
		Avatar:      user.Avatar,
	}}, nil
}

func (u userService) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	user, err := u.repo.GetUserById(req.UserId)
	if err != nil {
		return nil, err
	}
	user.Username = req.Username
	user.Email = req.Email
	user.Avatar = req.Avatar
	user.Description = req.Description

	err = u.repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return &proto.UpdateUserResponse{Message: "User updated successfully"}, nil
}

func (u userService) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	err := u.repo.DeleteUserById(req.UserId)
	if err != nil {
		return nil, err
	}
	return &proto.DeleteUserResponse{
		Message: "User deleted successfully",
	}, nil
}
