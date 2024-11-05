package app

import (
	"os"
	"push/config/cache"
	"push/internal/client"
	"push/internal/server"
	"push/internal/service"
)

func Run() {
	port := os.Getenv("PORT")
	todoClient := client.NewTodoClient(os.Getenv("TODO_SERVICE_URL"))
	redisClient := cache.NewRedisClient()

	pushService := service.NewNotificationService(todoClient, redisClient)

	pushServer := server.NewPushServer(pushService)

	pushServer.Run(port)
}
