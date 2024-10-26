package dtoconverter

import (
	"github.com/robinrev/go-rest-api/dto"
	"github.com/robinrev/go-rest-api/model"
)

func CompanyModelToDto(company model.Company) dto.Company {
	return dto.Company{
		CompanyID:      company.CompanyID,
		CompanyName:    company.CompanyName,
		CompanyCode:    company.CompanyCode,
		CompanyLogoUrl: company.CompanyLogoUrl,
		ExpiredDate:    company.ExpiredDate,
	}
}

func CompanyListModelToDto(companies []model.Company) []dto.Company {
	listDto := make([]dto.Company, len(companies))

	for i, company := range companies {
		listDto[i] = CompanyModelToDto(company)
	}

	return listDto
}
