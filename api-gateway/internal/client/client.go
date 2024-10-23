package client

import (
	"api-gateway/internal/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type AuthClient struct {
	Client proto.AuthServiceClient
}

func NewAuthClient(address string) *AuthClient {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	return &AuthClient{Client: proto.NewAuthServiceClient(conn)}
}
