package controllers

import (
	"encoding/json"
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
	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		utils.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Check if username is available
	if available, err := middlewares.UsernameIsAvailable(newUser.Username, uc.DB); err != nil {
		utils.Error(w, "Failed to check username availability", http.StatusInternalServerError)
		return
	} else if !available {
		utils.Error(w, "Username is not available", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := middlewares.HashPassword(newUser.Password)
	if err != nil {
		utils.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	newUser.Password = hashedPassword
	newUser.Balance = 0
	newUser.UserType = "customer"

	if err := uc.DB.Create(&newUser).Error; err != nil {
		utils.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	utils.Respond(w, newUser, http.StatusCreated)
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Check if password is correct
	if correct, err := middlewares.PasswordIsCorrect(user.Username, user.Password, uc.DB); err != nil {
		utils.Error(w, "Failed to check password", http.StatusInternalServerError)
		return
	} else if !correct {
		utils.Error(w, "Incorrect password", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "http://localhost:8080/login", http.StatusSeeOther)
	utils.Respond(w, user, http.StatusOK)
}

func (uc *UserController) GetRegister(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./views/pages/registration.html")
}

func (uc *UserController) GetLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./views/pages/login.html")
}
