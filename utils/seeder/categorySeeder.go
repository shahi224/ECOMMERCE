package seeder

import (
	"log"

	"ECOMMERCE/database"
	"ECOMMERCE/utils/models"
)

// add categories
func SeedCategories() {
	categories := []models.Category{
		{Name: "Electronics", Description: "Electronic gadgets and devices", CreatedBy: 1},
		{Name: "Fashion", Description: "Clothing, footwear, and accessories", CreatedBy: 1},
		{Name: "Home & Kitchen", Description: "Home appliances and kitchenware", CreatedBy: 1},
		{Name: "Makeup &Skincare", Description: "Lipstick", CreatedBy: 1},
	}

	for _, category := range categories {
		var existing models.Category
		err := database.DB.Where("name = ?", category.Name).First(&existing).Error
		if err == nil {
			continue // already exists
		}
		if err := database.DB.Create(&category).Error; err != nil {
			log.Printf("failed to seed category %s: %v", category.Name, err)
		} else {
			log.Printf("Seeded category: %s", category.Name)
		}
	}
}
