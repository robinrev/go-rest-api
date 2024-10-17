package model

import "time"

type GeneralTime struct {
	CreateDate time.Time `gorm:"column:create_date;type:timestamp" json:"createDate"`
	UpdateDate time.Time `gorm:"column:update_date;type:timestamp" json:"updateDate"`
	RecordDate time.Time `gorm:"column:record_date;type:date" json:"recordDate"`
}
