package mocks

import (
	"api-gateway/internal/client"
	"api-gateway/internal/proto/auth"
	"context"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type MockAuthClient struct {
	mock.Mock
	client.AuthClient
}

func (m *MockAuthClient) Register(ctx context.Context, req *auth.RegisterRequest, opts ...grpc.CallOption) (*auth.RegisterResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*auth.RegisterResponse), args.Error(1)
}

func (m *MockAuthClient) Login(ctx context.Context, req *auth.LoginRequest, opts ...grpc.CallOption) (*auth.LoginResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*auth.LoginResponse), args.Error(1)
}

func (m *MockAuthClient) Logout(ctx context.Context, req *auth.LogOutRequest, opts ...grpc.CallOption) (*auth.LogOutResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*auth.LogOutResponse), args.Error(1)
}

func (m *MockAuthClient) ValidateToken(ctx context.Context, req *auth.ValidateTokenRequest, opts ...grpc.CallOption) (*auth.ValidateTokenResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*auth.ValidateTokenResponse), args.Error(1)
}
