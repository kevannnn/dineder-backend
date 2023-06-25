package dataaccess

import (
	"github.com/kevannnn/dineder-backend/internal/models"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, newUser models.User) error {
	result := db.Table("dineder_user").Create(&newUser)
	return result.Error
}