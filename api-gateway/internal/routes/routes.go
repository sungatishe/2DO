package routes

import (
	"api-gateway/internal/handlers"
	"api-gateway/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Routes struct {
	router chi.Router
}

func NewRouter(r chi.Router) *Routes {
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	return &Routes{router: r}
}

func (r *Routes) SetupAuthRoutes(authHandlers *handlers.AuthHandlers) {
	r.router.Post("/register", authHandlers.Register)
	r.router.Post("/login", authHandlers.Login)
	r.router.Post("/logout", authHandlers.Logout)
}

func (r *Routes) SetupUserRoutes(userHandlers *handlers.UserHandlers) {
	r.router.Post("/user", middleware.AuthMiddleware(userHandlers.CreateUser))
	r.router.Get("/user", middleware.AuthMiddleware(userHandlers.GetUserById))
	r.router.Put("/user", middleware.AuthMiddleware(userHandlers.UpdateUser))
	//r.router.Delete("/user/{id}", middleware.AuthMiddleware(userHandlers.DeleteUser))
}

func (r *Routes) SetupTodoRoutes(todoHandlers *handlers.TodoHandlers) {
	r.router.Post("/todo", middleware.AuthMiddleware(todoHandlers.CreateTodo))
	r.router.Get("/todo/{id}", middleware.AuthMiddleware(todoHandlers.GetTodoById))
	r.router.Get("/todo", middleware.AuthMiddleware(todoHandlers.GetUsersListTodo))
	r.router.Put("/todo", middleware.AuthMiddleware(todoHandlers.UpdateTodo))
	r.router.Delete("/todo/{id}", middleware.AuthMiddleware(todoHandlers.DeleteTodo))
}
