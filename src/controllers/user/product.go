package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	services "ECOMMERCE/src/services/user"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductController(productUsecase services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productUsecase}
}

// get all products
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// get product by ID
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// search products
func (h *ProductHandler) SearchProducts(c *gin.Context) {
	query := c.DefaultQuery("query", "")
	products, err := h.productService.SearchProducts(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search products"})
		return
	}
	c.JSON(http.StatusOK, products)
}
