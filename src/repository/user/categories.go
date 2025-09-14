package repository

import (
	"ECOMMERCE/database"
	"ECOMMERCE/utils/models"
)

// get all categories
func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := database.DB.Find(&categories).Error
	return categories, err
}

// get produst by category ID
func GetProductByCategoryID(categoryID uint) ([]models.Product, error) {
	var products []models.Product
	err := database.DB.Where("category_id = ?", categoryID).Find(&products).Error
	return products, err
}
