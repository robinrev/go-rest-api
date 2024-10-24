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

func UpdateCustomer(param model.Customer) error {
	result := initializer.DB.Save(&param)
	return result.Error
}
