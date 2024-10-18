package model

type Customer struct {
	CustomerID      int       `gorm:"column:customer_id;primaryKey;autoIncrement" json:"customerId"`
	CustomerName    string    `gorm:"column:customer_name;type:varchar(255)" json:"customerName"`
	Username    string    `gorm:"column:username;type:varchar(255)" json:"username"`
	Password    string    `gorm:"column:password;type:varchar(255)" json:"password"`
	Email    string    `gorm:"column:email;type:varchar(50)" json:"email"`
	Phone    string    `gorm:"column:phone;type:varchar(50)" json:"phone"`
	LoyaltyPoint    int    `gorm:"column:loyalty_point" json:"loyaltyPoint"`
	BirthDate    string    `gorm:"column:birth_date;type:date" json:"birthDate"`
	BirthPlace    string    `gorm:"column:birth_place;type:varchar(100)" json:"birthPlace"`
	Address    string    `gorm:"column:address;type:text" json:"address"`
	LastLogin    string    `gorm:"column:last_login;type:datetime" json:"lastLogin"`
	GeneralTime
	CompanyID int    `gorm:"column:company_id;not null" json:"companyId"`                 // Foreign key linking to Company
    Company   Company `gorm:"foreignKey:CompanyID;references:CompanyID"` // Belongs to one Company
}