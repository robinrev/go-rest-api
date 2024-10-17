package repository

import (
	"time"

	"github.com/robinrev/go-rest-api/initializer"
	"github.com/robinrev/go-rest-api/model"
	"github.com/robinrev/go-rest-api/util"
)

func GetAllCompany() []model.Company {
	var company []model.Company
	initializer.DB.Find(&company)
	return company
}

func GetCompanyById(id string) model.Company {
	var company model.Company
	initializer.DB.Find(&company, id)
	return company
}

func GetCompanyByCompanyCode(companyCode string) model.Company {
	var company model.Company
	initializer.DB.Where("company_code = ?", companyCode).First(&company)
	return company
}

func AddNewCompany(company model.Company) error {
	AddNewLogActivity(util.LOG_ACTIVITY_ADD_COMPANY)
	company.RecordDate = time.Now()
	result := initializer.DB.Create(&company)
	return result.Error
}

func UpdateCompany(company model.Company) error {
	company.UpdateDate = time.Now()
	result := initializer.DB.Save(&company)
	return result.Error
}
