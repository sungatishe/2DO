package service

import (
	"context"
	"log"
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
	log.Printf("Requesting todos with deadline before: %s", deadline)

	todosResp, err := s.todoClient.Client.GetTodosByDeadline(context.Background(), &todo.GetTodosByDeadlineRequest{Deadline: deadline})
	if err != nil {
		log.Printf("Error fetching todos by deadline: %v", err)
		return nil, err
	}
	log.Printf("Number of todos fetched: %d", len(todosResp.Todos))

	var notifications []*proto.DeadlineNotification
	for _, todo := range todosResp.Todos {
		log.Printf("Creating notification for Todo: ID=%d, UserID=%d, Title=%s, Deadline=%s", todo.Id, todo.UserId, todo.Title, todo.Deadline)
		notifications = append(notifications, &proto.DeadlineNotification{
			Id:          strconv.FormatUint(todo.Id, 10),
			UserId:      strconv.FormatUint(todo.UserId, 10),
			Title:       todo.Title,
			Description: todo.Description,
			Deadline:    todo.Deadline,
		})
	}

	log.Printf("Number of notifications created: %d", len(notifications))
	return &proto.DeadlineNotificationResponse{Notifications: notifications}, nil
}
