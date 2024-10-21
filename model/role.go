package model

type Role struct {
	RoleID   int    `gorm:"column:role_id;primaryKey;autoIncrement" json:"roleId"`
	RoleName string `gorm:"column:role_name;type:varchar(255)" json:"roleName"`
	GeneralTime
	Employees []Employee `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // One-to-many relationship
}

func (Role) TableName() string {
	return "role"
}
