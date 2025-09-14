package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ECOMMERCE/database"
	"ECOMMERCE/middleware"
	controllers "ECOMMERCE/src/controllers/admin"
	repository "ECOMMERCE/src/repository/admin"
	services "ECOMMERCE/src/services/admin"
)

func AdminRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {
	//adminRoutes
	adminRepo := &repository.AdminRepository{DB: db}
	authService := &services.AuthService{Repo: adminRepo}
	authHandler := &controllers.AuthHandler{Services: authService}

	// r.GET("/login", func(c *gin.Context) {
	// 	c.HTML(200, "login.html", nil)
	// })

	adminAuth := r.Group("/Authentication")
	{
		adminAuth.POST("/signup", authHandler.SignupAdmin)
		adminAuth.POST("/login", authHandler.LoginAdmin)
		adminAuth.GET("/login", authHandler.ShowLoginPage)
		adminAuth.GET("/logout", authHandler.Logout)
	}

	// Profile repo/service/controller
	profileRepo := repository.NewProfileRepository(db)
	profileService := services.NewProfileService(profileRepo)
	profileController := controllers.NewProfileController(profileService)

	// Group routes under /admin/profile
	profile := r.Group("/profile", middleware.AdminAuthMiddleware())
	{
		profile.GET("/show", profileController.ShowProfile)      // show profile page
		profile.POST("/update", profileController.UpdateProfile) // update profile
	}

	// category routes
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := controllers.NewCategoryHandler(categoryService)
	admin := r.Group("/category", middleware.AdminAuthMiddleware())
	{
		admin.POST("/createCategory", categoryHandler.CreateCategory)
		admin.PUT("/update/:id", categoryHandler.UpdateCategory)
		admin.DELETE("/delete/:id", categoryHandler.DeleteCategory)
		admin.GET("/Categories", categoryHandler.GetAllCategories)
		admin.GET("/:id", categoryHandler.GetCategoryByID)
	}

	// product routes
	productRepo := repository.NewProductRepository(db)
	productUsecase := services.NewProductUsecase(productRepo)
	productHandler := controllers.NewProductHandler(productUsecase)

	product := r.Group("/products", middleware.AdminAuthMiddleware())
	{
		product.POST("/createProduct", productHandler.Create)
		product.GET("/getAllProducts", productHandler.GetAll)
		product.PUT("/:id", productHandler.Update)
		product.DELETE("/:id", productHandler.Delete)
	}

	// user routes
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserUsecase(userRepo)
	userHandler := controllers.NewUserHandler(userService)
	userGroup := r.Group("/users", middleware.AdminAuthMiddleware())
	{
		userGroup.POST("/CreateUser", userHandler.CreateUser)
		userGroup.PUT("/:id/UpdateUser", userHandler.UpdateUser)
		userGroup.GET("/getAllUsers", userHandler.GetAllUsers)
		userGroup.PUT("/:id/unblock", userHandler.UnblockUser)
		userGroup.PUT("/:id/block", userHandler.BlockUser)
		userGroup.DELETE("/delete/:id", userHandler.DeleteUser)
	}

	// order routes
	orderRepo := repository.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderHandler := controllers.NewOrderHandleer(orderService)

	order := r.Group("/orders", middleware.AdminAuthMiddleware())
	{
		order.GET("/getAllOrders", orderHandler.GetAllOrder)
		order.GET("/:id", orderHandler.GetOrderByID)
		order.PUT("/:id/status", orderHandler.UpdateOrderStatus)
		order.DELETE("/:id", orderHandler.DeleteOrder)
	}

	// dashboard routes
	dashRepo := repository.NewDashboardRepo(database.DB)
	dashService := services.NewDashboardService(dashRepo)
	dashHandler := controllers.NewDashboardHandler(dashService)

	dashboard := r.Group("/dashboard", middleware.AdminAuthMiddleware())
	{
		{
			dashboard.GET(" ", dashHandler.ShowDashboard)
		}
	}

	r.GET("/dashboard", func(c *gin.Context) {
		token, err := c.Cookie("admin_token")
		if err != nil || token == "" {
			c.Redirect(http.StatusSeeOther, "/admin/Authentication/login")
			return
		}
		c.HTML(200, "dashboard.html", gin.H{"title": "Admin Dashboard"})
	})

	return r
}
