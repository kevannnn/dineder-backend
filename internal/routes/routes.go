package routes

import (
	"encoding/json"
	"net/http"

	"github.com/kevannnn/dineder-backend/internal/handlers/users"
	"github.com/go-chi/chi"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/user/{username}/{password}", handlers.GetUser)
		r.Post("/user", handlers.createUser)
	}
}
