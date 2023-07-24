package dataaccess

import (
	"fmt"
	"github.com/kevannnn/dineder-backend/internal/models"
	"gorm.io/gorm"
)

func CreateDS(db *gorm.DB, newDS models.DinederSlot) error {
	result := db.Table("dineder_slots").Create(&newDS)
	if result.Error != nil {
		fmt.Println("Error creating dineder slot: ", result.Error)
	}
	return result.Error
}

func ReadTSID(db *gorm.DB, tsid string, user_id string) ([]models.User, error) {
	var users []models.User
	result := db.
	Joins("JOIN dineder_slots ON dineder_user.id = dineder_slots.user_id").
	Where("dineder_slots.timeslot_id = ? AND dineder_slots.user_id <> ?", tsid, user_id).
	Find(&users)
	if result.Error != nil {
        return users, result.Error
    }
	return users, nil
}
