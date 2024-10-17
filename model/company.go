package model

import "time"

type Company struct {
	CompanyID      int       `gorm:"column:company_id;primaryKey;autoIncrement" json:"companyId"`
	CompanyName    string    `gorm:"column:company_name;type:varchar(255)" json:"companyName"`
	CompanyCode    string    `gorm:"column:company_code;unique:varchar(10)" json:"companyCode"`
	CompanyLogoUrl string    `gorm:"column:company_logo_url" json:"companyLogoUrl"`
	ExpiredDate    time.Time `gorm:"column:expired_date;type:date" json:"expiredDate"`
	GeneralTime
}

func (Company) TableName() string {
	return "company"
}
