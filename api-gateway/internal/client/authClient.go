package client

import (
	"api-gateway/internal/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type AuthClient struct {
	Client auth.AuthServiceClient
}

func NewAuthClient(address string) *AuthClient {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	return &AuthClient{Client: auth.NewAuthServiceClient(conn)}
}