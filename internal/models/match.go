package models

import (
	"github.com/satori/go.uuid"
)

type Match struct {
	ID			uint 		`json:"id" gorm:"type:primaryKey"`
	User1ID		uuid.UUID 	`json:"user1_id" gorm:"type:uuid;foreignKey"`
	User2ID		uuid.UUID 	`json:"user2_id" gorm:"type:uuid;foreignKey"`
	TimeslotID	uint		`json:"timeslot_id" gorm:"type:foreignKey"`
}

func (Match) TableName() string {
	return "matches"
}