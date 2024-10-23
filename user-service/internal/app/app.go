package app

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"user-service/config/db"
	"user-service/internal/proto"
	"user-service/internal/repository"
	"user-service/internal/service"
)

func Run() {
	db.InitDB()

	userRepo := repository.NewUserRepository(db.DB)

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
