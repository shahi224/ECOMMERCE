package migration

import (
	"log"

	"ECOMMERCE/database"
	"ECOMMERCE/utils/models"
)

func Migration() {
	db := database.DB

	if err := db.AutoMigrate(
		models.User{},
		models.UserProfile{},
		models.Category{},
		models.Brand{},
		models.Product{},
		models.Order{},
		models.OrderItem{},
		models.Cart{},
		// // models.Image{},
		models.Payment{},
		models.Address{},
		// models.Wallet{},
		models.Wishlist{},
	); err != nil {
		log.Fatal("failed to migrate")
		// panic("K")
	}
}
