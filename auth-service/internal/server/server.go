package server

import (
	"auth-service/internal/proto"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"log"
	"net"
)

func InitGRPCServer(port string, channel *amqp.Channel, authRepo repository.AuthRepository) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userService := service.NewAuthService(authRepo, channel)
	proto.RegisterAuthServiceServer(grpcServer, userService)

	fmt.Println("Auth service is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
