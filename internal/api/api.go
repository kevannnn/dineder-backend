package api

import (
	"encoding/json"
	"github.com/kevannnn/dineder-backend/internal/models"
)

type Response struct {
	Data json.RawMessage `json:"data,omitempty"`
}

type GetAvailUsersResponse struct {
    Users []models.User `json:"users"`
}