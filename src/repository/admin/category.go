package repository

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"

	"ECOMMERCE/utils/models"
)

type CategoryRepository interface {
	Create(category *models.Category) error
	Update(category *models.Category) error
	Delete(id uint) error
	GetAll() ([]models.Category, error)
	GetByID(id uint) (*models.Category, error)
}

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepo{db: db}
}

// craete categories
func (r *categoryRepo) Create(category *models.Category) error {
	var existing models.Category
	if err := r.db.Where("LOWER(name) = ?", strings.ToLower(category.Name)).First(&existing).Error; err == nil {
		return fmt.Errorf("category %q already exists", category.Name)
	}
	if err := r.db.Create(category).Error; err != nil {
		fmt.Println("DB error:", err)
		return err
	}
	return nil
}

// update categories
func (r *categoryRepo) Update(category *models.Category) error {
	return r.db.Model(&models.Category{}).Where("id = ?", category.ID).Updates(map[string]interface{}{
		"name":       category.Name,
		"created_by": category.CreatedBy,
	}).Error
}

// delete category
func (r *categoryRepo) Delete(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}

// get all categories
func (r *categoryRepo) GetAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

// get categories by ID
func (r *categoryRepo) GetByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.First(&category, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &category, err
}
