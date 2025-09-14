package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	repository "ECOMMERCE/src/repository/admin"
	"ECOMMERCE/utils/models"
)

type UserUsecase interface {
	CreateUser(req models.SignUpRequest) error
	GetUser(userID uint) (*models.User, error)
	UpdateUser(userID uint, req models.SignUpRequest) error
	GetAllUsers() ([]models.User, error)
	BlockUser(userID uint) error
	UnblockUser(userID uint) error
	DeleteUser(userID uint) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (s *userUsecase) CreateUser(req models.SignUpRequest) error {
	// Check if phone already exists
	existingUser, _ := s.repo.GetByPhone(req.Phone)
	if existingUser != nil && existingUser.ID != 0 {
		return errors.New("user with this phone already exists")
	}

	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: string(hashedPass),
		Status:   "active",
		Role:     "user",
		IsAdmin:  false,
	}

	return s.repo.Create(user)
}

func (s *userUsecase) UpdateUser(userID uint, req models.SignUpRequest) error {
	user, err := s.repo.GetByID(userID)
	if err != nil {
		return err
	}

	user.Name = req.Name
	user.Email = req.Email
	user.Phone = req.Phone
	if req.Password != "" {
		hashedPass, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPass)
	}

	return s.repo.Update(user)
}

func (s *userUsecase) GetUser(userID uint) (*models.User, error) {
	return s.repo.GetByID(userID)
}

// get all users
func (u *userUsecase) GetAllUsers() ([]models.User, error) {
	return u.repo.GetAllUsers()
}

// block user
func (u *userUsecase) BlockUser(userID uint) error {
	return u.repo.UpdateUserStatus(userID, "blocked")
}

// unblock user
func (u *userUsecase) UnblockUser(userID uint) error {
	return u.repo.UpdateUserStatus(userID, "active")
}

// delete user
func (u *userUsecase) DeleteUser(userID uint) error {
	return u.repo.DeleteUser(userID)
}
