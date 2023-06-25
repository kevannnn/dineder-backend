package routes

import (
	"github.com/go-chi/chi"
	handlers "github.com/kevannnn/dineder-backend/internal/handlers"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/user/{email}/{password}", handlers.GetUserID)
		r.Post("/user", handlers.PostUser)
	}
}
