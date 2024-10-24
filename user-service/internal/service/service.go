package service

import (
	"context"
	"user-service/internal/proto"
)

type UserService interface {
	CreateUser(context.Context, *proto.CreateUserRequest) (*proto.CreateUserResponse, error)
	GetUserById(context.Context, *proto.GetUserByIdRequest) (*proto.GetUserByIdResponse, error)
	UpdateUser(context.Context, *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error)
	DeleteUser(context.Context, *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error)
}
