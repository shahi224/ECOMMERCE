package repository

import (
	"errors"

	"ECOMMERCE/database"
	"ECOMMERCE/utils/models"
)

func AddToWishlist(userID uint, req models.AddToWishlistRequest) error {
	var existingProduct models.Product
	if err := database.DB.First(&existingProduct, req.ProductID).Error; err != nil {
		return errors.New("product not found")
	}

	var existing models.Wishlist
	if err := database.DB.Where("user_id = ? AND product_id = ?", userID, req.ProductID).First(&existing).Error; err == nil {
		return nil // Already exists
	}

	wishlist := models.Wishlist{
		UserID:    userID,
		ProductID: req.ProductID,
	}
	return database.DB.Create(&wishlist).Error
}

func GetWishlist(userID uint) ([]models.Wishlist, error) {
	var items []models.Wishlist
	err := database.DB.Preload("Product").Where("user_id = ?", userID).Find(&items).Error
	return items, err
}

func RemoveFromWishlist(userID, productID uint) error {
	var wishlist models.Wishlist
	return database.DB.Where("user_id=? AND product_id=?", userID, productID).Delete(&wishlist).Error
}

func ClearWishlist(userID uint) error {
	return database.DB.Where("user_id = ?", userID).Delete(&models.Wishlist{}).Error
}
