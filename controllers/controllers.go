package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
	hashedPassword, err := utils.HashPassword(reqData.Password)
	if err != nil {
		utils.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Create user
	newUser := utils.ManufactureUser(reqData.Username, hashedPassword, reqData.Email, reqData.FirstName, reqData.LastName, reqData.Address, reqData.Phone, reqData.UserType)

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
	log.Printf("User type: %v\n", user.UserType)

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
		Token    string `json:"token"`
		ID       uint   `json:"id"`
		UserType string `json:"usertype"`
	}{
		User:     user,
		Token:    token,
		ID:       user.ID,
		UserType: user.UserType,
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

func Shop(w http.ResponseWriter, r *http.Request) {
	middlewares.ServePage(w, r)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		utils.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}
	utils.Respond(w, users, http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Get the user's information from the database
	var user models.User
	if errInf := config.DB.Where("id = ?", id).First(&user).Error; errInf != nil {
		utils.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	utils.Respond(w, user, http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		utils.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	var reqData struct {
		Balance float64 `json:"balance"`
	}
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		utils.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	user.Balance = uint(reqData.Balance)

	if err := config.DB.Save(&user).Error; err != nil {
		utils.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	utils.Respond(w, user, http.StatusOK)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		utils.Error(w, "Failed to retrieve item", http.StatusInternalServerError)
		return
	}

	var reqData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		utils.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	// validate the fields to be updated
	allowedFields := map[string]bool{"name": true, "price": true, "quantity": true}
	for field := range reqData {
		if !allowedFields[field] {
			utils.Error(w, "Invalid field: "+field, http.StatusBadRequest)
			return
		}
	}

	// update the fields
	for field, value := range reqData {
		switch field {
		case "name":
			item.Name = value.(string)
		case "price":
			item.Price = uint(value.(float64))
		case "quantity":
			newQuantity := uint(value.(float64))
			if newQuantity < 0 {
				utils.Error(w, "Invalid quantity: "+string(newQuantity), http.StatusBadRequest)
				return
			}
			item.Quantity = newQuantity
		}
	}

	if err := config.DB.Save(&item).Error; err != nil {
		utils.Error(w, "Failed to update item", http.StatusInternalServerError)
		return
	}

	utils.Respond(w, item, http.StatusOK)
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	var items []models.Item
	if err := config.DB.Find(&items).Error; err != nil {
		utils.Error(w, "Failed to retrieve items", http.StatusInternalServerError)
		return
	}
	utils.Respond(w, items, http.StatusOK)
}

func GetItemsSeller(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")

	var items []models.Item
	if err := config.DB.Where("user_id = ?", userID).Find(&items).Error; err != nil {
		utils.Error(w, "Failed to retrieve items", http.StatusInternalServerError)
		return
	}
	utils.Respond(w, items, http.StatusOK)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var reqData models.Item
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		log.Printf("Error decoding request data: %v\n", err)
		utils.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	tokenString := utils.ExtractToken(r)
	if tokenString == "" {
		utils.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	claims, err := utils.ParseToken(tokenString)
	if err != nil {
		log.Printf("Error parsing token: %v\n", err)
		utils.Error(w, "Failed to parse token", http.StatusInternalServerError)
		return
	}

	userID := uint(claims["user_id"].(float64))

	// Set the user ID of the item to the current user's ID
	reqData.UserID = userID

	if err := config.DB.Create(&reqData).Error; err != nil {
		log.Printf("Error creating item: %v\n", err)
		utils.Error(w, "Failed to create item", http.StatusInternalServerError)
		return
	}

	utils.Respond(w, reqData, http.StatusCreated)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Get the item's information from the database
	var item models.Item
	if errInf := config.DB.Where("id = ?", id).First(&item).Error; errInf != nil {
		utils.Error(w, "Failed to retrieve item", http.StatusInternalServerError)
		return
	}

	tokenString := utils.ExtractToken(r)
	if tokenString == "" {
		utils.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	claims, err := utils.ParseToken(tokenString)
	if err != nil {
		log.Printf("Error parsing token: %v\n", err)
		utils.Error(w, "Failed to parse token", http.StatusInternalServerError)
		return
	}

	userID := uint(claims["user_id"].(float64))

	if item.UserID != userID {
		utils.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := config.DB.Delete(&item).Error; err != nil {
		log.Printf("Error deleting item: %v\n", err)
		utils.Error(w, "Failed to delete item", http.StatusInternalServerError)
		return
	}

	utils.Respond(w, map[string]string{"message": "Item deleted"}, http.StatusOK)
}
