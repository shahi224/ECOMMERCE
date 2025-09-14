package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"size:100;not null" json:"email"`
	Password string `gorm:"size:255;not null" json:"-"`
	Phone    string `json:"phone"`
	OTP      string `json:"otp"`
	Role     string `gorm:"size:20;default:'user'" json:"role"`
	Status   string `gorm:"size:20;default:'active'" json:"status"`
	// UserProfile UserProfile `gorm:"foreignKey:UserID" json:"user_profile,omitempty"`
	IsAdmin bool    `json:"is_admin"`
	Orders  []Order `json:"orders,omitempty"`
}

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	Phone    string `json:"phone" validate:"required"`
}

type LoginRequest struct {
	Phone string `json:"phone" validate:"required"`
	// Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ChangePassword struct {
	Oldpassword string `json:"old_password"`
	Password    string `json:"password"`
	Repassword  string `json:"re_password"`
}

type ForgotPasswordSend struct {
	Phone string `json:"phone"`
}
type ForgotVerify struct {
	Phone       string `json:"phone" binding:"required" validate:"required"`
	Otp         string `json:"otp" binding:"required"`
	NewPassword string `json:"newpassword" binding:"required" validate:"min=6,max=20"`
}
