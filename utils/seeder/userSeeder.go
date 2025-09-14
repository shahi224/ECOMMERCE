package seeder

import (
	"log"

	"ECOMMERCE/database"
	"ECOMMERCE/utils/models"
)

// add users
func SeedUsers() {
	users := []models.User{
		{Name: "John Doe", Email: "john@example.com", Password: "password123", Phone: "9999900001", Role: "user", Status: "active", IsAdmin: false},
		{Name: "Alice Smith", Email: "alice@example.com", Password: "password123", Phone: "9999900002", Role: "user", Status: "active", IsAdmin: false},
		{Name: "Bob Johnson", Email: "bob@example.com", Password: "password123", Phone: "9999900003", Role: "admin", Status: "active", IsAdmin: false},
		{Name: "Charlie Brown", Email: "charlie@example.com", Password: "password123", Phone: "9999900004", Role: "user", Status: "inactive", IsAdmin: false},
		{Name: "David Wilson", Email: "david@example.com", Password: "password123", Phone: "9999900005", Role: "user", Status: "active", IsAdmin: false},
		{Name: "Eve Davis", Email: "eve@example.com", Password: "password123", Phone: "9999900006", Role: "admin", Status: "active", IsAdmin: false},
		{Name: "Frank Harris", Email: "frank@example.com", Password: "password123", Phone: "9999900007", Role: "user", Status: "active", IsAdmin: false},
		{Name: "Grace Miller", Email: "grace@example.com", Password: "password123", Phone: "9999900008", Role: "user", Status: "inactive", IsAdmin: false},
		{Name: "Hannah Clark", Email: "hannah@example.com", Password: "password123", Phone: "9999900009", Role: "user", Status: "active", IsAdmin: false},
		{Name: "Ivan Turner", Email: "ivan@example.com", Password: "password123", Phone: "9999900010", Role: "admin", Status: "active", IsAdmin: false},
	}

	for _, user := range users {
		var existing models.User
		err := database.DB.Where("email = ?", user.Email).First(&existing).Error
		if err == nil {
			continue
		}
		if err := database.DB.Create(&user).Error; err != nil {
			log.Printf("Failed to seed user %s: %v", user.Email, err)
		} else {
			log.Printf("Seeded user: %s", user.Email)
		}
	}
}
