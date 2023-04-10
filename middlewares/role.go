package middlewares

import (
	"shopping-system/config"
	"shopping-system/models"
)

func GetUserRole(userID uint) (string, error) {
	var user models.User
	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return "", err
	}
	return user.UserType, nil
}

func HasRole(userID uint, role string) (bool, error) {
	userRole, err := GetUserRole(userID)
	if err != nil {
		return false, err
	}
	return userRole == role, nil
}
