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

// create user address
func CreateAddress(c *gin.Context) {
	var req models.CreateAddressRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid requset", nil, err.Error()))
		return
	}

	if err := validator.New().Struct(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid requset", nil, err.Error()))
		return
	}

	userID := c.GetUint("user_id")
	if err := services.CreateAddress(userID, req); err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "invalid requset", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "address fetched", nil, nil))
}

// get all address
func GetAllAddress(c *gin.Context) {
	userID := c.GetUint("user_id")
	address, err := services.GetAllAdddress(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid requset", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "address fetched", address, nil))

}

// update users
func UpdateAddress(c *gin.Context) {
	addressIDStr := c.Param("address_id")
	addressID, err := strconv.Atoi(addressIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid requset", nil, err.Error()))
		return
	}

	var req models.UpdateAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid requset", nil, err.Error()))
		return
	}

	if err := validator.New().Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "validation failed", nil, err.Error()))
		return
	}

	userID := c.GetUint("user_id")
	if err := services.UpdateAddress(userID, uint(addressID), req); err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "invalid requset", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "address updated", nil, nil))

}

// delete user address
func DeleteAddress(c *gin.Context) {
	addressIDStr := c.Param("id")
	addressID, err := strconv.Atoi(addressIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "invalid address id", nil, err.Error()))
		return
	}

	userID := c.GetUint("user_id")
	if err := services.DeleteAddress(userID, uint(addressID)); err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "could not delete address", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "address deleted", nil, nil))

}
