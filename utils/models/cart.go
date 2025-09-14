package models

import "gorm.io/gorm"

type AddToCartRequest struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required,min=1"`
}

type UpdateCartRequest struct {
	Quantity int `json:"quantity" validate:"required,min=1"`
}
type Cart struct {
	gorm.Model
	UserID     uint    `gorm:"not null;index" json:"user_id"`
	ProductID  uint    `gorm:"not null;index" json:"product_id"`
	Quantity   float64 `gorm:"not null" json:"quantity"`
	TotalPrice float64 `gorm:"not null" json:"total_price"`

	User    User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Product Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
}

// If it's empty, don't include it in the response (omitempty = "only show if it's filled").

// The tag (gorm:"constraint:...") tells the database what to do if something changes.
