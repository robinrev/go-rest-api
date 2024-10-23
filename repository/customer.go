package repository

import (
	"errors"
	"log"
	"time"

	"github.com/robinrev/go-rest-api/initializer"
	"github.com/robinrev/go-rest-api/model"
	"github.com/robinrev/go-rest-api/util"
	"gorm.io/gorm"
)

func GetAllCustomer() ([]model.Customer, error) {
	var data []model.Customer
	if err := initializer.DB.Find(&data).Error; err != nil {
		return data, err
	}

	if len(data) == 0 {
		return data, errors.New("no customers found")
	}
	return data, nil
}

func GetCustomerById(id string) (model.Customer, error) {
	var data model.Customer

	// Fetch the customer with the given ID
	if err := initializer.DB.First(&data, "customer_id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return data, err
		}
		// Handle any other database error
		return data, err
	}

	return data, nil
}

func AddNewCustomer(param model.Customer) error {
	AddNewLogActivity(util.LOG_ACTIVITY_ADD_CUSTOMER)
	param.CreateDate = time.Now()
	param.UpdateDate = time.Now()
	param.RecordDate = time.Now()

	result := initializer.DB.Create(&param)
	if result.Error != nil {
		log.Println("failed to add customer")
	}
	return result.Error
}

func UpdateCustomer(param model.Customer) error {
	param.UpdateDate = time.Now()
	result := initializer.DB.Save(&param)
	return result.Error
}
