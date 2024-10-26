package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/robinrev/go-rest-api/model"
	"github.com/robinrev/go-rest-api/service"
	"github.com/robinrev/go-rest-api/util"
	"gorm.io/gorm"
)

func GetAllCustomer(c *gin.Context) {
	data, err := service.GetAllCustomer()

	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			c.IndentedJSON(http.StatusNotFound, gin.H{util.HTTP_ERROR: gorm.ErrRecordNotFound})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{util.HTTP_ERROR: gorm.ErrRecordNotFound})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, data)
}

func GetCustomerById(c *gin.Context) {
	id := c.Param("id")
	result, err := service.GetCustomerById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{util.HTTP_ERROR: gorm.ErrRecordNotFound.Error()})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{util.HTTP_ERROR: err})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, result)
}

func AddNewCustomer(c *gin.Context) {
	var data model.Customer
	if err := c.BindJSON(&data); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{util.HTTP_ERROR: "invalid request"})
		return
	}

	if err := service.AddNewCustomer(data); err != nil {
		if err.Error() == "username and password are required" {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{util.HTTP_ERROR: err.Error()})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{util.HTTP_ERROR: err})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, data)
}

func UpdateCustomer(c *gin.Context) {
	var newData model.Customer
	if err := c.BindJSON(&newData); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{util.HTTP_ERROR: err})
		return
	}

	existingData, err := service.GetCustomerById(strconv.Itoa(newData.CustomerID))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, gin.H{util.HTTP_ERROR: util.ERROR_NO_DATA_FOUND})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{util.HTTP_ERROR: err})
			return
		}
	}

	updatedCustomer, err := service.UpdateCustomer(existingData, newData)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{util.HTTP_ERROR: err})
		return
	}

	c.IndentedJSON(http.StatusOK, util.Response{ErrorCode: util.ERROR_CODE_SUCCESS,
		Message: util.MESSAGE_SUCCESS,
		Data:    updatedCustomer})
}
