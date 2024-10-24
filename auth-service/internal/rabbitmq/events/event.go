package events

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type UserRegisteredEvent struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func PublishUserRegisteredEvent(ch *amqp.Channel, event UserRegisteredEvent) error {
	eventBody, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		"userQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        eventBody,
		},
	)
	if err != nil {
		log.Printf("Failed to publish event: %s", err)
		return err
	}
	log.Printf("Published event: %v", event)
	return nil
}
