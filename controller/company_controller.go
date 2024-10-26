package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/robinrev/go-rest-api/dtoconverter"
	"github.com/robinrev/go-rest-api/model"
	"github.com/robinrev/go-rest-api/service"
	"github.com/robinrev/go-rest-api/util"
	"gorm.io/gorm"
)

func GetAllCompanies(c *gin.Context) {
	datas, err := service.GetAllCompanies()

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
		Data:    dtoconverter.CompanyListModelToDto(datas)})
}

func GetCompanyById(c *gin.Context) {
	id := c.Param("id")
	result, err := service.GetCompanyById(id)

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
		Data:    dtoconverter.CompanyModelToDto(result)})
}

func AddNewCompany(c *gin.Context) {
	var data model.Company
	if err := c.BindJSON(&data); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, util.Response{ErrorCode: util.ERROR_CODE_WRONG_JSON_FORMAT,
			Message: util.WRONG_JSON_FORMAT})
		return
	}

	if err := service.AddNewCompany(data); err != nil {
		if err.Error() == "company code and name are required" {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{util.HTTP_ERROR: err.Error()})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, util.Response{ErrorCode: util.ERROR_CODE_DATABASE,
				Message: err.Error()})
		}
		return
	}

	c.IndentedJSON(http.StatusOK, util.Response{ErrorCode: util.ERROR_CODE_SUCCESS,
		Message: util.MESSAGE_SUCCESS,
		Data:    dtoconverter.CompanyModelToDto(data)})
}

func UpdateCompany(c *gin.Context) {
	var newData model.Company
	if err := c.BindJSON(&newData); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, util.Response{ErrorCode: util.ERROR_CODE_WRONG_JSON_FORMAT,
			Message: util.WRONG_JSON_FORMAT})
		return
	}

	existingData, err := service.GetCompanyById(strconv.Itoa(newData.CompanyID))

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

	updatedCompany, err := service.UpdateCompany(existingData, newData)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, util.Response{ErrorCode: util.ERROR_CODE_DATABASE,
			Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, util.Response{ErrorCode: util.ERROR_CODE_SUCCESS,
		Message: util.MESSAGE_SUCCESS,
		Data:    dtoconverter.CompanyModelToDto(updatedCompany)})
}
