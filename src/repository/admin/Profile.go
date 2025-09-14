package repository

import (
	"gorm.io/gorm"

	"ECOMMERCE/utils/models"
)

type ProfileRepository interface {
	GetAdminByID(id uint) (models.User, error)
	UpdateAdmin(admin models.User) error
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db: db}
}

func (r *profileRepository) GetAdminByID(id uint) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ? AND is_admin = ?", id, true).First(&user).Error
	return user, err
}

func (r *profileRepository) UpdateAdmin(user models.User) error {
	return r.db.Save(&user).Error
}
