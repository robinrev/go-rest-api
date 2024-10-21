package model

type Employee struct {
	EmployeeID       int    `gorm:"column:employee_id;primaryKey;autoIncrement" json:"employeeId"`
	EmployeeName     string `gorm:"column:employee_name;type:varchar(255)" json:"employeeName"`
	EmployeeUsername string `gorm:"column:employee_username;type:varchar(255)" json:"employeeUsername"`
	Password         string `gorm:"column:password;type:varchar(255)" json:"password"`
	Email            string `gorm:"column:email;type:varchar(50)" json:"email"`
	Phone            string `gorm:"column:phone;type:varchar(50)" json:"phone"`
	IDCard           string `gorm:"column:id_card;type:varchar(30)" json:"id_card"`
	BirthDate        string `gorm:"column:birth_date;type:date" json:"birthDate"`
	BirthPlace       string `gorm:"column:birth_place;type:varchar(100)" json:"birthPlace"`
	Address          string `gorm:"column:address;type:text" json:"address"`
	LastLogin        string `gorm:"column:last_login;type:timestamp" json:"lastLogin"`
	GeneralTime
	Orders    []Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"orders"`
	CompanyID int     `gorm:"column:company_id;not null" json:"companyId"`
	RoleID    int     `gorm:"column:role_id;not null" json:"roleId"`
}

func (Employee) TableName() string {
	return "employee"
}
