package repository

import (
	"errors"

	"gorm.io/gorm"

	"ECOMMERCE/database"
	"ECOMMERCE/utils/models"
)

func CreateUserProfile(userID uint, profile *models.UserProfile) error {
	var existing models.UserProfile
	err := database.DB.Unscoped().Where("user_id = ?", userID).First(&existing).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return database.DB.Create(&profile).Error
		}
		return err
	}

	if existing.DeletedAt.Valid {
		return database.DB.Model(&existing).Updates(map[string]interface{}{
			"deleted_at": nil,
			"name":       profile.Name,
			"email":      profile.Email,
			"phone":      profile.Phone,
			"dob":        profile.DOB,
			"gender":     profile.Gender,
		}).Error
	}

	return errors.New("profile already exists")
}

// get user profile by ID
func GetUserProfileByUserID(userID uint) (*models.UserProfile, error) {
	var profile models.UserProfile
	err := database.DB.Where("user_id = ?", userID).First(&profile).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &profile, nil

}

// update user profile
func UpdateUserProfile(userID uint, req *models.UpdateUserProfileRequest, updates map[string]interface{}) error {

	return database.DB.Model(&models.UserProfile{}).Where("user_id = ?", userID).Updates(updates).Error
}

// delete user profile
func DeleteUserProfile(userID uint) error {
	return database.DB.Where("user_id = ?", userID).Delete(&models.UserProfile{}).Error
}
