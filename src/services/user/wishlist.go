package services

import (
	repository "ECOMMERCE/src/repository/user"
	"ECOMMERCE/utils/models"
)

// add to wishlist
func AddToWishlist(userID uint, req models.AddToWishlistRequest) error {
	return repository.AddToWishlist(userID, req)
}

// get wishlist
func GetWishlist(userID uint) ([]models.Wishlist, error) {
	return repository.GetWishlist(userID)
}

// remove from wishlist
func RemoveFromWishlist(userID, productID uint) error {
	return repository.RemoveFromWishlist(userID, productID)
}

// clear wishlist
func ClearWishlist(userID uint) error {
	return repository.ClearWishlist(userID)
}
