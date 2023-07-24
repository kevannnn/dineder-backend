package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/kevannnn/dineder-backend/internal/api"
	"github.com/kevannnn/dineder-backend/internal/dataaccess"
	"github.com/kevannnn/dineder-backend/internal/database"
	"github.com/kevannnn/dineder-backend/internal/models"
)

func PostMatch(w http.ResponseWriter, req *http.Request) {
	var newMatch models.Match
	err := json.NewDecoder(req.Body).Decode(&newMatch)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = dataaccess.CreateMatch(database.DB, newMatch)
	if err != nil {
		http.Error(w, "Failed to create match", http.StatusInternalServerError)
		return
	}

	user1Name, err := dataaccess.GetUserNameFromID(database.DB, newMatch.User1ID)
	if err != nil {
		http.Error(w, "Failed to fetch user 1's name", http.StatusInternalServerError)
		return
	}

	user2Name, err := dataaccess.GetUserNameFromID(database.DB, newMatch.User2ID)
	if err != nil {
		http.Error(w, "Failed to fetch user 2's name", http.StatusInternalServerError)
		return
	}

	timeslotTime, err := dataaccess.GetTSTimeFromID(database.DB, newMatch.TimeslotID)
	if err != nil {
		http.Error(w, "Failed to fetch timeslot time", http.StatusInternalServerError)
		return
	}

	response := api.Response{
		Data: json.RawMessage(
			fmt.Sprintf(`{"message": "Match created with %s and %s at %s"}`,
			user1Name,
			user2Name,
			timeslotTime,
		),
	),	
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}