package model

type LogActivity struct {
	LogActivityID   int    `gorm:"column:log_activity_id;primaryKey;autoIncrement" json:"logCompanyId"`
	LogActivityName string `gorm:"column:log_activity_name" json:"logCompanyName"`
	CompanyID       int    `gorm:"column:company_id" json:"companyId"`
	EmployeeID      int    `gorm:"column:employee_id" json:"employeeId"`
	GeneralTime
}

func (LogActivity) TableName() string {
	return "log_activity"
}
