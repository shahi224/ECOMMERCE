// internal/repository/user_repository.go
package repository

import (
	"errors"

	"gorm.io/gorm"

	"ECOMMERCE/utils/models"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByPhone(phone string) (*models.User, error)
	Update(user *models.User) error
	GetByID(userID uint) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	UpdateUserStatus(userID uint, status string) error
	DeleteUser(userID uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByPhone(phone string) (*models.User, error) {
	var user models.User
	err := r.db.Where("phone = ?", phone).First(&user).Error
	return &user, err
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) GetByID(userID uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, userID).Error
	return &user, err
}

// get all users
func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	// if err := r.db.Preload("UserProfile").Find(&users).Error; err != nil {
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// update user status
func (r *userRepository) UpdateUserStatus(userID uint, status string) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("status", status).Error
}

// delete user
func (r *userRepository) DeleteUser(userID uint) error {
	res := r.db.Delete(&models.User{}, userID)
	if res.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return res.Error
}
