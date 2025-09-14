package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"ECOMMERCE/database"
	"ECOMMERCE/utils/models"
)

func AddToCart(userID uint, req models.AddToCartRequest) error {
	var existing models.Cart
	err := database.DB.Where("user_id = ? AND product_id = ?", userID, req.ProductID).First(&existing).Error
	if err == nil {
		existing.Quantity += float64(req.Quantity)
		return database.DB.Save(&existing).Error
	}

	product := models.Product{}
	if err := database.DB.First(&product, req.ProductID).Error; err != nil {
		return errors.New("product not found")
	}

	cart := models.Cart{
		UserID:     userID,
		ProductID:  req.ProductID,
		Quantity:   float64(req.Quantity),
		TotalPrice: float64(req.Quantity) * product.Price,
	}
	return database.DB.Create(&cart).Error
}

func GetAllCartProducts(userID uint) ([]models.Cart, error) {
	var carts []models.Cart

	err := database.DB.Preload("Product").Where("user_id = ?", userID).Find(&carts).Error
	return carts, err
}

func UpdateCartItem(userID, cartID uint, req models.UpdateCartRequest) error {
	var cart models.Cart
	// err := database.DB.First(&cart, "id = ? AND user_id = ?", cartID, userID).Error
	if err := database.DB.Where("id = ? AND user_id = ?", cartID, userID).First(&cart).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("cart item not found")
		}
		return err
	}
	cart.Quantity = float64(req.Quantity)
	// cart.TotalPrice = float64(req.Quantity) * cart.Product.Price
	return database.DB.Save(&cart).Error
}
func RemoveCartItem(userID, cartID uint) error {
	var cart models.Cart
	return database.DB.Where("id = ? AND user_id =?", cartID, userID).Delete(&cart).Error
}

func ClearCart(userID uint) error {
	var cart models.Cart
	return database.DB.Where("user_id = ?", userID).Delete(&cart).Error
}
