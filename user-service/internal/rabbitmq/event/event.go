package event

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
	"log"
	"user-service/internal/models"
	"user-service/internal/repository"
)

type UserRegisteredEvent struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func ListenForUserRegisteredEvents(ch *amqp.Channel, repo repository.UserRepository) {
	msgs, err := ch.Consume(
		"userQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	for msg := range msgs {
		var event UserRegisteredEvent
		err := json.Unmarshal(msg.Body, &event)
		if err != nil {
			log.Printf("Failed to unmarshal message: %s", err)
		}
		createUserInDatabase(event, repo)
	}
}

func createUserInDatabase(event UserRegisteredEvent, repo repository.UserRepository) {
	user := models.User{
		Model:       gorm.Model{ID: event.UserID},
		Username:    event.Username,
		Email:       event.Email,
		Description: event.Username,
		Avatar:      "https://greekherald.com.au/wp-content/uploads/2020/07/default-avatar.png",
	}

	err := repo.CreateUser(&user)
	if err != nil {
		log.Printf("Failed to create user: %s", err)
		return
	}

	log.Printf("User created: %s", user.Username)
}
