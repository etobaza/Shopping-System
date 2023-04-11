package middlewares

import (
	"github.com/jinzhu/gorm"
	"log"
	"shopping-system/config"
	"shopping-system/models"
	"shopping-system/utils"
)

func UsernameIsAvailable(username string) (bool, error) {
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return true, nil
		} else {
			return false, err
		}
	}
	return false, nil
}

func PasswordIsCorrect(username, password string) (bool, error) {
	var user models.User
	print("username: " + username + "\n")
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			log.Printf("[PIC] Error retrieving record: %v\n", err)
			return false, nil
		} else {
			log.Printf("[PIC] Some other issue: %v\n", err)
			return false, err
		}
	}

	if utils.CheckPasswordHash(password, user.Password) {
		return true, nil
	}

	return false, nil
}
