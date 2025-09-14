package models

import "gorm.io/gorm"

type Wishlist struct {
	gorm.Model
	UserID    uint    `json:"user_id" gorm:"not null"`
	ProductID uint    `json:"product_id" gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID"`
}

type AddToWishlistRequest struct {
	ProductID uint `json:"product_id" validate:"required"`
}
