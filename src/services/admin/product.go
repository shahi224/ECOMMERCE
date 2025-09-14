package services

import (
	repository "ECOMMERCE/src/repository/admin"
	"ECOMMERCE/utils/models"
)

type ProductUsecase interface {
	GetAll() ([]models.Product, error)
	Create(product *models.Product) error
	Update(id uint, updated *models.Product) error
	Delete(id uint) error
}

type productUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{repo: repo}
}

// get all products
func (u *productUsecase) GetAll() ([]models.Product, error) {
	return u.repo.GetAll()
}

// create products
func (u *productUsecase) Create(product *models.Product) error {
	return u.repo.Create(product)
}

// update product
func (u *productUsecase) Update(id uint, updated *models.Product) error {
	existing, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}
	existing.Name = updated.Name
	existing.Description = updated.Description
	existing.Price = updated.Price
	existing.Stock = updated.Stock
	existing.Category = updated.Category

	return u.repo.Update(existing)
}

// delete product
func (u *productUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
