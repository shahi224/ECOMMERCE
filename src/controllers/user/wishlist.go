package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	services "ECOMMERCE/src/services/user"
	"ECOMMERCE/utils/models"
	"ECOMMERCE/utils/response"
)

var validate = validator.New()

// add products to wishlist
func AddToWishlist(c *gin.Context) {
	var req models.AddToWishlistRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid request", nil, err.Error()))
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "validation failed", nil, err.Error()))
		return
	}

	userID := c.GetUint("user_id")
	if err := services.AddToWishlist(userID, req); err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "could not add to wishlist", nil, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.ClientResponse(http.StatusCreated, "added to wishlist", nil, nil))
}

// get wishlist
func GetWishlist(c *gin.Context) {
	userID := c.GetUint("user_id")
	items, err := services.GetWishlist(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "could not get wishlist", nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "wishlist", items, nil))
}

// remove products from wishlist
func RemoveFromWishlist(c *gin.Context) {
	productIDStr := c.Param("product_id")
	if productIDStr == "" {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "product id is required", nil, nil))
		return
	}

	productID, err := strconv.Atoi(productIDStr)
	// id, _ := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid product id", nil, err.Error()))
		return
	}

	userID := c.GetUint("user_id")
	if err := services.RemoveFromWishlist(userID, uint(productID)); err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "could not remove wishlist item", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "removed from wishlist", nil, nil))
}

// clear wishlist
func ClearWishlist(c *gin.Context) {
	userID := c.GetUint("user_id")
	if err := services.ClearWishlist(userID); err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "could not remove wishlist item", nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "clear wishlist", nil, nil))
}
