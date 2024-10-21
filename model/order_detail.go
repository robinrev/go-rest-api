package model

type OrderDetail struct {
	OrderDetailID int    `gorm:"column:order_detail_id;primaryKey;autoIncrement" json:"orderDetailId"`
	ItemName      string `gorm:"column:item_name;type:varchar(255)" json:"itemName"`
	IsCharged     bool   `gorm:"column:is_charged" json:"isCharged"`
	GeneralTime
	OrderID int `gorm:"column:order_id;not null" json:"orderId"`
	ItemID  int `gorm:"column:item_id;not null" json:"itemId"`
}

func (OrderDetail) TableName() string {
	return "order_detail"
}
