package models

import (
	"time"
)

type Timeslot struct {
	ID 			uint  		`json:"id" gorm:"type:primaryKey"`
	Time 		time.Time 	`json:"time" gorm:"type:timestamp"`
}

func (Timeslot) TableName() string {
	return "timeslots"
}