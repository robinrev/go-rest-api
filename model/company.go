package model

import "time"

type Company struct {
	CompanyID      int       `gorm:"column:company_id;primaryKey;autoIncrement" json:"companyId"`
	CompanyName    string    `gorm:"column:company_name;type:varchar(255)" json:"companyName"`
	CompanyCode    string    `gorm:"column:company_code;unique:varchar(10)" json:"companyCode"`
	CompanyLogoUrl string    `gorm:"column:company_logo_url" json:"companyLogoUrl"`
	ExpiredDate    time.Time `gorm:"column:expired_date;type:date" json:"expiredDate"`
	GeneralTime
	Employees       []Employee `gorm:"foreignKey:EmployeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // One-to-many relationship
	Customers       []Customer `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // One-to-many relationship
	Items       []Item `gorm:"foreignKey:ItemID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // One-to-many relationship
}

func (Company) TableName() string {
	return "company"
}
