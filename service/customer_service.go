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

	if err := repository.AddNewCustomer(data); err != nil {
		log.Println("Error inserting customer:", err)
		return err
	}
	return nil
}

func UpdateCustomer(updateData model.Customer, newData model.Customer) (model.Customer, error) {
	log.Println("update customer")
	updateData.CustomerName = newData.CustomerName
	updateData.Address = newData.Address
	updateData.Email = newData.Email
	updateData.Phone = newData.Phone
	updateData.UpdateDate = time.Now()

	err := repository.UpdateCustomer(updateData)
	return updateData, err
}
