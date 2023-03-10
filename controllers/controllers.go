package controllers

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"shopping-system/middlewares"
	"shopping-system/models"
	"shopping-system/utils"
)

type UserController struct {
	DB *gorm.DB
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")
	firstName := r.FormValue("firstname")
	lastName := r.FormValue("lastname")
	address := r.FormValue("address")
	phone := r.FormValue("phone")

	// Check if username is available
	if available, err := middlewares.UsernameIsAvailable(username, uc.DB); err != nil {
		utils.Error(w, "Failed to check username availability", http.StatusInternalServerError)
		return
	} else if !available {
		utils.Error(w, "Username is not available", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := middlewares.HashPassword(password)
	if err != nil {
		utils.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Create user
	newUser := middlewares.ManufactureUser(username, hashedPassword, email, firstName, lastName, address, phone)

	if err := uc.DB.Create(&newUser).Error; err != nil {
		utils.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	utils.Respond(w, newUser, http.StatusCreated)
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Check if password is correct
	if correct, err := middlewares.PasswordIsCorrect(username, password, uc.DB); err != nil {
		utils.Error(w, "Failed to check password", http.StatusInternalServerError)
		return
	} else if !correct {
		utils.Error(w, "Incorrect password", http.StatusBadRequest)
		return
	}

	// Retrieve the user
	var user models.User
	if err := uc.DB.Where("username = ?", username).First(&user).Error; err != nil {
		utils.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	utils.Respond(w, user, http.StatusOK)
}

func (uc *UserController) GetRegister(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./views/pages/registration.html")
}

func (uc *UserController) GetLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./views/pages/login.html")
}
