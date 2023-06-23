package router

import (
	"github.com/kevannnn/dineder-backend/internal/routes"
	"github.com/go-chi/chi"
)

func Setup() chi.Router {
	r := chi.NewRouter()
	setUpRoutes(r)
	return r
}

func setUpRoutes(r chi.Router) {
	r.Group(routes.GetRoutes())
}
