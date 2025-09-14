package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `gorm:"size:150;not null" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock       int     `gorm:"not null;default:0" json:"stock"`

	CategoryID uint     `json:"category_id"`
	BrandID    uint     `json:"brand_id"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Brand      Brand    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

// type CreateProductRequest struct {
// 	Name        string  `json:"name" binding:"required"`
// 	Description string  `json:"description"`
// 	Price       float64 `json:"price" binding:"required,gt=0"`
// 	Stock       int     `json:"stock" binding:"required,gte=0"`
// 	CategoryID  uint    `json:"Category_id" binding:"required"`
// }

// type UpdateProductRequest struct {
// 	ID          uint    `json:"id" binding:"required"`
// 	Name        string  `json:"name"`
// 	Description string  `json:"description"`
// 	Price       float64 `json:"price"`
// 	Stock       int     `json:"stock"`
// 	CategoryID  uint    `json:"category_id"`
// }
