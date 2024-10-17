package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robinrev/go-rest-api/model"
	"github.com/robinrev/go-rest-api/repository"
)

func GetAllCompany(context *gin.Context) {
	log.Println("get all company")
	company := repository.GetAllCompany()
	context.IndentedJSON(http.StatusOK, company)
}

func GetCompanyById(context *gin.Context) {
	log.Println("get company by id")
	id := context.Param("id")
	log.Println("id : ", id)
	company := repository.GetCompanyById(id)
	context.IndentedJSON(http.StatusOK, company)
}

func GetCompanyByCompanyCode(context *gin.Context) {
	log.Println("get company by company code")
	companyCode := context.Query("companyCode")
	log.Println("company code : ", companyCode)
	company := repository.GetCompanyByCompanyCode(companyCode)
	context.IndentedJSON(http.StatusOK, company)
}

func AddNewCompany(context *gin.Context) {
	log.Println("add new company")
	var company model.Company
	if err := context.BindJSON(&company); err != nil {
		return
	}
	err := repository.AddNewCompany(company)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, err)
	}
	context.IndentedJSON(http.StatusOK, company)
}

func UpdateCompany(context *gin.Context) {
	log.Println("update company")
	var company model.Company
	if err := context.BindJSON(&company); err != nil {
		return
	}
	err := repository.UpdateCompany(company)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, err)
	}
	context.IndentedJSON(http.StatusOK, company)
}
