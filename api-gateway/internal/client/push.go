package client

import (
	push "api-gateway/internal/proto/push"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type PushClient struct {
	Client push.NotificationServiceClient
}

func NewPushClient(address string) *PushClient {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	return &PushClient{Client: push.NewNotificationServiceClient(conn)}
}
