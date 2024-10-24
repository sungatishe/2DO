package app

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"user-service/config/db"
	"user-service/config/rabbitmq"
	"user-service/internal/proto"
	"user-service/internal/rabbitmq/event"
	"user-service/internal/repository"
	"user-service/internal/service"
)

func Run() {
	db.InitDB()

	conn, err := rabbitmq.ConnectToRabbitMQ()
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
		"userQueue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}

	userRepo := repository.NewUserRepository(db.DB)

	go event.ListenForUserRegisteredEvents(ch, userRepo)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userService := service.NewUserService(userRepo)
	proto.RegisterUserServiceServer(grpcServer, userService)

	fmt.Println("User service is running on port :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
