package app

import (
	"api-gateway/internal/client"
	"api-gateway/internal/handlers"
	"api-gateway/internal/routes"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
)

func Run() {
	authClient := client.NewAuthClient(os.Getenv("AUTH_SERVICE_URL"))
	userClient := client.NewUserClient(os.Getenv("USER_SERIVE_URL"))

	authHandler := handlers.NewAuthHandlers(authClient)
	userHandler := handlers.NewUserHandlers(userClient)

	r := chi.NewRouter()

	rt := routes.NewRouter(r)
	rt.SetupAuthRoutes(authHandler)
	rt.SetupUserRoutes(userHandler)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Error in server: %v", err)
	}
}
