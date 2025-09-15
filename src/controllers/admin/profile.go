package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	services "ECOMMERCE/src/services/admin"
)

type ProfileController struct {
	service services.ProfileService
}

func NewProfileController(service services.ProfileService) *ProfileController {
	return &ProfileController{service: service}
}

// Show profile page
func (pc *ProfileController) ShowProfile(c *gin.Context) {
	// Get adminID from context (set by middleware)
	adminID, exists := c.Get("adminID")
	if !exists {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{"error": "Unauthorized access"})
		return
	}

	admin, err := pc.service.GetAdmin(adminID.(uint))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Failed to fetch profile: " + err.Error()})
		return
	}

	c.HTML(http.StatusOK, "profile.html", gin.H{"Admin": admin})
}

// Handle profile update
func (pc *ProfileController) UpdateProfile(c *gin.Context) {
	// Get adminID from context (set by middleware)
	adminID, exists := c.Get("adminID")
	if !exists {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{"error": "Unauthorized access"})
		return
	}

	admin, err := pc.service.GetAdmin(adminID.(uint))
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

	// Add success message and redirect to profile page
	c.Redirect(http.StatusSeeOther, "/admin/profile?success=Profile updated successfully")
}

// Add this method to handle the redirect with success message
func (pc *ProfileController) ShowProfileWithMessage(c *gin.Context) {
	// Get adminID from context (set by middleware)
	adminID, exists := c.Get("adminID")
	if !exists {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{"error": "Unauthorized access"})
		return
	}

	admin, err := pc.service.GetAdmin(adminID.(uint))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Failed to fetch profile: " + err.Error()})
		return
	}

	// Check for success message in query params
	success := c.Query("success")

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"Admin":   admin,
		"success": success,
	})
}
