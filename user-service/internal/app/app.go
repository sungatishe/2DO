package app

import (
	"user-service/config/db"
	"user-service/config/rabbitmq"
	"user-service/internal/rabbitmq/event"
	"user-service/internal/repository"
	"user-service/internal/server"
)

func Run() {
	db.InitDB()

	conn, ch := rabbitmq.InitRabbitMQ()
	defer conn.Close()
	defer ch.Close()

	rabbitmq.DeclareQueue(ch, "userQueue")

	userRepo := repository.NewUserRepository(db.DB)
	go event.ListenForUserRegisteredEvents(ch, userRepo)

	server.InitGRPCServer(":50052", userRepo)
}
