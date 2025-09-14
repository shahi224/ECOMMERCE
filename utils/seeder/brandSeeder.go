package seeder

import (
	"log"

	"ECOMMERCE/database"
	"ECOMMERCE/utils/models"
)

// add brands
func SeedBrands() {
	brands := []models.Brand{
		{Name: "Apple", CategoryID: 1},   // Electronics
		{Name: "Samsung", CategoryID: 1}, // Electronics
		{Name: "Nike", CategoryID: 3},    // Fashion
		{Name: "Adidas", CategoryID: 3},  // Fashion
		{Name: "Philips", CategoryID: 2}, // Home & Kitchen
		{Name: "Dazler", CategoryID: 4},
	}

	for _, brand := range brands {
		var existing models.Brand
		err := database.DB.Where("name = ?", brand.Name).First(&existing).Error
		if err == nil {
			continue
		}
		if err := database.DB.Create(&brand).Error; err != nil {
			log.Printf("Failed to seed brand %s: %v", brand.Name, err)
		} else {
			log.Printf("Seeded brand: %s", brand.Name)
		}
	}
}
