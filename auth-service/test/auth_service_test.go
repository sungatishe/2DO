package test

//
//import (
//	"auth-service/internal/models"
//	"auth-service/internal/proto"
//	"auth-service/internal/repository"
//	"auth-service/internal/service"
//	"context"
//	"errors"
//	"github.com/stretchr/testify/mocks"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//)
//
//func TestRegister(t *testing.T) {
//	mockRepo := new(MockAuthRepository)
//	mockChannel := new(MockChannel) // Используем мок-канал
//	authService := service.NewAuthService(mockRepo, mockChannel)
//
//	tests := []struct {
//		name                        string
//		input                       *proto.RegisterRequest
//		expectedMessage             string
//		expectedError               error
//		repoGetUserByEmailReturn    *models.User
//		repoGetUserByEmailError     error
//		repoGetUserByUsernameReturn *models.User
//		repoGetUserByUsernameError  error
//		repoRegisterUserError       error
//	}{
//		{
//			name: "successful registration",
//			input: &proto.RegisterRequest{
//				Username: "testuser",
//				Email:    "test@example.com",
//				Password: "password123",
//			},
//			expectedMessage:             "User registered successfully",
//			expectedError:               nil,
//			repoGetUserByEmailReturn:    nil,
//			repoGetUserByEmailError:     repository.ErrNotFound,
//			repoGetUserByUsernameReturn: nil,
//			repoGetUserByUsernameError:  repository.ErrNotFound,
//			repoRegisterUserError:       nil,
//		},
//		{
//			name: "email already taken",
//			input: &proto.RegisterRequest{
//				Username: "testuser",
//				Email:    "taken@example.com",
//				Password: "password123",
//			},
//			expectedMessage:             "",
//			expectedError:               errors.New("Email already taken"),
//			repoGetUserByEmailReturn:    &models.User{},
//			repoGetUserByEmailError:     nil,
//			repoGetUserByUsernameReturn: nil,
//			repoGetUserByUsernameError:  repository.ErrNotFound,
//			repoRegisterUserError:       nil,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mockRepo.On("GetUserByEmail", tt.input.Email).Return(tt.repoGetUserByEmailReturn, tt.repoGetUserByEmailError)
//			mockRepo.On("GetUserByUsername", tt.input.Username).Return(tt.repoGetUserByUsernameReturn, tt.repoGetUserByUsernameError)
//			mockRepo.On("RegisterUser", mocks.Anything).Return(tt.repoRegisterUserError)
//
//			// Ожидаем, что Publish будет вызван в случае успешной регистрации
//			if tt.expectedError == nil {
//				mockChannel.On("Publish", mocks.Anything, mocks.Anything, false, false, mocks.Anything).Return(nil)
//			}
//
//			resp, err := authService.Register(context.Background(), tt.input)
//
//			if tt.expectedError != nil {
//				assert.Error(t, err)
//				assert.Equal(t, tt.expectedError.Error(), err.Error())
//			} else {
//				assert.NoError(t, err)
//				assert.Equal(t, tt.expectedMessage, resp.Message)
//			}
//
//			mockRepo.AssertExpectations(t)
//			mockChannel.AssertExpectations(t) // Проверяем вызовы мок-канала
//		})
//	}
//}
