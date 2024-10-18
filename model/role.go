package model

type Role struct {
	RoleID      int       `gorm:"column:role_id;primaryKey;autoIncrement" json:"roleId"`
	RoleName    string    `gorm:"column:role_name;type:varchar(255)" json:"roleName"`
	Employees       []Employee `gorm:"foreignKey:EmployeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // One-to-many relationship
	GeneralTime
}