package middlewares

import (
	"github.com/jinzhu/gorm"
	"shopping-system/models"
)

func UsernameIsAvailable(username string, db *gorm.DB) (bool, error) {
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return true, nil
		} else {
			return false, err
		}
	}
	return false, nil
}

func PasswordIsCorrect(username, password string, db *gorm.DB) (bool, error) {
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		} else {
			return false, err
		}
	}

	if CheckPasswordHash(password, user.Password) {
		return true, nil
	}

	return false, nil
}
