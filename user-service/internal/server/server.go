package server

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"user-service/internal/proto"
)

type UserServer struct {
	server      *grpc.Server
	userService proto.UserServiceServer
}

func NewUserServer(userService proto.UserServiceServer) *UserServer {
	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, userService)

	return &UserServer{
		server:      grpcServer,
		userService: userService,
	}
}

func (s *UserServer) Run(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Printf("User service is running on port %s\n", port)
	if err := s.server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
