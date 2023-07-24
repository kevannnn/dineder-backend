package dataaccess

import (
	"fmt"
	"github.com/kevannnn/dineder-backend/internal/models"
	"gorm.io/gorm"
)

func CreateMatch(db *gorm.DB, newMatch models.Match) error {
	result := db.Table("matches").Create(&newMatch)
	if result.Error != nil {
		fmt.Println("Error creating match: ", result.Error)
	}
	return result.Error
}