package model

type Item struct {
	ItemID      int       `gorm:"column:item_id;primaryKey;autoIncrement" json:"itemId"`
	ItemName    string    `gorm:"column:item_name;type:varchar(255)" json:"itemName"`
	GeneralTime
	OrderDetails       []OrderDetail `gorm:"foreignKey:OrderDetailID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // One-to-many relationship
	CompanyID int    `gorm:"column:company_id;not null" json:"companyId"`                 // Foreign key linking to Company
    Company   Company `gorm:"foreignKey:CompanyID;references:CompanyID"` // Belongs to one Company
}