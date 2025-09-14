package services

import (
	repository "ECOMMERCE/src/repository/admin"
	"ECOMMERCE/utils/models"
)

type ProfileService interface {
	GetAdmin(id uint) (models.User, error)
	UpdateAdmin(admin models.User) error
}

type profileService struct {
	repo repository.ProfileRepository
}

func NewProfileService(repo repository.ProfileRepository) ProfileService {
	return &profileService{repo: repo}
}

func (s *profileService) GetAdmin(id uint) (models.User, error) {
	return s.repo.GetAdminByID(id)
}

func (s *profileService) UpdateAdmin(admin models.User) error {
	return s.repo.UpdateAdmin(admin)
}
