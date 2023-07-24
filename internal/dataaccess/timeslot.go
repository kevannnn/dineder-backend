package dataaccess

import (
	"fmt"
	"time"

	"github.com/kevannnn/dineder-backend/internal/models"
	"gorm.io/gorm"
)

func CreateTS(db *gorm.DB, newTS models.Timeslot) (models.Timeslot, error) {
	result := db.Table("timeslots").Create(&newTS)
	if result.Error != nil {
		fmt.Println("Error creating timeslot: ", result.Error)
		return newTS, result.Error
	}
	return newTS, nil
}

func GetTSTimeFromID(db *gorm.DB, TSID uint) (string, error) {
	var timeslot models.Timeslot
	result := db.First(&timeslot, "id = ?", TSID)
	if result.Error != nil {
		fmt.Println("Error fetching timeslot: ", result.Error)
	}
	return timeslot.Time.String(), result.Error
}

func GetTimeslotByTime(db *gorm.DB, timeTS time.Time) (models.Timeslot, error) {
	var timeslot models.Timeslot
	result := db.First(&timeslot, "time = ?", timeTS)
	if result.Error != nil {
		fmt.Println("Error fetching timeslot: ", result.Error)
	}
	return timeslot, result.Error

}