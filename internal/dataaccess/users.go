package users

import (
	"github.com/kevannnn/dineder-backend/internal/database"
	"github.com/kevannnn/dineder-backend/internal/models"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, newUser models.User) error {
	result := db.Create(&newUser)
	return result.Error
}