package app

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"todo-service/config/db"
	"todo-service/internal/proto/proto"
	"todo-service/internal/repository"
	"todo-service/internal/service"
)

func Run() {
	db.InitDB()

	todoRepo := repository.NewTodoRepository(db.DB)
	todoService := service.NewTodoService(todoRepo)

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterTodoServiceServer(grpcServer, todoService)
	fmt.Println("Auth service is running on port :50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
