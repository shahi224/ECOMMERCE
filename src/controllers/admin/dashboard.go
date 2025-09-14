package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	services "ECOMMERCE/src/services/admin"
)

type DashboardHandler struct {
	service services.DashboardService
}

func NewDashboardHandler(s services.DashboardService) *DashboardHandler {
	return &DashboardHandler{service: s}
}

func (h *DashboardHandler) ShowDashboard(c *gin.Context) {
	// ✅ Token check
	token, err := c.Cookie("admin_token")
	if err != nil || token == "" {
		c.Redirect(http.StatusSeeOther, "/admin/Authentication/login")
		return
	}
	_, err = h.service.ValidateToken(token)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/admin/Authentication/login")
		return
	}

	// ✅ Get Overview (includes RecentOrders)
	overview, err := h.service.GetOverview()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Failed to load dashboard"})
		return
	}

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"Overview": overview,
	})
}

// ✅ API for frontend JS (/admin/orders/getAllOrders)
func (h *DashboardHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.service.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}
