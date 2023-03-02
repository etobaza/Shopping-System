package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"shopping-system/config"
	"shopping-system/controllers"
	"shopping-system/models"
	"shopping-system/routes"
)

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(config.DBDialect, config.DBURI)
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}
	defer db.Close()
	db.AutoMigrate(&models.User{})
	fmt.Println("Connected to database")

	r := mux.NewRouter()
	uc := &controllers.UserController{DB: db}

	routes.SetupRoutes(r, uc)

	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", r)
}
