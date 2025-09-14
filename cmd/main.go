package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"ECOMMERCE/config"
	"ECOMMERCE/database"
	"ECOMMERCE/migration"
	"ECOMMERCE/routes"
	"ECOMMERCE/seeder"
)

func main() {
	r := gin.Default()

	// ✅ Load HTML templates from templates folder
	r.LoadHTMLGlob("templates/*.html") // Make sure you have templates folder with .html files

	// Load config
	cfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading the config file")
	}

	// Database connection
	db, err := database.ConnectDB(cfig)
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	// Migration & Seeder
	migration.Migration()
	seeder.GroupSeeder()

	// Register routes
	routes.RegisterRoutes(r, db)

	// ✅ Add a test route for HTML page
	r.GET("/dashboard", func(c *gin.Context) {
		c.HTML(200, "dashboard.html", gin.H{
			"title":   "Admin Dashboard",
			"message": "Welcome to your dashboard!",
		})
	})

	log.Printf("Starting server at %s\n", cfig.ServerAddress)
	if err := r.Run(cfig.ServerAddress); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
