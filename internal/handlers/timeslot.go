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

func PostTimeslot(w http.ResponseWriter, req *http.Request) {
	var newTS models.Timeslot
	err := json.NewDecoder(req.Body).Decode(&newTS)

	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	existingTS, err := dataaccess.GetTimeslotByTime(database.DB, newTS.Time)
	if existingTS.ID != 0 {
		TSExistResponse := api.Response{
			Data: json.RawMessage(
				fmt.Sprintf(`{
					"message": "Exisiting timeslot at %s",
					"tsid": "%d"
					}`, existingTS.Time.String(), existingTS.ID),
				),
		}	
		//w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(TSExistResponse)
		if err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	} else {
		newTS, err = dataaccess.CreateTS(database.DB, newTS)
	
		if err != nil {
		http.Error(w, "Failed to create timeslot", http.StatusInternalServerError)
		return
		}

		response := api.Response{
			Data: json.RawMessage(
				fmt.Sprintf(`{
					"message": "Timeslot created at %s",
					"tsid": "%d"
					}`, newTS.Time.String(), newTS.ID),
				),
		}	

		err = json.NewEncoder(w).Encode(response)

		if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	}
}