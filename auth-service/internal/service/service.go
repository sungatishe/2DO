package service

import (
	"auth-service/internal/proto"
	"context"
)

type Authorization interface {
	Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error)
	Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error)
	Logout(ctx context.Context, req *proto.LogOutRequest) (*proto.LogOutResponse, error)
	ValidateToken(ctx context.Context, req *proto.ValidateTokenRequest) (*proto.ValidateTokenResponse, error)
}
