package middlewares

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
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
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
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

func ValidateTokenFromRequest(r *http.Request) (bool, error) {
	// Extract token from the request
	tokenString := utils.ExtractToken(r)
	if tokenString == "" {
		fmt.Println("Token is missing")
		return false, fmt.Errorf("missing token")
	}

	// Parse the token to get the claims
	claims, err := utils.ParseToken(tokenString)
	if err != nil {
		fmt.Println("Failed to parse token:", err)
		return false, fmt.Errorf("failed to parse token")
	}

	// Check if the token is valid
	if !utils.IsTokenValid(claims) {
		fmt.Println("Token is invalid")
		return false, fmt.Errorf("token has expired")
	}

	fmt.Println("Token is valid")
	return true, nil
}
