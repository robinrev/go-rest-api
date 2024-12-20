package repository

import (
	"log"

	"github.com/robinrev/go-rest-api/initializer"
	"github.com/robinrev/go-rest-api/model"
	"github.com/robinrev/go-rest-api/util"
	"gorm.io/gorm"
)

func GetAllCustomer() ([]model.Customer, error) {
	var data []model.Customer
	result := initializer.DB.Find(&data)

	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return data, nil
}

func GetAllCustomersByCompany(companyID int) ([]model.Customer, error) {
	var data []model.Customer
	result := initializer.DB.Where("company_id = ?", companyID).Find(&data)

	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return data, nil
}

func GetCustomerById(id string) (model.Customer, error) {
	var data model.Customer
	result := initializer.DB.First(&data, "customer_id = ?", id)

	if result.Error != nil {
		return model.Customer{}, result.Error
	} else if result.RowsAffected == 0 {
		return model.Customer{}, gorm.ErrRecordNotFound
	}

	return data, nil
}

func AddNewCustomer(param model.Customer) error {
	AddNewLogActivity(util.LOG_ACTIVITY_ADD_CUSTOMER)

	result := initializer.DB.Create(&param)
	if result.Error != nil {
		log.Println("failed to add customer")
	}
	return result.Error
}

func UpdateCustomer(customerID int, updateFields map[string]interface{}) error {
	result := initializer.DB.Model(&model.Customer{}).Where("customer_id = ?", customerID).Updates(updateFields)
	return result.Error
}
