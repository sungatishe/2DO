package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

func ConnectToRabbitMQ() (*amqp.Connection, error) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return nil, err
	}
	return conn, nil
}
