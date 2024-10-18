package model

type Order struct {
	OrderID      int       `gorm:"column:order_id;primaryKey;autoIncrement" json:"orderId"`
	InvoiceNo    string    `gorm:"column:invoice_no;primaryKey;type:varchar(255)" json:"invoiceNo"`
	TotalWeight    float32    `gorm:"column:total_weight" json:"totalWeight"`
	TotalPoint    decimal.decimal    `gorm:"column:total_point" json:"totalPoint"`
	Tax1    decimal.Decimal    `gorm:"column:tax_1" json:"tax1"`
	Tax2    decimal.Decimal    `gorm:"column:tax_2" json:"tax2"`
	TotalTax    decimal.Decimal    `gorm:"column:total_tax" json:"totalTax"`
	SubTotal    decimal.Decimal    `gorm:"column:sub_total" json:"subTotal"`
	GrandTotal    decimal.Decimal    `gorm:"column:grand_total" json:"grandTotal"`
	Status    string    `gorm:"column:status;type:varchar(50)" json:"status"`
	EstimateFinishDate    time.Time    `gorm:"column:estimate_finish_date;type:date" json:"estimateFinishDate"`
	ActualFinishDate    time.Time    `gorm:"column:actual_finish_date;type:date" json:"actualFinishDate"`
	DeliveryDate    time.Time    `gorm:"column:delivery_date;type:date" json:"deliveryDate"`
	IsDelivered    bool    `gorm:"column:is_delivered" json:"isDelivered"`
	IsPaid    bool    `gorm:"column:is_paid" json:"isPaid"`
	GeneralTime
	OrderDetails       []OrderDetail `gorm:"foreignKey:OrderDetailID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // One-to-many relationship
	CompanyID int    `gorm:"column:company_id;not null" json:"companyId"`                 // Foreign key linking to Company
    Company   Company `gorm:"foreignKey:CompanyID;references:CompanyID"` // Belongs to one Company
	CustomerID int    `gorm:"column:customer_id;not null" json:"customerid"`                 // Foreign key linking to Company
    Customer   Company `gorm:"foreignKey:CustomerID;references:CustomerID"` // Belongs to one Company
	EmployeeID int    `gorm:"column:employee_id;not null" json:"employeeId"`                 // Foreign key linking to Company
    Employee   Company `gorm:"foreignKey:EmployeeID;references:EmployeeID"` // Belongs to one Company
}