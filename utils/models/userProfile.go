package models

import (
	"time"

	"gorm.io/gorm"
)

type UserProfile struct {
	gorm.Model
	UserID uint       `gorm:"unique;not null"`
	Name   string     `gorm:"size:100" json:"name"`
	Email  string     `gorm:"size:100" json:"email"`
	Phone  string     `gorm:"size:20" json:"phone"`
	DOB    *time.Time `json:"dob,omitempty"`
	Gender string     `gorm:"size:20" json:"gender"`
}
type CreateUserProfileRequest struct {
	Name   string `json:"name" validate:"required"`
	Email  string `json:"email" validate:"required,email"`
	Phone  string `json:"phone"`
	DOB    string `json:"dob"`
	Gender string `json:"gender"`
}

type UpdateUserProfileRequest struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	DOB    string `json:"dob"`
	Gender string `json:"gender"`
}
