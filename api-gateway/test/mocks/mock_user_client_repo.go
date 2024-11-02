package mocks

import (
	"api-gateway/internal/client"
	"api-gateway/internal/proto/user"
	"context"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type MockUserClient struct {
	mock.Mock
	client.UserClient
}

func (m *MockUserClient) CreateUser(ctx context.Context, req *user.CreateUserRequest, opts ...grpc.CallOption) (*user.CreateUserResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*user.CreateUserResponse), args.Error(1)
}

func (m *MockUserClient) GetUserById(ctx context.Context, req *user.GetUserByIdRequest, opts ...grpc.CallOption) (*user.GetUserByIdResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*user.GetUserByIdResponse), args.Error(1)
}

func (m *MockUserClient) UpdateUser(ctx context.Context, req *user.UpdateUserRequest, opts ...grpc.CallOption) (*user.UpdateUserResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*user.UpdateUserResponse), args.Error(1)
}

func (m *MockUserClient) DeleteUser(ctx context.Context, req *user.DeleteUserRequest, opts ...grpc.CallOption) (*user.DeleteUserResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*user.DeleteUserResponse), args.Error(1)
}
