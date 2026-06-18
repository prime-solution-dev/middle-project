package models

import (
	"time"
)

type SaleBomItem struct {
	Id                   string    `gorm:"column:id" json:"id"`
	ProductCode          string    `gorm:"column:product_code" json:"product_code"`
	Qty                  float64   `gorm:"column:qty" json:"qty"`
	UnitCode             string    `gorm:"column:unit_code" json:"unit_code"`
	PriceUnit            float64   `gorm:"column:price_unit" json:"price_unit"`
	PriceAllocate        float64   `gorm:"column:price_allocate" json:"price_allocate"`
	PriceAllocatePercent float64   `gorm:"column:price_allocate_percent" json:"price_allocate_percent"`
	CreateBy             string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm            time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy             string    `gorm:"column:update_by" json:"update_by"`
	UpdateDtm            time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	SaleBomId            string    `gorm:"column:sale_bom_id" json:"sale_bom_id"`
}

func (SaleBomItem) TableName() string {
	return "sale_bom_item"
}
