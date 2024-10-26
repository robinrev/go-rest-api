package dto

import "time"

type Company struct {
	CompanyID      int       `json:"companyId"`
	CompanyName    string    `json:"companyName"`
	CompanyCode    string    `json:"companyCode"`
	CompanyLogoUrl string    `json:"companyLogoUrl"`
	ExpiredDate    time.Time `json:"expiredDate"`
}
