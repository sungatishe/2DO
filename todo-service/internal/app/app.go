package app

import (
	"os"
	"todo-service/config/db"
	"todo-service/internal/repository"
	"todo-service/internal/server"
	"todo-service/internal/service"
)

func Run() {
	port := os.Getenv("PORT")
	db.InitDB()

	todoRepo := repository.NewTodoRepository(db.DB)
	todoService := service.NewTodoService(todoRepo)

	pushServer := server.NewTodoServer(todoService)

	pushServer.Run(port)
}
