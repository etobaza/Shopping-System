package models

import (
	"time"
)

type Item struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Category    string    `gorm:"not null" json:"category"`
	Description string    `gorm:"not null" json:"description"`
	Price       uint      `gorm:"not null" json:"price"`
	Quantity    uint      `gorm:"not null" json:"quantity"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at"`
	UserID      uint      `gorm:"not null;index" json:"user_id"` // Add the foreign key
}
