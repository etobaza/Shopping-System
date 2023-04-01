package controllers

import (
	"encoding/json"
	"net/http"
	"shopping-system/middlewares"
	"shopping-system/models"
	"shopping-system/utils"

	"github.com/jinzhu/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var reqData struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		Email     string `json:"email"`
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Address   string `json:"address"`
		Phone     string `json:"phone"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		utils.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}
	// Check if username is available
	if available, err := middlewares.UsernameIsAvailable(reqData.Username, uc.DB); err != nil {
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
	newUser := middlewares.ManufactureUser(reqData.Username, hashedPassword, reqData.Email, reqData.FirstName, reqData.LastName, reqData.Address, reqData.Phone)

	if err := uc.DB.Create(&newUser).Error; err != nil {
		utils.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	utils.Respond(w, newUser, http.StatusCreated)
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var reqData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		utils.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}
	// Check if password is correct
	if correct, err := middlewares.PasswordIsCorrect(reqData.Username, reqData.Password, uc.DB); err != nil {
		utils.Error(w, "Failed to check password", http.StatusInternalServerError)
		return
	} else if !correct {
		utils.Error(w, "Incorrect password", http.StatusBadRequest)
		return
	}

	// Retrieve the user
	var user models.User
	if err := uc.DB.Where("username = ?", reqData.Username).First(&user).Error; err != nil {
		utils.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	utils.Respond(w, user, http.StatusOK)
}
