package seeder

import (
	"log"

	"ECOMMERCE/database"
	"ECOMMERCE/utils/models"
)

// add products
func SeedProducts() {
	products := []models.Product{
		{
			Name:        "iPhone 14 Pro Max",
			Description: "Latest Apple smartphone with A16 chip",
			Price:       139999.00,
			Stock:       25,
			CategoryID:  1, // Electronics
			BrandID:     1, // Apple
		},
		{
			Name:        "Samsung Galaxy Z Fold 4",
			Description: "Foldable phone with AMOLED display",
			Price:       154999.99,
			Stock:       15,
			CategoryID:  1,
			BrandID:     2,
		},
		{
			Name:        "Nike Air Max",
			Description: "Comfortable running shoes",
			Price:       9999.00,
			Stock:       50,
			CategoryID:  3, // Fashion
			BrandID:     3, // Nike
		},
		{
			Name:        "Adidas Shoes",
			Description: "Comforatable running shoes",
			Price:       6999.00,
			Stock:       50,
			CategoryID:  3,
			BrandID:     4,
		},
		{
			Name:        "Philips Air Fryer",
			Description: "Oil-free cooking air fryer",
			Price:       7999.99,
			Stock:       40,
			CategoryID:  4, // Home & Kitchen
			BrandID:     5, // Philips
		},
		{
			Name:        "Dazler Lipstick",
			Description: "Long Lasting Lipstick",
			Price:       300.00,
			Stock:       1,
			CategoryID:  5,
			BrandID:     6,
		},
	}

	for _, product := range products {
		var existing models.Product
		err := database.DB.Where("name = ?", product.Name).First(&existing).Error
		if err == nil {
			continue
		}
		if err := database.DB.Create(&product).Error; err != nil {
			log.Printf("Failed to seed product %s: %v", product.Name, err)
		} else {
			log.Printf("Seeded product: %s", product.Name)
		}
	}
}
