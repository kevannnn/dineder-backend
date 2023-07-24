package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kevannnn/dineder-backend/internal/api"
	"github.com/kevannnn/dineder-backend/internal/dataaccess"
	"github.com/kevannnn/dineder-backend/internal/database"
	"github.com/kevannnn/dineder-backend/internal/models"
)

func PostDinederSlot(w http.ResponseWriter, req *http.Request) {
	var newDS models.DinederSlot
	err := json.NewDecoder(req.Body).Decode(&newDS)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = dataaccess.CreateDS(database.DB, newDS)
	if err != nil {
		http.Error(w, "Failed to create dineder slot", http.StatusInternalServerError)
		return
	}

	userName, err := dataaccess.GetUserNameFromID(database.DB, newDS.UserID)
	if err != nil {
		http.Error(w, "Failed to fetch user's name", http.StatusInternalServerError)
		return
	}

	timeslotTime, err := dataaccess.GetTSTimeFromID(database.DB, newDS.TimeslotID)
	if err != nil {
		http.Error(w, "Failed to fetch user's name", http.StatusInternalServerError)
		return
	}

	response := api.Response{
		Data: json.RawMessage(
			fmt.Sprintf(`{"message": "Dineder slot created at %s by %s"}`, timeslotTime, userName),
			),
	}	

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func GetAvailUsers(w http.ResponseWriter, req *http.Request) {
    tsid := chi.URLParam(req, "tsid")
	user_id := chi.URLParam(req, "user_id")
	
	users, err := dataaccess.ReadTSID(database.DB, tsid, user_id)
    if err != nil {
        http.Error(w, "Failed to retrieve available users", http.StatusInternalServerError)
        return
    }

	response := api.GetAvailUsersResponse{
        Users: users,
    }
   
	if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "failed to encode response", http.StatusInternalServerError)
        return
    }
}