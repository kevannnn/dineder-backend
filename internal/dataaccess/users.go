package dataaccess

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/kevannnn/dineder-backend/internal/models"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

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

func CreateUser(db *gorm.DB, newUser models.User) (models.User, error) {
	result := db.Table("dineder_user").Create(&newUser)
	if result.Error != nil {
		fmt.Println("Error creating user:", result.Error)
		return newUser, result.Error
	}
	return newUser, nil
}

func GetUserNameFromID(db *gorm.DB, userID uuid.UUID) (string, error) {
	var user models.User
	result := db.First(&user, "id = ?", userID)
	if result.Error != nil {
		fmt.Println("Error fetching user: ", result.Error)
	}
	return user.Name, result.Error
}

func UpdateUserGender(db *gorm.DB, userID string, userGender int) (models.User, error) {
	var user models.User
	userResult := db.Table("dineder_user").Where("id = ?", userID).First(&user)
	if userResult.Error != nil {
		fmt.Println("Error fetching user: ", userResult.Error)
		return user, userResult.Error
	}
	
	user.Gender = userGender
	updateResult := db.Save(&user)
	if updateResult.Error != nil {
		fmt.Println("Failed to update gender: ", updateResult.Error)
		return user, updateResult.Error
	}
	return user, nil
}

func UpdateUserPrefGender(db *gorm.DB, userID string, userPrefGender int) (models.User, error) {
	var user models.User
	userResult := db.Table("dineder_user").Where("id = ?", userID).First(&user)
	if userResult.Error != nil {
		fmt.Println("Error fetching user: ", userResult.Error)
		return user, userResult.Error
	}
	
	user.PrefGender = userPrefGender
	updateResult := db.Save(&user)
	if updateResult.Error != nil {
		fmt.Println("Failed to update preferred gender: ", updateResult.Error)
		return user, updateResult.Error
	}
	return user, nil
}

func UpdateUserFaculty(db *gorm.DB, userID string, userFaculty int) (models.User, error) {
	var user models.User
	userResult := db.Table("dineder_user").Where("id = ?", userID).First(&user)
	if userResult.Error != nil {
		fmt.Println("Error fetching user: ", userResult.Error)
		return user, userResult.Error
	}
	
	user.Faculty = userFaculty
	updateResult := db.Save(&user)
	if updateResult.Error != nil {
		fmt.Println("Failed to update faculty: ", updateResult.Error)
		return user, updateResult.Error
	}
	return user, nil
}

func UpdateUserField(db *gorm.DB, userID string, fieldName string, fieldValue interface{}) (models.User, error) {
	var user models.User
	userResult := db.Table("dineder_user").Where("id = ?", userID).First(&user)
	if userResult.Error != nil {
		fmt.Println("Error fetching user: ", userResult.Error)
		return user, userResult.Error
	}

	reflect.ValueOf(&user).Elem().FieldByName(fieldName).Set(reflect.ValueOf(fieldValue))

	updateResult := db.Save(&user)
	if updateResult.Error != nil {
		fmt.Printf("Failed to update %s: %s\n", fieldName, updateResult.Error)
		return user, updateResult.Error
	}
	return user, nil
}