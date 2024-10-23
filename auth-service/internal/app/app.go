package app

import (
	"auth-service/config/db"
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

	authRepo := repository.NewAuthRepository(db.Db)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userService := service.NewAuthService(authRepo)
	proto.RegisterAuthServiceServer(grpcServer, userService)

	fmt.Println("Auth service is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
