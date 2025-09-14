package services

import (
	repository "ECOMMERCE/src/repository/user"
	"ECOMMERCE/utils/models"
)

// add to cart
func AddToCart(userID uint, req models.AddToCartRequest) error {
	return repository.AddToCart(userID, req)
}

// get all cart items
func GetAllCartProducts(userID uint) ([]models.Cart, error) {
	return repository.GetAllCartProducts(userID)
}

// update cart items
func UpdateCartItem(userID, cartID uint, req models.UpdateCartRequest) error {
	return repository.UpdateCartItem(userID, cartID, req)
}

// remove items from cart
func RemoveCartItem(userID, cartID uint) error {
	return repository.RemoveCartItem(userID, cartID)
}

// clear cart
func ClearCart(userID uint) error {
	return repository.ClearCart(userID)
}
