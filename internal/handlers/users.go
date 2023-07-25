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
    email := chi.URLParam(req, "email")
	password := chi.URLParam(req, "password")
	
	user, err := dataaccess.ReadUserID(database.DB, email, password)
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

	newUser, err = dataaccess.CreateUser(database.DB, newUser)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	response := api.Response{
		Data: json.RawMessage(fmt.Sprintf(`{
			"message": "Welcome to dineder, %s",
			"user_id": "%s"
			}`, newUser.Name, newUser.ID)),
	}	

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func UpdateUserProfilePic(w http.ResponseWriter, req *http.Request) {
	UserID := chi.URLParam(req, "id")
	var requestBody struct {
		ImageUrl  string  `json:"image_url"`
	}
	
	err := json.NewDecoder(req.Body).Decode(&requestBody) 
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	
	updatedUser, err := dataaccess.UpdateUserField(database.DB, UserID, "ImageUrl", requestBody.ImageUrl)
	if err != nil {
		http.Error(w, "Failed to update user profile picture", http.StatusInternalServerError)
		return
	}

	response := api.Response{
		Data: json.RawMessage(
			fmt.Sprintf(`{"message": "Updated profile picture for %s as %d"}`, updatedUser.Name, updatedUser.MealPref)),
	}	

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func UpdateUserGender(w http.ResponseWriter, req *http.Request) {
	UserID := chi.URLParam(req, "id")
	var requestBody struct {
		Gender int `json:"gender"`
	}
	
	err := json.NewDecoder(req.Body).Decode(&requestBody) 
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	
	updatedUser, err := dataaccess.UpdateUserField(database.DB, UserID, "Gender", requestBody.Gender)
	if err != nil {
		http.Error(w, "Failed to update user gender", http.StatusInternalServerError)
		return
	}

	response := api.Response{
		Data: json.RawMessage(
			fmt.Sprintf(`{"message": "Updated gender for %s as %d"}`, updatedUser.Name, updatedUser.Gender)),
	}	

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func UpdateUserPrefGender(w http.ResponseWriter, req *http.Request) {
	UserID := chi.URLParam(req, "id")
	var requestBody struct {
		PrefGender 	int `json:"pref_gender"`
	}
	
	err := json.NewDecoder(req.Body).Decode(&requestBody) 
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	
	updatedUser, err := dataaccess.UpdateUserField(database.DB, UserID, "PrefGender", requestBody.PrefGender)
	if err != nil {
		http.Error(w, "Failed to update user preferred gender", http.StatusInternalServerError)
		return
	}

	response := api.Response{
		Data: json.RawMessage(
			fmt.Sprintf(`{"message": "Updated preferred gender for %s as %d"}`, updatedUser.Name, updatedUser.PrefGender)),
	}	

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func UpdateUserFaculty(w http.ResponseWriter, req *http.Request) {
	UserID := chi.URLParam(req, "id")
	var requestBody struct {
		Faculty	int  `json:"faculty"`
	}
	
	err := json.NewDecoder(req.Body).Decode(&requestBody) 
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	
	updatedUser, err := dataaccess.UpdateUserField(database.DB, UserID, "Faculty", requestBody.Faculty)
	if err != nil {
		http.Error(w, "Failed to update user faculty", http.StatusInternalServerError)
		return
	}

	response := api.Response{
		Data: json.RawMessage(
			fmt.Sprintf(`{"message": "Updated faculty for %s as %d"}`, updatedUser.Name, updatedUser.Faculty)),
	}	

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func UpdateUserMealPref(w http.ResponseWriter, req *http.Request) {
	UserID := chi.URLParam(req, "id")
	var requestBody struct {
		MealPref  int  `json:"meal_pref"`
	}
	
	err := json.NewDecoder(req.Body).Decode(&requestBody) 
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	
	updatedUser, err := dataaccess.UpdateUserField(database.DB, UserID, "MealPref", requestBody.MealPref)
	if err != nil {
		http.Error(w, "Failed to update user meal preference", http.StatusInternalServerError)
		return
	}

	response := api.Response{
		Data: json.RawMessage(
			fmt.Sprintf(`{"message": "Updated meal preference for %s as %d"}`, updatedUser.Name, updatedUser.MealPref)),
	}	

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}