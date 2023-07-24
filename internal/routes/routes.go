package routes

import (
	"github.com/go-chi/chi"
	handlers "github.com/kevannnn/dineder-backend/internal/handlers"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/user/{email}/{password}", handlers.GetUserID)
		r.Post("/user", handlers.PostUser)
		r.Patch("/user/{id}/gender", handlers.UpdateUserGender)
		r.Patch("/user/{id}/prefGender", handlers.UpdateUserPrefGender)
		r.Patch("/user/{id}/faculty", handlers.UpdateUserFaculty)
		r.Patch("/user/{id}/mealPref", handlers.UpdateUserMealPref)
		r.Post("/timeslot", handlers.PostTimeslot)
		r.Post("/dinederslot", handlers.PostDinederSlot)
		r.Get("/dinederslot/{user_id}/{tsid}", handlers.GetAvailUsers)
		r.Post("/match", handlers.PostMatch)
	}
}
