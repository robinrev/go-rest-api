package model

type Item struct {
	ItemID   int    `gorm:"column:item_id;primaryKey;autoIncrement" json:"itemId"`
	ItemName string `gorm:"column:item_name;type:varchar(255)" json:"itemName"`
	GeneralTime
	OrderDetails []OrderDetail `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"orderDetails"`
	CompanyID    int           `gorm:"column:company_id;not null" json:"companyId"`
}

func (Item) TableName() string {
	return "item"
}
