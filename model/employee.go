package model

type Employee struct {
	EmployeeID      int       `gorm:"column:role_id;primaryKey;autoIncrement" json:"roleId"`
	EmployeeName    string    `gorm:"column:role_name;type:varchar(255)" json:"roleName"`
	EmployeeUsername    string    `gorm:"column:employee_username;type:varchar(255)" json:"roleName"`
	Password    string    `gorm:"column:password;type:varchar(255)" json:"password"`
	Email    string    `gorm:"column:email;type:varchar(50)" json:"email"`
	Phone    string    `gorm:"column:phone;type:varchar(50)" json:"phone"`
	IDCard    string    `gorm:"column:id_card;type:varchar(30)" json:"id_card"`
	BirthDate    string    `gorm:"column:birth_date;type:date" json:"birthDate"`
	BirthPlace    string    `gorm:"column:birth_place;type:varchar(100)" json:"birthPlace"`
	Address    string    `gorm:"column:address;type:text" json:"address"`
	LastLogin    string    `gorm:"column:last_login;type:datetime" json:"lastLogin"`
	GeneralTime
	CompanyID int    `gorm:"column:company_id;not null" json:"companyId"`                 // Foreign key linking to Company
    Company   Company `gorm:"foreignKey:CompanyID;references:CompanyID"` // Belongs to one Company
	RoleID int    `gorm:"column:role_id;not null" json:"roleId"`                 // Foreign key linking to Company
    Role   Role `gorm:"foreignKey:ROleID;references:RoleID"` // Belongs to one Company
}