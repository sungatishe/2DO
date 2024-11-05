package app

import (
	"auth-service/config/db"
	"auth-service/config/logs"
	"auth-service/config/rabbitmq"
	"auth-service/internal/repository"
	"auth-service/internal/server"
	"auth-service/internal/service"
	"os"
)

func Run() {
	port := os.Getenv("PORT")

	db.InitDb()
	logs.InitLogger()

	conn, ch := rabbitmq.InitRabbitMQ()
	defer conn.Close()
	defer ch.Close()

	rabbitmq.DeclareQueue(ch, "userQueue")

	authRepo := repository.NewAuthRepository(db.DB)
	authService := service.NewAuthService(authRepo, ch)

	authServer := server.NewAuthServer(authService)

	authServer.Run(port)

}
