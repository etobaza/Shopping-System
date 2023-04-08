package models

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
	Category    string `gorm:"not null" json:"category"`
	Description string `gorm:"not null" json:"description"`
	Price       uint   `gorm:"not null" json:"price"`
	Quantity    uint   `gorm:"not null" json:"quantity"`
	Owner       User   `gorm:"foreignkey:OwnerID" json:"owner"`
	OwnerID     uint   `gorm:"not null" json:"ownerid"`
}
