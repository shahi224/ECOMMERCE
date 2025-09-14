package services

import (
	repository "ECOMMERCE/src/repository/admin"
	"ECOMMERCE/utils/models"
)

type CategoryService interface {
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	DeleteCategory(id uint) error
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id uint) (*models.Category, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

// create category
func (s *categoryService) CreateCategory(category *models.Category) error {
	return s.repo.Create(category)
}

// update categories
func (s *categoryService) UpdateCategory(category *models.Category) error {
	return s.repo.Update(category)
}

// delete category
func (s *categoryService) DeleteCategory(id uint) error {
	return s.repo.Delete(id)
}

// get all categories
func (s *categoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.GetAll()
}

// get category by ID
func (s *categoryService) GetCategoryByID(id uint) (*models.Category, error) {
	return s.repo.GetByID(id)
}
