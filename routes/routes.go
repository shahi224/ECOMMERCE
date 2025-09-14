package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	userGroup := router.Group("/user")
	UserRoutes(userGroup, db)

	adminGroup := router.Group("/admin")
	AdminRoutes(adminGroup, db)
}
