package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `gorm:"size:100;unique;not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`

	// Relationships
	Brands   []Brand   `gorm:"foreignKey:CategoryID" json:"brands,omitempty"`
	Products []Product `gorm:"foreignKey:CategoryID" json:"products,omitempty"`

	CreatedBy uint `json:"created_by"`
}
type Brand struct {
	gorm.Model
	Name       string   `gorm:"size:100;not null;unique" json:"name"`
	CategoryID uint     `json:"category_id"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

// OnDelete:SET NULL :	If a category is deleted, don’t delete the product—just remove the category link. (Set category_id to NULL)
