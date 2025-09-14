package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID    uint   `json:"user_id" gorm:"not null"`
	FullName  string `json:"full_name" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	House     string `json:"house" validate:"required"`
	Street    string `json:"street" validate:"required"`
	City      string `json:"city" validate:"required"`
	State     string `json:"state" validate:"required"`
	Pincode   string `json:"pincode" validate:"required"`
	IsDefault bool   `json:"is_default"`
}

type CreateAddressRequest struct {
	FullName string `json:"full_name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	House    string `json:"house" validate:"required"`
	Street   string `json:"street"`
	City     string `json:"city" vlaidate:"required"`
	State    string `json:"state" validate:"required"`
	Pincode  string `json:"pincode" validate:"required"`
}

type UpdateAddressRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	House    string `json:"house" binding:"required"`
	Street   string `json:"street" binding:"required"`
	City     string `json:"city" binding:"required"`
	State    string `json:"state" binding:"required"`
	Pincode  string `json:"pincode" binding:"required"`
}
