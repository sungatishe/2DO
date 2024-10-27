package app

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"push/internal/client"
	"push/internal/proto/proto"
	"push/internal/service"
)

func Run() {
	todoClient := client.NewTodoClient(os.Getenv("TODO_SERVICE_URL"))

	pushService := service.NewNotificationService(todoClient)

	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterNotificationServiceServer(grpcServer, pushService)
	fmt.Println("Auth service is running on port :50054")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
