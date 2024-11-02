package app

import (
	"auth-service/config/db"
	"auth-service/config/logs"
	"auth-service/config/rabbitmq"
	"auth-service/internal/repository"
	"auth-service/internal/server"
)

func Run() {
	db.InitDb()
	logs.InitLogger()

	conn, ch := rabbitmq.InitRabbitMQ()
	defer conn.Close()
	defer ch.Close()

	rabbitmq.DeclareQueue(ch, "userQueue")

	authRepo := repository.NewAuthRepository(db.Db)
	server.InitGRPCServer(":50051", ch, authRepo)
}
