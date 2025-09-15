package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	services "ECOMMERCE/src/services/admin"
)

type AuthHandler struct {
	Services *services.AuthService
}

// admin signup
func (h *AuthHandler) SignupAdmin(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := h.Services.SignupAdmin(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "admin registered", "user": user})
}

// show login page
func (h *AuthHandler) ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// admin login
func (h *AuthHandler) LoginAdmin(c *gin.Context) {
	var req struct {
		Email    string `form:"email" json:"email" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	if err := c.ShouldBind(&req); err != nil {
		if strings.Contains(c.GetHeader("Accept"), "text/html") {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Invalid input"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		}
		return
	}

	token, err := h.Services.LoginAdmin(req.Email, req.Password)
	if err != nil {
		if strings.Contains(c.GetHeader("Accept"), "text/html") {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		}
		return
	}

	// Set cookie for browser requests
	if strings.Contains(c.GetHeader("Accept"), "text/html") {
		c.SetCookie("admin_token", token, 3600, "/", "", false, true)
		c.Redirect(http.StatusSeeOther, "/admin/dashboard")
	} else {
		// Return token for API requests (like Postman)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

// admin logout
func (h *AuthHandler) Logout(c *gin.Context) {
	c.SetCookie("admin_token", "", -1, "/", "", false, true)

	if strings.Contains(c.GetHeader("Accept"), "text/html") {
		c.Redirect(http.StatusSeeOther, "/admin/Authentication/login")
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
	}
}
