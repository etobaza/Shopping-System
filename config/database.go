package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"shopping-system/models"
)

const DBDialect = "postgres"
const DBHost = "localhost"
const DBPort = "5432"
const DBUser = "postgres"
const DBName = "ecommerce"
const DBPassword = "darkside"

var DbUri = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", DBHost, DBPort, DBUser, DBName, DBPassword)
var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(DBDialect, DbUri)
	if err != nil {
		panic("Failed to connect to database")
	}
	DB = db
	DB.AutoMigrate(&models.User{}, &models.Item{})
}
