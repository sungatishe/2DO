package handlers

import (
	"api-gateway/internal/client"
	proto "api-gateway/internal/proto/push"
	"api-gateway/pgk/utils"
	"context"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

type PushHandlers struct {
	proto.UnimplementedNotificationServiceServer
	clients   map[string]*websocket.Conn
	clientsMu sync.Mutex
	client    *client.PushClient
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewPushHandlers(client *client.PushClient) *PushHandlers {
	return &PushHandlers{
		client:  client,
		clients: make(map[string]*websocket.Conn),
	}

}

func (h *PushHandlers) WebSocketHandler(rw http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		http.Error(rw, "Failed to upgrade to WebSocket", http.StatusInternalServerError)
		return
	}

	id, err := utils.ExtractUserIdFromToken(r)
	if err != nil || id == "" {
		http.Error(rw, "Unauthorized", http.StatusUnauthorized)
		return
	}

	h.clientsMu.Lock()
	h.clients[id] = conn
	h.clientsMu.Unlock()

	go h.sendNotifications(id)
}

func (h *PushHandlers) sendNotifications(userID string) {
	for {
		// Каждые 10 минут проверяем на дедлайны
		time.Sleep(10 * time.Minute)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Получаем уведомления от push-service
		resp, err := h.client.Client.NotifyOnDeadlineCheck(ctx, &proto.DeadlineNotificationRequest{})
		if err != nil {
			log.Printf("failed to get notifications: %v", err)
			continue
		}

		// Отправляем уведомления пользователю через WebSocket
		h.clientsMu.Lock()
		clientConn, exists := h.clients[userID]
		h.clientsMu.Unlock()

		if exists {
			for _, notification := range resp.Notifications {
				if notification.UserId == userID {
					err := clientConn.WriteJSON(notification)
					if err != nil {
						log.Printf("failed to send notification: %v", err)
						clientConn.Close()
						h.clientsMu.Lock()
						delete(h.clients, userID)
						h.clientsMu.Unlock()
						break
					}
				}
			}
		}
	}
}
