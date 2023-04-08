package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"shopping-system/config"
	"shopping-system/middlewares"
	"shopping-system/models"
	"shopping-system/utils"
	"time"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var reqData struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		Email     string `json:"email"`
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Address   string `json:"address"`
		Phone     string `json:"phone"`
		UserType  string `json:"usertype"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		utils.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}
	// Check if username is available
	if available, err := middlewares.UsernameIsAvailable(reqData.Username); err != nil {
		utils.Error(w, "Failed to check username availability", http.StatusInternalServerError)
		return
	} else if !available {
		utils.Error(w, "Username is not available", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := middlewares.HashPassword(reqData.Password)
	if err != nil {
		utils.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Create user
	newUser := middlewares.ManufactureUser(reqData.Username, hashedPassword, reqData.Email, reqData.FirstName, reqData.LastName, reqData.Address, reqData.Phone, reqData.UserType)

	if err := config.DB.Create(&newUser).Error; err != nil {
		utils.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	utils.Respond(w, newUser, http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var reqData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		utils.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}
	// Check if password is correct
	if correct, err := middlewares.PasswordIsCorrect(reqData.Username, reqData.Password); err != nil {
		fmt.Println(err)
		utils.Error(w, "Failed to check password", http.StatusInternalServerError)
		return
	} else if !correct {
		utils.Error(w, "Incorrect password", http.StatusBadRequest)
		return
	}

	// Retrieve the user
	var user models.User
	if err := config.DB.Where("username = ?", reqData.Username).First(&user).Error; err != nil {
		utils.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	// Generate a token for the user
	token, err := utils.GenerateToken(&user)
	if err != nil {
		utils.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Create a cookie with the JWT token
	cookie := utils.CreateCookie(token)

	// Set the cookie in the response
	http.SetCookie(w, cookie)

	// Respond with the user and token
	response := struct {
		models.User
		Token string `json:"token"`
	}{
		User:  user,
		Token: token,
	}
	utils.Respond(w, response, http.StatusOK)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Delete the HttpOnly cookie that contains the JWT token by setting the cookie's expiration time to a past date.
	cookie := &http.Cookie{
		Name:     "jwtToken",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	}
	http.SetCookie(w, cookie)

	// Respond with a success message
	utils.Respond(w, map[string]string{"message": "Logout successful"}, http.StatusOK)
}

func Home(w http.ResponseWriter, r *http.Request) {
	// Validate the token
	isValid, err := middlewares.ValidateTokenFromRequest(r)
	if !isValid {
		utils.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// If the token is valid, serve the file
	buildDir := "./client/build/"
	http.ServeFile(w, r, filepath.Join(buildDir, "index.html"))
}
