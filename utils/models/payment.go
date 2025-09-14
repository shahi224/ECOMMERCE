package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	OrderID       uint    `gorm:"uniqueIndex;not null" json:"order_id"`
	Amount        float64 `gorm:"type:decimal(10,2);not null" json:"amount"`
	PaymentMethod string  `gorm:"size:50;not null" json:"payment_method"` // e.g., Razorpay, card, COD
	Status        string  `gorm:"size:50;not null" json:"status"`         // completed, pending, failed
	TransactionID string  `gorm:"size:100;uniqueIndex" json:"transaction_id"`
}

//api request model

type PaymentRequest struct {
	OrderID     uint   `json:"order_id" validation:"required"`
	PaymentMode string `json:"payment_mode" validate:"required,oneof=COD RAZORPAY STRIPE"`
}
