package service

import (
	"context"
	"push/internal/client"
	"push/internal/proto/proto"
	"push/internal/proto/todo"
	"strconv"
	"time"
)

type NotificationService struct {
	proto.UnimplementedNotificationServiceServer
	todoClient *client.TodoClient
}

func NewNotificationService(todoClient *client.TodoClient) proto.NotificationServiceServer {
	return &NotificationService{todoClient: todoClient}
}

func (s *NotificationService) NotifyOnDeadlineCheck(ctx context.Context, req *proto.DeadlineNotificationRequest) (*proto.DeadlineNotificationResponse, error) {
	deadline := time.Now().Add(12 * time.Hour).Format(time.RFC3339)

	todosResp, err := s.todoClient.Client.GetTodosByDeadline(context.Background(), &todo.GetTodosByDeadlineRequest{Deadline: deadline})
	if err != nil {
		return nil, err
	}

	var notifications []*proto.DeadlineNotification
	for _, todo := range todosResp.Todos {
		notifications = append(notifications, &proto.DeadlineNotification{
			Id:          strconv.FormatUint(todo.Id, 10),
			UserId:      strconv.FormatUint(todo.UserId, 10),
			Title:       todo.Title,
			Description: todo.Description,
			Deadline:    todo.Deadline,
		})
	}

	return &proto.DeadlineNotificationResponse{Notifications: notifications}, nil

}
