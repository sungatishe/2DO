package service

import (
	"context"
	"encoding/json"
	"log"
	"push/config/cache"
	"push/internal/client"
	"push/internal/proto/proto"
	"push/internal/proto/todo"
	"strconv"
	"time"
)

type NotificationService struct {
	proto.UnimplementedNotificationServiceServer
	todoClient  *client.TodoClient
	redisClient *cache.RedisClient
}

func NewNotificationService(todoClient *client.TodoClient, redisClient *cache.RedisClient) proto.NotificationServiceServer {
	return &NotificationService{todoClient: todoClient, redisClient: redisClient}
}

func (s *NotificationService) NotifyOnDeadlineCheck(ctx context.Context, req *proto.DeadlineNotificationRequest) (*proto.DeadlineNotificationResponse, error) {
	deadline := time.Now().Add(12 * time.Hour).Format(time.RFC3339)
	log.Printf("Requesting todos with deadline before: %s", deadline)

	cachedNotifications, err := s.redisClient.GetNotification(ctx, "deadline_notifications")
	if err == nil && cachedNotifications != "" {
		log.Println("Returning cached notifications")
		return &proto.DeadlineNotificationResponse{
			Notifications: parseNotifications(cachedNotifications),
		}, nil
	}

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

	err = s.redisClient.SetNotification(ctx, "deadline_notifications", formatNotifications(notifications), 3*time.Hour)
	if err != nil {
		log.Printf("Failed to cache notifications: %v", err)
	}

	log.Printf("Number of notifications created: %d", len(notifications))
	return &proto.DeadlineNotificationResponse{Notifications: notifications}, nil
}

func parseNotifications(data string) []*proto.DeadlineNotification {
	var notifications []*proto.DeadlineNotification
	err := json.Unmarshal([]byte(data), &notifications)
	if err != nil {
		log.Printf("Failed to deserialize notifications: %v", err)
		return nil
	}
	return notifications
}

func formatNotifications(notifications []*proto.DeadlineNotification) string {
	data, err := json.Marshal(notifications)
	if err != nil {
		log.Printf("Failed to serialize notifications: %v", err)
		return ""
	}
	return string(data)
}
