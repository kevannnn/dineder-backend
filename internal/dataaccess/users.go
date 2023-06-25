package dataaccess

import (
	"errors"
	"fmt"

	"github.com/kevannnn/dineder-backend/internal/models"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, newUser models.User) error {
	result := db.Table("dineder_user").Create(&newUser)
	if result.Error != nil {
		fmt.Println("Error creating user:", result.Error)
	}
	return result.Error
}

func ReadUserID(db *gorm.DB, email string, password string) (models.User, error) {
	var user models.User
	result := db.Where(&models.User{Email: email, Password: password}).First(&user)
	if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return user, errors.New("user not found")
        }
        return user, result.Error
    }
	return user, nil
}