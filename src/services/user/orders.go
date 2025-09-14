package services

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	repository "ECOMMERCE/src/repository/user"
	"ECOMMERCE/utils/models"
)

// create order
func CreateOrder(userID uint, req models.CreateOrderRequest) error {
	var total float64
	var items []models.OrderItem

	for _, item := range req.Items {
		product, err := repository.GetProductById(item.ProductID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("product with ID %d not found", item.ProductID)
			}
			return fmt.Errorf("error fetching product: %v", err)
		}

		price := product.Price * float64(item.Quantity)
		total += price

		items = append(items, models.OrderItem{
			ProductID: &product.ID,
			Quantity:  item.Quantity,
			Price:     price,
		})
	}

	order := models.Order{
		UserID: userID,
		// TotalAmount: total,
		AddressID:   &req.AddressID,
		TotalAmount: total,
		Status:      "pending",
		// ShippingAddress: ,
		OrderItems: items,
	}

	return repository.CreateOrder(&order)
}

// get user orders
func GetUserOrders(userID uint) ([]models.Order, error) {
	return repository.GetUserOrders(userID)
}

// get order by id
func GetOrderById(orderID, userID uint) (models.Order, error) {
	return repository.GetOrderById(orderID, userID)
}
