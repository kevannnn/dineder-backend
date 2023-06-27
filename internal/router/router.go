package router

import (
	"github.com/kevannnn/dineder-backend/internal/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func Setup() chi.Router {
	r := chi.NewRouter()
	
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:19006"}, 
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(corsMiddleware.Handler)
	setUpRoutes(r)
	return r
}

func setUpRoutes(r chi.Router) {
	r.Group(routes.GetRoutes())
}
