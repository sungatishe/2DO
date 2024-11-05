package app

import (
	"os"
	"user-service/config/db"
	"user-service/config/rabbitmq"
	"user-service/internal/rabbitmq/event"
	"user-service/internal/repository"
	"user-service/internal/server"
	"user-service/internal/service"
)

func Run() {
	db.InitDB()

	conn, ch := rabbitmq.InitRabbitMQ()
	defer conn.Close()
	defer ch.Close()

	rabbitmq.DeclareQueue(ch, "userQueue")

	userRepo := repository.NewUserRepository(db.DB)
	go event.ListenForUserRegisteredEvents(ch, userRepo)

	userService := service.NewUserService(userRepo)

	userServer := server.NewUserServer(userService)

	userServer.Run(os.Getenv("PORT"))
}
