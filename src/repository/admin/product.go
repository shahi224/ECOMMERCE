package repository

import (
	"errors"

	"gorm.io/gorm"

	"ECOMMERCE/utils/models"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetByID(id uint) (*models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// get all products
func (r *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}

// get produst by ID
func (r *productRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("product not found")
	}
	return &product, err
}

// create product
func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

// update product
func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

// delete product
func (r *productRepository) Delete(id uint) error {
	res := r.db.Delete(&models.Product{}, id)
	if res.RowsAffected == 0 {
		return errors.New("product not found")
	}
	return res.Error
}
