package model

type OrderDetail struct {
	OrderDetailID      int       `gorm:"column:order_detail_id;primaryKey;autoIncrement" json:"orderDEtailId"`
	ItemName    string    `gorm:"column:item_name;primaryKey;type:varchar(255)" json:"itemName"`
	IsCharged    bool    `gorm:"column:is_charged" json:"isCHarged"`
	GeneralTime
	OrderID int    `gorm:"column:order_id;not null" json:"orderId"`                 // Foreign key linking to Company
    Order   Order `gorm:"foreignKey:OrderID;references:OrderID"` // Belongs to one Company
	ItemID int    `gorm:"column:item_id;not null" json:"itemId"`                 // Foreign key linking to Company
    Item   Item `gorm:"foreignKey:ItemID;references:ItemID"` // Belongs to one Company
}