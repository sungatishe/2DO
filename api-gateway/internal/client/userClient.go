package client

import (
	"api-gateway/internal/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type UserClient struct {
	Client user.UserServiceClient
}

func NewUserClient(address string) *UserClient {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	return &UserClient{Client: user.NewUserServiceClient(conn)}
}
