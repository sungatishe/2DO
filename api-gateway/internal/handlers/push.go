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

	log.Printf("Users push id: %v", id)

	h.clientsMu.Lock()
	h.clients[id] = conn
	h.clientsMu.Unlock()

	go h.sendNotifications(id)
}

func (h *PushHandlers) sendNotifications(userID string) {
	for {
		time.Sleep(20 * time.Minute)
		log.Println("Push sent!!!")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		resp, err := h.client.Client.NotifyOnDeadlineCheck(ctx, &proto.DeadlineNotificationRequest{})
		if err != nil {
			log.Printf("failed to get notifications: %v", err)
			continue
		}

		log.Printf("Number of notifications received: %d", len(resp.Notifications))

		// Отправляем уведомления пользователю через WebSocket
		h.clientsMu.Lock()
		clientConn, exists := h.clients[userID]
		h.clientsMu.Unlock()
		log.Println(exists)

		if exists {
			for _, notification := range resp.Notifications {
				log.Printf("Received notification: %+v", notification) // Логируем каждое уведомление

				if notification.UserId == userID {
					log.Println("Sending notification to user:", userID)
					log.Println("Notification Description:", notification.Description)
					log.Println("Notification Deadline:", notification.Deadline)

					err := clientConn.WriteJSON(notification)
					if err != nil {
						log.Printf("failed to send notification: %v", err)
						clientConn.Close()
						h.clientsMu.Lock()
						delete(h.clients, userID)
						h.clientsMu.Unlock()
						break
					} else {
						log.Println("Notification sent successfully") // Логируем успешную отправку уведомления
					}
				}
			}
		}
	}
}
