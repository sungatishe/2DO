package service

import (
	"context"
	"time"
	"todo-service/internal/models"
	"todo-service/internal/proto/proto"
	"todo-service/internal/repository"
)

type todoService struct {
	repo repository.TodoRepository
	proto.UnimplementedTodoServiceServer
}

func NewTodoService(repo repository.TodoRepository) proto.TodoServiceServer {
	return &todoService{repo: repo}
}

func (s *todoService) CreateTodo(ctx context.Context, req *proto.CreateTodoRequest) (*proto.CreateTodoResponse, error) {
	deadline, err := time.Parse(time.RFC3339, req.Deadline)
	if err != nil {
		return nil, err
	}

	todo := &models.Todo{
		UserID:      uint(req.UserId),
		Title:       req.Title,
		Description: req.Description,
		IsDone:      req.IsDone,
		Deadline:    deadline,
	}

	err = s.repo.CreateTodo(todo)
	if err != nil {
		return nil, err
	}

	return &proto.CreateTodoResponse{
		Message: "Todo created successfully",
		Todo: &proto.Todo{
			Id:          uint64(todo.ID),
			UserId:      uint64(todo.UserID),
			Title:       todo.Title,
			Description: todo.Description,
			IsDone:      todo.IsDone,
			Deadline:    todo.Deadline.Format(time.RFC3339),
		},
	}, nil

}

func (s *todoService) GetTodoById(ctx context.Context, req *proto.GetTodosByIdRequest) (*proto.GetTodosByIdResponse, error) {
	todo, err := s.repo.GetTodoById(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &proto.GetTodosByIdResponse{Todo: &proto.Todo{
		Id:          uint64(todo.ID),
		UserId:      uint64(todo.UserID),
		Title:       todo.Title,
		Description: todo.Description,
		IsDone:      todo.IsDone,
		Deadline:    todo.Deadline.Format(time.RFC3339),
	}}, nil
}

func (s *todoService) UpdateTodo(ctx context.Context, req *proto.UpdateTodoRequest) (*proto.UpdateTodoResponse, error) {
	todo, err := s.repo.GetTodoById(uint(req.Id))
	if err != nil {
		return nil, err
	}
	todo.Title = req.Title
	todo.Description = req.Description
	todo.IsDone = req.IsDone

	err = s.repo.UpdateTodo(todo)
	if err != nil {
		return nil, err
	}

	return &proto.UpdateTodoResponse{Todo: &proto.Todo{
		Id:          uint64(todo.ID),
		UserId:      uint64(todo.UserID),
		Title:       todo.Title,
		Description: todo.Description,
		IsDone:      todo.IsDone,
		Deadline:    todo.Deadline.Format(time.RFC3339),
	}}, nil
}

func (s *todoService) DeleteTodo(ctx context.Context, req *proto.DeleteTodoRequest) (*proto.DeleteTodoResponse, error) {
	err := s.repo.DeleteTodoById(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &proto.DeleteTodoResponse{Message: "Todo deleted successfully"}, nil
}

func (s *todoService) ListTodo(ctx context.Context, req *proto.ListTodoRequest) (*proto.ListTodoResponse, error) {
	todos, err := s.repo.ListTodoByUserId(uint(req.UserId))
	if err != nil {
		return nil, err
	}
	var todoList []*proto.Todo
	for _, todo := range todos {
		todoList = append(todoList, &proto.Todo{
			Id:          uint64(todo.ID),
			UserId:      uint64(todo.UserID),
			Title:       todo.Title,
			Description: todo.Description,
			IsDone:      todo.IsDone,
			Deadline:    todo.Deadline.Format(time.RFC3339),
		})
	}

	return &proto.ListTodoResponse{Todos: todoList}, nil
}

func (s *todoService) GetTodosByDeadline(ctx context.Context, req *proto.GetTodosByDeadlineRequest) (*proto.GetTodosByDeadlineResponse, error) {
	deadlineTime := time.Now().Add(12 * time.Hour).Format(time.RFC3339)
	todos, err := s.repo.GetTodosByDeadline(deadlineTime)
	if err != nil {
		return nil, err
	}
	var todoList []*proto.Todo
	for _, todo := range todos {
		todoList = append(todoList, &proto.Todo{
			Id:          uint64(todo.ID),
			UserId:      uint64(todo.UserID),
			Title:       todo.Title,
			Description: todo.Description,
			IsDone:      todo.IsDone,
			Deadline:    todo.Deadline.Format(time.RFC3339),
		})
	}
	return &proto.GetTodosByDeadlineResponse{Todos: todoList}, nil
}
