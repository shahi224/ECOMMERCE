package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	services "ECOMMERCE/src/services/admin"
	"ECOMMERCE/utils/helper"
)

type ProfileController struct {
	service services.ProfileService
}

func NewProfileController(service services.ProfileService) *ProfileController {
	return &ProfileController{service: service}
}

// Show profile page
func (pc *ProfileController) ShowProfile(c *gin.Context) {
	// Get adminID from JWT token instead of context
	tokenString, err := c.Cookie("admin_token")
	if err != nil {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{"error": "Unauthorized access"})
		return
	}

	// Extract admin ID from token
	adminID, err := extractAdminIDFromToken(tokenString)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{"error": "Invalid token: " + err.Error()})
		return
	}

	admin, err := pc.service.GetAdmin(adminID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Failed to fetch profile: " + err.Error()})
		return
	}

	c.HTML(http.StatusOK, "profile.html", gin.H{"Admin": admin})
}

// Handle profile update
func (pc *ProfileController) UpdateProfile(c *gin.Context) {
	// Get adminID from JWT token
	tokenString, err := c.Cookie("admin_token")
	if err != nil {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{"error": "Unauthorized access"})
		return
	}

	// Extract admin ID from token
	adminID, err := extractAdminIDFromToken(tokenString)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{"error": "Invalid token: " + err.Error()})
		return
	}

	admin, err := pc.service.GetAdmin(adminID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Profile not found: " + err.Error()})
		return
	}

	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	if name != "" {
		admin.Name = name
	}
	if email != "" {
		admin.Email = email
	}
	if password != "" {
		// Hash the password before storing
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Failed to hash password"})
			return
		}
		admin.Password = string(hashedPassword)
	}

	if err := pc.service.UpdateAdmin(admin); err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Update failed: " + err.Error()})
		return
	}

	c.Redirect(http.StatusSeeOther, "/admin/profile")
}

// Helper function to extract admin ID from JWT token
func extractAdminIDFromToken(tokenString string) (uint, error) {
	claims, err := helper.ParseToken(tokenString)
	if err != nil {
		return 0, err
	}

	// Your JWT uses "user_id" claim, not "admin_id"
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("invalid user ID in token")
	}

	return uint(userIDFloat), nil
}
