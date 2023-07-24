package models

import (
	"time"
	"github.com/satori/go.uuid"
)

type User struct {
	ID       	uuid.UUID 	`json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name 	 	string 		`json:"name"`
	Email    	string 		`json:"email"`
	DOB 		time.Time 	`json:"dob"`
	Password 	string 		`json:"password"`
	Gender 		int 		`json:"gender"`
	PrefGender 	int 		`json:"pref_gender"`
	Faculty 	int 		`json:"faculty"`
	MealPref 	int 		`json:"meal_pref"`
}

func (User) TableName() string {
	return "dineder_user"
}
