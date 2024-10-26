package service

import (
	"errors"
	"log"
	"time"

	"github.com/robinrev/go-rest-api/model"
	"github.com/robinrev/go-rest-api/repository"
)

func GetAllCustomer() ([]model.Customer, error) {
	log.Println("get all customer")
	return repository.GetAllCustomer()
}

func GetAllCustomersByCompany(companyID int) ([]model.Customer, error) {
	log.Println("get all customer")
	return repository.GetAllCustomersByCompany(companyID)
}

func GetCustomerById(id string) (model.Customer, error) {
	log.Println("get customer by id :", id)
	return repository.GetCustomerById(id)
}

func AddNewCustomer(data model.Customer) error {
	log.Println("add new customer")

	if data.Username == "" || data.Password == "" {
		return errors.New("username and password are required")
	}

	data.CreateDate = time.Now()
	data.UpdateDate = time.Now()
	data.RecordDate = time.Now()
	data.LastLogin = time.Time{}

	if err := repository.AddNewCustomer(data); err != nil {
		log.Println("Error inserting customer:", err)
		return err
	}
	return nil
}

func UpdateCustomer(existingData model.Customer, newData model.Customer) (model.Customer, error) {
	log.Println("update customer")

	updateFields := make(map[string]interface{})

	if newData.Address != "" {
		updateFields["address"] = newData.Address
	}

	if newData.BirthDate != "" {
		updateFields["birth_date"] = newData.BirthDate
	}

	if newData.BirthPlace != "" {
		updateFields["birth_place"] = newData.BirthPlace
	}

	if newData.CustomerName != "" {
		updateFields["customer_name"] = newData.CustomerName
	}

	if newData.Email != "" {
		updateFields["email"] = newData.Email
	}

	if newData.Password != "" {
		updateFields["password"] = newData.Password
	}

	if newData.Username != "" {
		updateFields["username"] = newData.Username
	}

	if newData.LoyaltyPoint != 0 {
		updateFields["loyalty_point"] = newData.LoyaltyPoint
	}

	if newData.LoyaltyPoint != 0 {
		updateFields["last_login"] = newData.LastLogin
	}

	updateFields["update_date"] = time.Now()

	err := repository.UpdateCustomer(existingData.CustomerID, updateFields)
	if err != nil {
		return model.Customer{}, err
	}
	return existingData, nil
}
