package routes

import (
	"api-gateway/internal/handlers"
	"github.com/go-chi/chi/v5"
)

type Routes struct {
	router chi.Router
}

func NewRouter(r chi.Router) *Routes {
	return &Routes{router: r}
}

func (r *Routes) SetupAuthRoutes(authHandlers *handlers.AuthHandlers) {
	r.router.Post("/register", authHandlers.Register)
	r.router.Post("/login", authHandlers.Login)
}

func (r *Routes) SetupUserRoutes(userHandlers *handlers.UserHandlers) {
	r.router.Post("/user", userHandlers.CreateUser)
	r.router.Get("/user/{id}", userHandlers.GetUserById)
	r.router.Put("/user", userHandlers.UpdateUser)
	r.router.Delete("/user/{id}", userHandlers.DeleteUser)
}
