package dataaccess

import (
	"fmt"

	"github.com/kevannnn/dineder-backend/internal/models"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func CreateDS(db *gorm.DB, newDS models.DinederSlot) error {
	result := db.Table("dineder_slots").Create(&newDS)
	if result.Error != nil {
		fmt.Println("Error creating dineder slot: ", result.Error)
	}
	return result.Error
}

func GetExistingDS(db*gorm.DB, tsid uint, user_id uuid.UUID)(models.DinederSlot) {
	var dinederSlot models.DinederSlot
	result := db.First(&dinederSlot, "user_id = ? AND timeslot_id = ?", user_id, tsid)
	if result.Error != nil {
		fmt.Println("Error fetching timeslot: ", result.Error)
	}
	return dinederSlot
}


func ReadTSID(db *gorm.DB, tsid string, user_id string) ([]models.User, error) {
	var users []models.User
	
	user, err := GetUserFromID(db, user_id)
    if err != nil {
        return nil, err
    }
	result := db.
	Joins("JOIN dineder_slots ON dineder_user.id = dineder_slots.user_id").
	Where("dineder_slots.timeslot_id = ? AND dineder_user.gender = ? AND dineder_slots.user_id <> ?", tsid, user.PrefGender, user_id).
	Find(&users)
	if result.Error != nil {
        return users, result.Error
    }
	return users, nil
}
