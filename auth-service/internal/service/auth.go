package service

import (
	"auth-service/internal/models"
	"auth-service/internal/proto"
	"auth-service/internal/rabbitmq/events"
	"auth-service/internal/repository"
	"auth-service/pgk/utils"
	"context"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthService struct {
	Repo    repository.AuthRepository
	channel *amqp.Channel
	proto.UnimplementedAuthServiceServer
}

func NewAuthService(repo repository.AuthRepository, channel *amqp.Channel) proto.AuthServiceServer {
	authService := &AuthService{Repo: repo, channel: channel}
	return authService
}

func (a *AuthService) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	hashPassword, err := utils.GenerateHashPassword(req.Password)
	if err != nil {
		return nil, errors.New("error with create password hash")
	}

	user := &models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hashPassword,
	}

	existingUser, err := a.Repo.GetUserByEmail(user.Email)
	if err != nil && !errors.Is(repository.ErrNotFound, err) {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("Email already taken")
	}

	existingUser, err = a.Repo.GetUserByUsername(user.Username)
	if err != nil && !errors.Is(repository.ErrNotFound, err) {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("Username already taken")
	}

	err = a.Repo.RegisterUser(user)
	if err != nil {
		return nil, err
	}

	event := events.UserRegisteredEvent{
		UserID:   uint(user.ID),
		Username: user.Username,
		Email:    user.Email,
	}

	err = events.PublishUserRegisteredEvent(a.channel, event)
	if err != nil {
		log.Printf("Failed to publish user registred event: %s", err)
		return nil, err
	}

	return &proto.RegisterResponse{Message: "User registered successfully"}, nil

}

func (a *AuthService) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	user, err := a.Repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateJWT(uint(user.ID))
	if err != nil {
		return nil, err
	}

	return &proto.LoginResponse{Token: token}, nil
}

func (a *AuthService) ValidateToken(ctx context.Context, req *proto.ValidateTokenRequest) (*proto.ValidateTokenResponse, error) {
	userId, err := utils.ValidateJWT(req.Token)
	if err != nil {
		return &proto.ValidateTokenResponse{IsValid: false}, nil
	}

	return &proto.ValidateTokenResponse{IsValid: true, UserId: userId}, nil
}
