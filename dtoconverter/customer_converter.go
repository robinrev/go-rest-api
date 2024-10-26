package dtoconverter

import (
	"github.com/robinrev/go-rest-api/dto"
	"github.com/robinrev/go-rest-api/model"
)

func CustomerModelToDto(customer model.Customer) dto.CustomerDto {
	return dto.CustomerDto{
		CustomerID:   customer.CustomerID,
		CustomerName: customer.CustomerName,
		Username:     customer.Username,
		Password:     customer.Password,
		Email:        customer.Email,
		Phone:        customer.Phone,
		LoyaltyPoint: customer.LoyaltyPoint,
		BirthDate:    customer.BirthDate,
		BirthPlace:   customer.BirthPlace,
		Address:      customer.Address,
	}
}

func CustomerListModelToDto(customers []model.Customer) []dto.CustomerDto {
	listDto := make([]dto.CustomerDto, len(customers))

	for i, customer := range customers {
		listDto[i] = CustomerModelToDto(customer)
	}

	return listDto
}
