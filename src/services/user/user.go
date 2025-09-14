package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	repository "ECOMMERCE/src/repository/user"
	"ECOMMERCE/utils/helper"
	"ECOMMERCE/utils/models"
)

// user signup
func UserSignUp(req models.SignUpRequest) error {
	existsEmail, _ := repository.GetUserByEmail(req.Email)
	if existsEmail != nil {
		return errors.New("user already exists")
	}

	existPhone, _ := repository.CheckUserByPhone(req.Phone)
	if existPhone != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return err
	}

	otp := helper.GenerateOTP()

	user := models.User{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: string(hashedPassword),
		OTP:      otp,
		Role:     "user",
	}

	return repository.CreateUser(&user)
}

// user login
func Login(req models.LoginRequest) (string, error) {

	user, err := repository.CheckUserByPhone(req.Phone)
	if err != nil {
		return "", errors.New("user not found")
	}
	// user, err := repository.GetUserByEmail(req.Email)
	// if err != nil {
	// 	return " ", errors.New("Invalid credentials")
	// }

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := helper.GenerateJWT(user.ID, user.Email, "user")
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return token, nil
}
