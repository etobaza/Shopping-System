package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"uniqueIndex;not null" json:"username"`
	Password  string `gorm:"not null" json:"password"`
	Email     string `gorm:"uniqueIndex;not null" json:"email"`
	FirstName string `gorm:"not null" json:"firstname"`
	LastName  string `gorm:"not null" json:"lastname"`
	Balance   uint   `gorm:"not null" json:"balance"`
	Address   string `gorm:"not null" json:"address"`
	Phone     string `gorm:"not null" json:"phone"`
	UserType  string `gorm:"not null" json:"user_type"`
}
