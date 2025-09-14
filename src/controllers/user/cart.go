package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	services "ECOMMERCE/src/services/user"
	"ECOMMERCE/utils/models"
	"ECOMMERCE/utils/response"
)

// add products to cart
func AddToCart(c *gin.Context) {
	var req models.AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid request", nil, err.Error()))
		return
	}

	if err := validator.New().Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "validation failed", nil, err.Error()))
		return
	}

	userIDVal, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.ClientResponse(http.StatusUnauthorized, "User not authenticated", nil, nil))
		return
	}
	userID, ok := userIDVal.(uint)
	if !ok || userID == 0 {
		c.JSON(http.StatusUnauthorized, response.ClientResponse(http.StatusUnauthorized, "Invalid user ID", nil, nil))
		return
	}
	// userID := userID.(uint)
	cartErr := services.AddToCart(userID, req)
	if cartErr != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "could not add to cart", nil, cartErr.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.ClientResponse(http.StatusCreated, "added to cart", cartErr, nil))
}

// get all cart products
func GetAllCartProducts(c *gin.Context) {
	userID, exists := c.Get("user_id")
	fmt.Println("user_id", userID)
	if !exists {
		c.JSON(http.StatusUnauthorized, response.ClientResponse(http.StatusUnauthorized, "user not authenticated", nil, nil))
		return
	}

	cart, err := services.GetAllCartProducts(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "could not get cart", nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "user cart", cart, nil))
}

// update cart items
func UpdateCartItem(c *gin.Context) {
	cartIDStr := c.Param("id")
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid cart ID", nil, err.Error()))
		return
	}

	var req models.UpdateCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid data", nil, err.Error()))
		return
	}

	userID := c.GetUint("user_id")
	err = services.UpdateCartItem(userID, uint(cartID), req)
	if err != nil {
		if err.Error() == "cart item not found" {
			c.JSON(http.StatusNotFound, response.ClientResponse(http.StatusNotFound, err.Error(), nil, nil))
			return
		}
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "could not update cart", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "cart item updated", nil, nil))
}

// remove products from cart
func RemoveCartItem(c *gin.Context) {
	cartIDStr := c.Param("id")
	cartID, err := strconv.Atoi(cartIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "Invalid Cart ID", nil, err.Error()))
		return
	}

	userID := c.GetUint("user_id")
	if err := services.RemoveCartItem(userID, uint(cartID)); err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "could not remove item", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "cart item removed", nil, nil))
}

// clear cart
func ClearCart(c *gin.Context) {
	userID := c.GetUint("user_id")
	if err := services.ClearCart(userID); err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "colud not clear cart", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "cart cleared", nil, nil))
}
