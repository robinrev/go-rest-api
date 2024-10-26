package service

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robinrev/go-rest-api/model"
	"github.com/robinrev/go-rest-api/repository"
)

func GetAllCompanies() ([]model.Company, error) {
	log.Println("get all company")
	return repository.GetAllCompanies()
}

func GetCompanyByCompanyCode(context *gin.Context) {
	log.Println("get company by company code")
	companyCode := context.Query("companyCode")
	log.Println("company code : ", companyCode)
	company := repository.GetCompanyByCompanyCode(companyCode)
	context.IndentedJSON(http.StatusOK, company)
}

func GetCompanyById(id string) (model.Company, error) {
	log.Println("get customer by id :", id)
	return repository.GetCompanyById(id)
}

func AddNewCompany(data model.Company) error {
	log.Println("add new customer")

	if data.CompanyCode == "" || data.CompanyName == "" {
		return errors.New("company code and name are required")
	}

	data.CreateDate = time.Now()
	data.UpdateDate = time.Now()
	data.RecordDate = time.Now()

	if err := repository.AddNewCompany(data); err != nil {
		log.Println("Error inserting company:", err)
		return err
	}
	return nil
}

func UpdateCompany(existingData model.Company, newData model.Company) (model.Company, error) {
	log.Println("update company")

	updateFields := make(map[string]interface{})

	if newData.CompanyCode != "" {
		updateFields["company_code"] = newData.CompanyCode
	}

	if newData.CompanyLogoUrl != "" {
		updateFields["company_logo_url"] = newData.CompanyLogoUrl
	}

	if newData.CompanyName != "" {
		updateFields["company_name"] = newData.CompanyName
	}

	updateFields["expired_date"] = newData.ExpiredDate
	updateFields["update_date"] = time.Now()

	err := repository.UpdateCompany(existingData.CompanyID, updateFields)
	if err != nil {
		return model.Company{}, err
	}
	return existingData, nil
}
