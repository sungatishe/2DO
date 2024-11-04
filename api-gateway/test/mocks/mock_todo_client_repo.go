package mocks

import (
	"api-gateway/internal/client"
	"api-gateway/internal/proto/todo"
	"context"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type MockTodoClient struct {
	mock.Mock
	client.TodoClient
}

func (m *MockTodoClient) CreateTodo(ctx context.Context, req *todo.CreateTodoRequest, opts ...grpc.CallOption) (*todo.CreateTodoResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*todo.CreateTodoResponse), args.Error(1)
}

func (m *MockTodoClient) GetTodoById(ctx context.Context, req *todo.GetTodosByIdRequest, opts ...grpc.CallOption) (*todo.GetTodosByIdResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*todo.GetTodosByIdResponse), args.Error(1)
}

func (m *MockTodoClient) UpdateTodo(ctx context.Context, req *todo.UpdateTodoRequest, opts ...grpc.CallOption) (*todo.UpdateTodoResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*todo.UpdateTodoResponse), args.Error(1)
}

func (m *MockTodoClient) DeleteTodo(ctx context.Context, req *todo.DeleteTodoRequest, opts ...grpc.CallOption) (*todo.DeleteTodoResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*todo.DeleteTodoResponse), args.Error(1)
}

func (m *MockTodoClient) ListTodo(ctx context.Context, req *todo.ListTodoRequest, opts ...grpc.CallOption) (*todo.ListTodoResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*todo.ListTodoResponse), args.Error(1)
}
