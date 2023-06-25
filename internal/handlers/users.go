package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"

	//"github.com/go-chi/chi"
	"github.com/kevannnn/dineder-backend/internal/database"
	"github.com/kevannnn/dineder-backend/internal/models"
	//"github.com/kevannnn/dineder-backend/internal/dataaccess"
	"github.com/kevannnn/dineder-backend/internal/api"
	"gorm.io/gorm"
	
)

func CreateUser(db *gorm.DB, newUser models.User) error {
	result := db.Table("dineder_user2").Create(&newUser)
	if result.Error != nil {
		fmt.Println("Error creating user:", result.Error)
	}
	return result.Error
}

func PostUser(w http.ResponseWriter, req *http.Request) {
	var newUser models.User
	err := json.NewDecoder(req.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = CreateUser(database.DB, newUser)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	//error: id always returns 0 
	response := api.Response{
		Data: json.RawMessage(fmt.Sprintf(`{"id": %d, "username": "%s" }`, newUser.ID, newUser.Username)),
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}