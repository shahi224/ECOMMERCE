package services

import (
	repository "ECOMMERCE/src/repository/admin"
	"ECOMMERCE/utils/models"
)

type OrderService interface {
	GetAllOrder() ([]models.Order, error)
	GetOrderByID(id uint) (*models.Order, error)
	UpdateOrderStatus(id uint, status string) error
	DeleteOrder(id uint) error
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(r repository.OrderRepository) OrderService {
	return &orderService{repo: r}
}

// get all orders
func (s *orderService) GetAllOrder() ([]models.Order, error) {
	return s.repo.GetAllOrders()
}

// get orders by ID
func (s *orderService) GetOrderByID(id uint) (*models.Order, error) {
	return s.repo.GetOrderByID(id)
}

// update order status
func (s *orderService) UpdateOrderStatus(id uint, status string) error {
	return s.repo.UpdateOrderStatus(id, status)
}

// delete order
func (s *orderService) DeleteOrder(id uint) error {
	return s.repo.DeleteOrder(id)
}
