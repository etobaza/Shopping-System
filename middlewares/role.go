package middlewares

import (
	"shopping-system/config"
	"shopping-system/models"
)

// GetUserRole retrieves the user's role from the database using the username.
func GetUserRole(userID uint) (string, error) {
	var user models.User
	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return "", err
	}
	return user.UserType, nil
}

// HasRole HasRoleByID checks if the user has the specified role using the user ID.
func HasRole(userID uint, role string) (bool, error) {
	userRole, err := GetUserRole(userID)
	if err != nil {
		return false, err
	}
	return userRole == role, nil
}
