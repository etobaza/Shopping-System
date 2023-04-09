package controllers

import (
	"encoding/json"
	"log"
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

	if addErr := config.DB.Create(&newUser).Error; addErr != nil {
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
		log.Printf("Error decoding request data: %v\n", err)
		utils.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// Check if password is correct
	if correct, err := middlewares.PasswordIsCorrect(reqData.Username, reqData.Password); err != nil {
		log.Printf("Error checking password: %v\n", err)
		utils.Error(w, "Failed to check password", http.StatusInternalServerError)
		return
	} else if !correct {
		log.Println("Incorrect password")
		utils.Error(w, "Incorrect password", http.StatusBadRequest)
		return
	}

	// Retrieve the user
	var user models.User
	if err := config.DB.Where("username = ?", reqData.Username).First(&user).Error; err != nil {
		log.Printf("Error retrieving user: %v\n", err)
		utils.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	token, err := utils.GenerateToken(&user)
	if err != nil {
		log.Printf("Error generating token: %v\n", err)
		utils.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	cookie := utils.CreateCookie(token)
	http.SetCookie(w, cookie)

	response := struct {
		models.User
		Token string `json:"token"`
	}{
		User:  user,
		Token: token,
	}
	utils.Respond(w, response, http.StatusOK)

	log.Println("Login successful")
}

func Logout(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	isValid, err := middlewares.ValidateTokenFromRequest(r)
	if !isValid {
		utils.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	buildDir := "./client/build/"
	http.ServeFile(w, r, filepath.Join(buildDir, "index.html"))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		utils.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	utils.Respond(w, users, http.StatusOK)
}
