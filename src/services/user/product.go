package services

import (
	repository "ECOMMERCE/src/repository/user"
	"ECOMMERCE/utils/models"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	SearchProducts(query string) ([]models.Product, error)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductUsecase(productRepo repository.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

// get all products
func (u *productService) GetAllProducts() ([]models.Product, error) {
	return u.productRepo.GetAllProducts()
}

// get produst by ID
func (u *productService) GetProductByID(id uint) (*models.Product, error) {
	return u.productRepo.GetProductByID(id)
}

// search product
func (u *productService) SearchProducts(query string) ([]models.Product, error) {
	return u.productRepo.SearchProducts(query)
}
