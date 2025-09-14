package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ECOMMERCE/middleware"
	controllers "ECOMMERCE/src/controllers/user"
	repository "ECOMMERCE/src/repository/user"
	services "ECOMMERCE/src/services/user"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) *gin.RouterGroup {

	r.POST("/signup", controllers.UserSignUp)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)

	// r.POST("/otp-verify", controllers.VerifyOTP)
	// r.GET("/resend-otp", controllers.ResendOTP)

	//product routes
	productRepo := repository.NewProductRepository(db)
	productUsecase := services.NewProductUsecase(productRepo)
	productcontrollers := controllers.NewProductController(productUsecase)

	products := r.Group("/products")
	{
		products.GET("/getAllProducts", productcontrollers.GetAllProducts)
		products.GET("/getID/:id", productcontrollers.GetProductByID)
		products.GET("/search", productcontrollers.SearchProducts)
	}

	// category routs
	category := r.Group("/categories")
	{
		category.GET("/listAllCategory", controllers.ListAllCategories)
		category.GET("/:id/products", controllers.GetProductByCategoryID)
	}

	// userProfile routes
	userRoutes := r.Group("/userProfile", middleware.UserAuthMiddleware())
	{
		userRoutes.POST("/createUserProfile", controllers.CreateUserProfile)
		userRoutes.GET("/getUserProfile", controllers.GetUserProfile)
		userRoutes.PUT("/updateProfile", controllers.UpdateUserProfile)
		userRoutes.DELETE("/deleteUserProfile", controllers.DeleteUserProfile)
	}

	//  carts routes
	carts := r.Group("/cart", middleware.UserAuthMiddleware())
	{
		carts.POST("/addCart", controllers.AddToCart)
		carts.GET("/getAllCartProducts", controllers.GetAllCartProducts)
		carts.PUT("/update/:id", controllers.UpdateCartItem)
		carts.DELETE("/remove/:id", controllers.RemoveCartItem)
		carts.DELETE("/clearCart", controllers.ClearCart)
	}

	//  wishlist routes
	wishlist := r.Group("/wishlist", middleware.UserAuthMiddleware())
	{
		wishlist.POST("/addWishlist", controllers.AddToWishlist)
		wishlist.GET("/getWishlist", controllers.GetWishlist)
		wishlist.DELETE("/remove/:product_id", controllers.RemoveFromWishlist)
		wishlist.DELETE("/clearWishList", controllers.ClearWishlist)
	}

	// orders routes
	order := r.Group("/orders")
	order.Use(middleware.UserAuthMiddleware())
	{
		order.POST("/createOrder", controllers.CreateOrder)
		order.GET("/getUserOrders", controllers.GetUserOrders)
		order.GET("/:order_id", controllers.GetOrderById)
	}

	//  address routes
	address := r.Group("/address").Use(middleware.UserAuthMiddleware())
	{
		address.POST("/createAddress", controllers.CreateAddress)
		address.GET("/getAllAddress", controllers.GetAllAddress)
		address.PUT("/update/:address_id", controllers.UpdateAddress)
		address.DELETE("/delete/:id", controllers.DeleteAddress)
	}
	return r

}
