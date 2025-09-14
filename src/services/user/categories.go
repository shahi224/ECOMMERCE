package services

import (
	repository "ECOMMERCE/src/repository/user"
	"ECOMMERCE/utils/models"
)

// get all categories
func GetAllCategories() ([]models.Category, error) {
	return repository.GetAllCategories()
}

// get product by category ID
func GetProductByCategoryID(categoryID uint) ([]models.Product, error) {
	return repository.GetProductByCategoryID(categoryID)
}
