package services

import (
	"errors"
	"time"

	repository "ECOMMERCE/src/repository/user"
	"ECOMMERCE/utils/models"
)

// create user profile
func CreateUserProfile(userID uint, req *models.CreateUserProfileRequest) error {
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return err
	}

	profile := &models.UserProfile{
		UserID: userID,
		Name:   req.Name,
		Email:  req.Email,
		Phone:  req.Phone,
		DOB:    &dob,
		Gender: req.Gender,
	}

	return repository.CreateUserProfile(userID, profile)
}

// get user profile
func GetUserProfile(userID uint) (*models.UserProfile, error) {
	profile, err := repository.GetUserProfileByUserID(userID)
	if err != nil {
		return nil, err
	}
	if profile == nil {
		return nil, errors.New("profile not found")
	}
	return profile, nil
}

// update user profile
func UpdateUserProfile(userID uint, req *models.UpdateUserProfileRequest) error {
	updates := map[string]interface{}{}

	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Phone != "" {
		updates["phone"] = req.Phone
	}
	if req.DOB != "" {
		dob, err := time.Parse("2006-01-02", req.DOB)
		if err != nil {
			return err
		}
		updates["dob"] = dob
	}
	if req.Gender != "" {
		updates["gender"] = req.Gender
	}

	if len(updates) == 0 {
		return errors.New("no fields to update")
	}

	return repository.UpdateUserProfile(userID, req, updates)
}

// delete user profile
func DeleteUserProfile(userID uint) error {
	return repository.DeleteUserProfile(userID)
}
