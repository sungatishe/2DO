package client

import (
	"api-gateway/internal/proto/todo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type TodoClient struct {
	Client todo.TodoServiceClient
}

func NewTodoClient(address string) *TodoClient {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	return &TodoClient{Client: todo.NewTodoServiceClient(conn)}
}
