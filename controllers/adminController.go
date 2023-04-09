package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"path/filepath"
	"shopping-system/config"
	"shopping-system/middlewares"
	"shopping-system/models"
	"shopping-system/utils"
)

func AdminPanel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	tokenString := utils.ExtractToken(r)
	claims, _ := utils.ParseToken(tokenString)
	userID := uint(claims["user_id"].(float64))

	hasRole, err := middlewares.HasRole(userID, "admin")
	if err != nil {
		utils.Error(w, "Error checking user role", http.StatusInternalServerError)
		return
	}
	if !hasRole {
		utils.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	buildDir := "./client/build/"
	http.ServeFile(w, r, filepath.Join(buildDir, "index.html"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	tokenString := utils.ExtractToken(r)
	claims, _ := utils.ParseToken(tokenString)
	userID := uint(claims["user_id"].(float64))

	hasRole, err := middlewares.HasRole(userID, "admin")
	if err != nil {
		fmt.Println("Error checking user role:", err)
		utils.Error(w, "Error checking user role", http.StatusInternalServerError)
		return
	}
	if !hasRole {
		fmt.Println("Not allowed", err)
		utils.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	if err != nil {
		utils.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if errDel := config.DB.Delete(&models.User{}, id).Error; errDel != nil {
		fmt.Println("Failed to delete user:", errDel)
		utils.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	utils.Respond(w, map[string]string{"message": "User deleted successfully"}, http.StatusOK)
}
