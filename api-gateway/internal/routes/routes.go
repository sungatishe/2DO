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
