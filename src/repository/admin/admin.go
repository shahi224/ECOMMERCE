package repository

import (
	"gorm.io/gorm"

	"ECOMMERCE/utils/models"
)

type AdminRepository struct {
	DB *gorm.DB
}

// create user
func (r *AdminRepository) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

// find user by email
func (r *AdminRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.DB.Where("email = ?", email).First(&user)
	return &user, result.Error
}
