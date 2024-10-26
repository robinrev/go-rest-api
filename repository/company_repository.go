package repository

import (
	"log"

	"github.com/robinrev/go-rest-api/initializer"
	"github.com/robinrev/go-rest-api/model"
	"github.com/robinrev/go-rest-api/util"
	"gorm.io/gorm"
)

func GetAllCompanies() ([]model.Company, error) {
	var company []model.Company
	result := initializer.DB.Find(&company)

	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return company, nil
}

func GetCompanyByCompanyCode(companyCode string) model.Company {
	var company model.Company
	initializer.DB.Where("company_code = ?", companyCode).First(&company)
	return company
}

func GetCompanyById(id string) (model.Company, error) {
	var data model.Company
	result := initializer.DB.First(&data, "company_id = ?", id)

	if result.Error != nil {
		return model.Company{}, result.Error
	} else if result.RowsAffected == 0 {
		return model.Company{}, gorm.ErrRecordNotFound
	}

	return data, nil
}

func AddNewCompany(param model.Company) error {
	AddNewLogActivity(util.LOG_ACTIVITY_ADD_CUSTOMER)

	result := initializer.DB.Create(&param)
	if result.Error != nil {
		log.Println("failed to add company")
	}
	return result.Error
}

func UpdateCompany(companyID int, updateFields map[string]interface{}) error {
	result := initializer.DB.Model(&model.Company{}).Where("company_id = ?", companyID).Updates(updateFields)
	return result.Error
}
