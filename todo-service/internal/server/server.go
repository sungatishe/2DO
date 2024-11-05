package server

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"todo-service/internal/proto/proto"
)

type TodoServer struct {
	server      *grpc.Server
	todoService proto.TodoServiceServer
}

func NewTodoServer(todoService proto.TodoServiceServer) *TodoServer {
	grpcServer := grpc.NewServer()
	proto.RegisterTodoServiceServer(grpcServer, todoService)

	return &TodoServer{
		server:      grpcServer,
		todoService: todoService,
	}
}

func (s *TodoServer) Run(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Printf("Todo service is running on port %s\n", port)
	if err := s.server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
