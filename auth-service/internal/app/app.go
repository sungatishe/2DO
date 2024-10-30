package app

import (
	"auth-service/config/db"
	"auth-service/config/logs"
	"auth-service/config/rabbitmq"
	"auth-service/internal/proto"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Run() {
	db.InitDb()
	logs.InitLogger()

	conn, err := rabbitmq.ConnectRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	// Открытие канала RabbitMQ
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	// Убедимся, что очередь "userQueue" существует
	_, err = ch.QueueDeclare(
		"userQueue", // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	authRepo := repository.NewAuthRepository(db.Db)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userService := service.NewAuthService(authRepo, ch)
	proto.RegisterAuthServiceServer(grpcServer, userService)

	fmt.Println("Auth service is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
