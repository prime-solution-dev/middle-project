package models

import (
	"time"
)

type PriceList struct {
	Id            string    `gorm:"column:id" json:"id"`
	ProductCode   string    `gorm:"column:product_code" json:"product_code"`
	PriceUnit     float64   `gorm:"column:price_unit" json:"price_unit"`
	Currency      string    `gorm:"column:currency" json:"currency"`
	EffectiveDate time.Time `gorm:"column:effective_date" json:"effective_date"`
	Remark        string    `gorm:"column:remark" json:"remark"`
	IsActive      bool      `gorm:"column:is_active" json:"is_active"`
	CreateBy      string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm     time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy      string    `gorm:"column:update_by" json:"update_by"`
	UpdateDtm     time.Time `gorm:"column:update_dtm" json:"update_dtm"`
}

func (PriceList) TableName() string {
	return "price_list"
}
