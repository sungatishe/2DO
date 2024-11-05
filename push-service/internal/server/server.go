package server

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"push/internal/proto/proto"
)

type PushServer struct {
	server      *grpc.Server
	pushService proto.NotificationServiceServer
}

func NewPushServer(pushService proto.NotificationServiceServer) *PushServer {
	grpcServer := grpc.NewServer()
	proto.RegisterNotificationServiceServer(grpcServer, pushService)

	return &PushServer{
		server:      grpcServer,
		pushService: pushService,
	}
}

func (s *PushServer) Run(port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Printf("Push service is running on port %s\n", port)
	if err := s.server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
