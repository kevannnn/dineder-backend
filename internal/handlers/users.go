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

func GetUserID(w http.ResponseWriter, req *http.Request) {
    username := chi.URLParam(req, "username")
	password := chi.URLParam(req, "password")
	
	user, err := dataaccess.ReadUserID(database.DB, username, password)
    if err != nil {
        http.Error(w, "Failed to retrieve user ID", http.StatusInternalServerError)
        return
    }

	response := api.Response{
        Data: json.RawMessage(fmt.Sprintf(`{"id": "%s"}`, user.ID.String())),
    }
   
	if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "failed to encode response", http.StatusInternalServerError)
        return
    }


}

func PostUser(w http.ResponseWriter, req *http.Request) {
	var newUser models.User
	err := json.NewDecoder(req.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = dataaccess.CreateUser(database.DB, newUser)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	response := api.Response{
		Data: json.RawMessage(fmt.Sprintf(`{"message": "Welcome to dineder, %s"}`, newUser.Username)),
	}	

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}