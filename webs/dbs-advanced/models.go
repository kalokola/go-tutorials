package main

import (
	"time"

	"gorm.io/gorm"
)

// User model
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	FirstName string         `gorm:"not null" json:"first_name"`
	LastName  string         `gorm:"not null" json:"last_name"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"not null" json:"-"` // "-" means don't include in JSON
	APIKey    string         `gorm:"uniqueIndex;not null" json:"api_key"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Album model with optional user relationship
type Album struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"not null" json:"title"`
	Artist    string         `gorm:"not null" json:"artist"`
	Price     float64        `gorm:"not null" json:"price"`
	UserID    *uint          `gorm:"index" json:"user_id,omitempty"` // Pointer = nullable
	User      *User          `gorm:"foreignKey:UserID" json:"user,omitempty"` // Pointer = optional
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Response structs
type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}