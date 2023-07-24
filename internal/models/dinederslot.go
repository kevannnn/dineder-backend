package models

import (
	"github.com/satori/go.uuid"
)

type DinederSlot struct {
	ID 			uint  		`json:"id" gorm:"type:primaryKey"`
	UserID		uuid.UUID 	`json:"user_id" gorm:"type:uuid;foreignKey"`
	TimeslotID	uint 		`json:"timeslot_id" gorm:"foreignKey"`
}

func (DinederSlot) TableName() string {
	return "dineder_slots"
}