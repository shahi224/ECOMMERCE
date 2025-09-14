package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	services "ECOMMERCE/src/services/user"
	"ECOMMERCE/utils/response"
)

// list all categories
func ListAllCategories(c *gin.Context) {
	categories, err := services.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "failed to fetch categories", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "categories fetched", categories, nil))
}

// get produst by category ID
func GetProductByCategoryID(c *gin.Context) {
	categoryIDStr := c.Param("id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid category ID", nil, err.Error()))
		return
	}

	products, err := services.GetProductByCategoryID(uint(categoryID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "failed to fetch products", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "products fetched", products, nil))

}
