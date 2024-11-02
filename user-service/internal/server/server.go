package server

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"user-service/internal/proto"
	"user-service/internal/repository"
	"user-service/internal/service"
)

func InitGRPCServer(port string, userRepo repository.UserRepository) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userService := service.NewUserService(userRepo)
	proto.RegisterUserServiceServer(grpcServer, userService)

	log.Printf("User service is running on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
