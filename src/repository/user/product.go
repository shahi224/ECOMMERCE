package repository

import (
	"gorm.io/gorm"

	"ECOMMERCE/utils/models"
)

type ProductRepository interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	SearchProducts(query string) ([]models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// get all products
func (r *productRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// pet produst by ID
func (r *productRepository) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// seach products
func (r *productRepository) SearchProducts(query string) ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Where("name LIKE ?", "%"+query+"%").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
