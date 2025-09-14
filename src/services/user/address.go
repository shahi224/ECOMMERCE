package services

import (
	repository "ECOMMERCE/src/repository/user"
	"ECOMMERCE/utils/models"
)

// create user address
func CreateAddress(userID uint, req models.CreateAddressRequest) error {

	address := models.Address{
		UserID:   userID,
		FullName: req.FullName,
		Phone:    req.Phone,
		House:    req.House,
		Street:   req.Street,
		City:     req.City,
		State:    req.State,
		Pincode:  req.Pincode,
	}
	return repository.CreateAddress(userID, &address)
}

// get all address
func GetAllAdddress(userID uint) ([]models.Address, error) {
	return repository.GetAllAddress(userID)
}

// update address
func UpdateAddress(userID, addressID uint, req models.UpdateAddressRequest) error {
	address := models.Address{
		FullName: req.FullName,
		Phone:    req.Phone,
		House:    req.House,
		Street:   req.Street,
		City:     req.City,
		State:    req.State,
		Pincode:  req.Pincode,
	}

	return repository.UpdateAddress(userID, addressID, &address)
}

// delete address
func DeleteAddress(userID, addressID uint) error {
	return repository.DeleteAddress(userID, addressID)
}
