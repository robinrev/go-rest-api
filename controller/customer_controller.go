package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	dtocoverter "github.com/robinrev/go-rest-api/dto_coverter"
	"github.com/robinrev/go-rest-api/model"
	"github.com/robinrev/go-rest-api/service"
	"github.com/robinrev/go-rest-api/util"
	"gorm.io/gorm"
)

func GetAllCustomer(c *gin.Context) {
	data, err := service.GetAllCustomer()

	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			c.IndentedJSON(http.StatusNotFound, util.Response{ErrorCode: util.ERROR_CODE_EMPTY_DATA,
				Message: gorm.ErrRecordNotFound.Error()})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, util.Response{ErrorCode: util.ERROR_CODE_DATABASE,
				Message: err.Error()})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, util.Response{ErrorCode: util.ERROR_CODE_SUCCESS,
		Message: util.MESSAGE_SUCCESS,
		Data:    dtocoverter.CustomerListModelToDto(data)})
}

func GetAllCustomersByCompany(c *gin.Context) {
	companyIDStr := c.Param("companyId")
	companyID, err := strconv.Atoi(companyIDStr)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, util.Response{ErrorCode: util.ERROR_CODE_PARAMETER,
			Message: util.WRONG_PARAMETER_FORMAT})
		return
	}

	data, err := service.GetAllCustomersByCompany(companyID)

	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			c.IndentedJSON(http.StatusNotFound, util.Response{ErrorCode: util.ERROR_CODE_EMPTY_DATA,
				Message: gorm.ErrRecordNotFound.Error()})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, util.Response{ErrorCode: util.ERROR_CODE_DATABASE,
				Message: err.Error()})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, util.Response{ErrorCode: util.ERROR_CODE_SUCCESS,
		Message: util.MESSAGE_SUCCESS,
		Data:    dtocoverter.CustomerListModelToDto(data)})
}

func GetCustomerById(c *gin.Context) {
	id := c.Param("id")
	result, err := service.GetCustomerById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, util.Response{ErrorCode: util.ERROR_CODE_EMPTY_DATA,
				Message: gorm.ErrRecordNotFound.Error()})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, util.Response{ErrorCode: util.ERROR_CODE_DATABASE,
				Message: err.Error()})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, util.Response{ErrorCode: util.ERROR_CODE_SUCCESS,
		Message: util.MESSAGE_SUCCESS,
		Data:    dtocoverter.CustomerModelToDto(result)})
}

func AddNewCustomer(c *gin.Context) {
	var data model.Customer
	if err := c.BindJSON(&data); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, util.Response{ErrorCode: util.ERROR_CODE_WRONG_JSON_FORMAT,
			Message: util.WRONG_JSON_FORMAT})
		return
	}

	if err := service.AddNewCustomer(data); err != nil {
		if err.Error() == "username and password are required" {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{util.HTTP_ERROR: err.Error()})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, util.Response{ErrorCode: util.ERROR_CODE_DATABASE,
				Message: err.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, util.Response{ErrorCode: util.ERROR_CODE_SUCCESS,
		Message: util.MESSAGE_SUCCESS,
		Data:    dtocoverter.CustomerModelToDto(data)})
}

func UpdateCustomer(c *gin.Context) {
	var newData model.Customer
	if err := c.BindJSON(&newData); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, util.Response{ErrorCode: util.ERROR_CODE_WRONG_JSON_FORMAT,
			Message: util.WRONG_JSON_FORMAT})
		return
	}

	existingData, err := service.GetCustomerById(strconv.Itoa(newData.CustomerID))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusNotFound, util.Response{ErrorCode: util.ERROR_CODE_EMPTY_DATA,
				Message: util.ERROR_NO_DATA_FOUND})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, util.Response{ErrorCode: util.ERROR_CODE_DATABASE,
				Message: err.Error()})
			return
		}
	}

	updatedCustomer, err := service.UpdateCustomer(existingData, newData)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, util.Response{ErrorCode: util.ERROR_CODE_DATABASE,
			Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, util.Response{ErrorCode: util.ERROR_CODE_SUCCESS,
		Message: util.MESSAGE_SUCCESS,
		Data:    dtocoverter.CustomerModelToDto(updatedCustomer)})
}
