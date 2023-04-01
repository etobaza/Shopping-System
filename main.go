package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"net/http"
	"shopping-system/config"
	"shopping-system/controllers"
	"shopping-system/middlewares"
	"shopping-system/models"
	"shopping-system/routes"
)

var db *gorm.DB

func main() {
	db, err := gorm.Open(config.DBDialect, config.DBURI)
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
	defer db.Close()
	db.AutoMigrate(&models.User{})
	r := mux.NewRouter()
	r.Use(middlewares.CorsMiddleware)
	uc := &controllers.UserController{DB: db}
	routes.SetupRoutes(r, uc)
	middlewares.ServeMessage()
	http.ListenAndServe(":8080", r)
}
